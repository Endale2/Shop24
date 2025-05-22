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

// OptionalProfile holds non-required profile fields at registration.
type OptionalProfile struct {
	FirstName, LastName, Phone, Address, City, State, PostalCode, Country string
}

// CustomerRegisterService creates both Customer & AuthCustomer and returns full Customer.
func CustomerRegisterService(
	authCust *authModels.AuthCustomer,
	profile *OptionalProfile,
) (*customerModels.Customer, error) {
	// 1) Basic required
	if authCust.Username == "" || authCust.Email == "" || authCust.Password == "" {
		return nil, errors.New("username, email and password are required")
	}

	// 2) Dup check
	if existing, _ := authRepo.FindAuthCustomerByEmail(authCust.Email); existing != nil {
		return nil, errors.New("customer already exists")
	}

	// 3) Build Customer record, merge optional profile
	cust := &customerModels.Customer{
		UserName:   authCust.Username,
		Email:      authCust.Email,
		FirstName:  profile.FirstName,
		LastName:   profile.LastName,
		Phone:      profile.Phone,
		Address:    profile.Address,
		City:       profile.City,
		State:      profile.State,
		PostalCode: profile.PostalCode,
		Country:    profile.Country,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	res, err := customerRepo.CreateCustomer(cust)
	if err != nil {
		return nil, err
	}

	// 4) Link AuthCustomer → Customer
	oid := res.InsertedID.(primitive.ObjectID)
	authCust.CustomerID = oid
	if _, err := authRepo.CreateAuthCustomer(authCust); err != nil {
		return nil, err
	}

	// Return the complete customer
	cust.ID = oid
	return cust, nil
}

// CustomerLoginService authenticates a customer and returns Customer + tokens.
func CustomerLoginService(email, password string) (*customerModels.Customer, string, string, error) {
	// 1) Lookup
	foundAuth, err := authRepo.FindAuthCustomerByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("customer not found")
	}
	// 2) Password check
	if password != foundAuth.Password {
		return nil, "", "", errors.New("invalid credentials")
	}
	// 3) Load Customer
	cust, err := customerRepo.GetCustomerByID(foundAuth.CustomerID)
	if err != nil {
		return nil, "", "", errors.New("failed to load customer data")
	}
	// 4) Tokens
	at, err := utils.CreateToken(foundAuth.CustomerID.Hex(), 5*time.Minute)
	if err != nil {
		return nil, "", "", errors.New("failed to generate access token")
	}
	rt, err := utils.CreateToken(foundAuth.CustomerID.Hex(), 7*24*time.Hour)
	if err != nil {
		return nil, "", "", errors.New("failed to generate refresh token")
	}
	return cust, at, rt, nil
}
