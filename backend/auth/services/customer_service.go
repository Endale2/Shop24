package services

import (
	"errors"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	customerModels "github.com/Endale2/DRPS/customers/models"
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerRegisterService registers a new customer by creating both a Customer record (detailed data)
// and an AuthCustomer record (authentication data), linking them via CustomerID.

func CustomerRegisterService(authCustomer *models.AuthCustomer) error {
	
	// Validate required fields.
	if authCustomer.Email == "" || authCustomer.Password == "" || authCustomer.Username == "" {
		return errors.New("missing required fields")
	}

	// Check if an authentication record already exists.
	if existing, _ := authRepo.FindAuthCustomerByEmail(authCustomer.Email); existing != nil {
		return errors.New("customer already exists")
	}

	// Create a new detailed Customer record.
	newCustomer := &customerModels.Customer{
		Name:  authCustomer.Username, // Use username as default display name.
		Email: authCustomer.Email,
		// Additional fields such as Phone or Address can be added here.
	}

	// Insert the new customer record.
	res, err := customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return err
	}
	newCustomer.ID = res.InsertedID.(primitive.ObjectID)

	// Link the AuthCustomer record to the new Customer.
	authCustomer.CustomerID = newCustomer.ID

	// Create the AuthCustomer record.
	_, err = authRepo.CreateAuthCustomer(authCustomer)
	return err
}

// CustomerLoginService validates login credentials and retrieves both authentication data and the detailed Customer record.
func CustomerLoginService(authCustomer *models.AuthCustomer) (*models.AuthCustomer, *customerModels.Customer, error) {
	// Look up the AuthCustomer record by email.
	foundAuth, err := authRepo.FindAuthCustomerByEmail(authCustomer.Email)
	if err != nil {
		return nil, nil, errors.New("customer not found")
	}

	// For simplicity, compare passwords directly. In production, use secure hashed comparisons.
	if authCustomer.Password != foundAuth.Password {
		return nil, nil, errors.New("password mismatch")
	}

	// Retrieve the detailed Customer record using the stored CustomerID.
	customerData, err := customerRepo.GetCustomerByID(foundAuth.CustomerID)
	if err != nil {
		return nil, nil, errors.New("customer details not found")
	}

	return foundAuth, customerData, nil
}
