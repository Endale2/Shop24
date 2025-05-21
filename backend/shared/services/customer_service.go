package services

import (
	"github.com/Endale2/DRPS/customers/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateCustomerService creates a new customer.
func CreateCustomerService(customer *models.Customer) (*mongo.InsertOneResult, error) {
	// Add any business logic or validations here (e.g., check email format).
	return repositories.CreateCustomer(customer)
}

// GetCustomerByIDService retrieves a customer by its ID.
func GetCustomerByIDService(id string) (*models.Customer, error) {
	return repositories.GetCustomerByID(id)
}

// GetAllCustomersService returns all customers.
func GetAllCustomersService() ([]models.Customer, error) {
	return repositories.GetAllCustomers()
}

// UpdateCustomerService updates fields of a customer identified by its ID.
func UpdateCustomerService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	// Place to add service-level validations or business rules (e.g., phone number format).
	return repositories.UpdateCustomer(id, updatedData)
}

// DeleteCustomerService removes a customer by its ID.
func DeleteCustomerService(id string) (*mongo.DeleteResult, error) {
	return repositories.DeleteCustomer(id)
}
