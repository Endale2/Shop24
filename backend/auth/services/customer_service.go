package services

import (
	"errors"
	"time"

	authModels "github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	customerModels "github.com/Endale2/DRPS/customers/models"
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	"github.com/Endale2/DRPS/auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerRegisterService registers a new customer:
// 1) Inserts a Customer record.
// 2) Sets authCustomer.CustomerID to that new ID.
// 3) Inserts an AuthCustomer record.
func CustomerRegisterService(authCustomer *authModels.AuthCustomer) error {
	// 1) Required fields
	if authCustomer.Username == "" || authCustomer.Email == "" || authCustomer.Password == "" {
		return errors.New("username, email and password are required")
	}

	// 2) Check for existing auth record
	if existing, _ := authRepo.FindAuthCustomerByEmail(authCustomer.Email); existing != nil {
		return errors.New("customer already exists")
	}

	// 3) Create detailed Customer record (no Phone/Address on AuthCustomer)
	newCust := &customerModels.Customer{
		UserName:  authCustomer.Username,
		Email: authCustomer.Email,
		// if you want to capture phone/address, add those fields to AuthCustomer model
	}
	res, err := customerRepo.CreateCustomer(newCust)
	if err != nil {
		return err
	}

	// 4) Link AuthCustomer → Customer
	newID := res.InsertedID.(primitive.ObjectID)
	authCustomer.CustomerID = newID

	// 5) Insert AuthCustomer record
	_, err = authRepo.CreateAuthCustomer(authCustomer)
	return err
}

// CustomerLoginService authenticates a customer and issues JWTs.
// Returns Customer profile, accessToken, refreshToken, error.
func CustomerLoginService(authCustomer *authModels.AuthCustomer) (*customerModels.Customer, string, string, error) {
	// 1) Lookup credentials
	foundAuth, err := authRepo.FindAuthCustomerByEmail(authCustomer.Email)
	if err != nil {
		return nil, "", "", errors.New("customer not found")
	}

	// 2) Verify password
	if authCustomer.Password != foundAuth.Password {
		return nil, "", "", errors.New("invalid credentials")
	}

	// 3) Load full Customer record
	custData, err := customerRepo.GetCustomerByID(foundAuth.CustomerID)
	if err != nil {
		return nil, "", "", errors.New("failed to load customer data")
	}

	// 4) Issue JWTs via utils
	accessToken, err := utils.CreateToken(foundAuth.CustomerID.Hex(), 5*time.Minute)
	if err != nil {
		return nil, "", "", errors.New("failed to generate access token")
	}
	refreshToken, err := utils.CreateToken(foundAuth.CustomerID.Hex(), 7*24*time.Hour)
	if err != nil {
		return nil, "", "", errors.New("failed to generate refresh token")
	}

	return custData, accessToken, refreshToken, nil
}
