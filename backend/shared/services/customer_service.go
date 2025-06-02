package services

import (
	custModel "github.com/Endale2/DRPS/customers/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateCustomerService creates a new customer.
func CreateCustomerService(customer *custModel.Customer) (*mongo.InsertOneResult, error) {
	return repositories.CreateCustomer(customer)
}

// GetCustomerByIDService retrieves a customer by its ID.
func GetCustomerByIDService(id string) (*custModel.Customer, error) {
	return repositories.GetCustomerByID(id)
}

// GetAllCustomersService returns all customers.
func GetAllCustomersService() ([]custModel.Customer, error) {
	return repositories.GetAllCustomers()
}

// UpdateCustomerService updates fields of a customer identified by its ID.
func UpdateCustomerService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	return repositories.UpdateCustomer(id, updatedData)
}

// DeleteCustomerService removes a customer by its ID.
func DeleteCustomerService(id string) (*mongo.DeleteResult, error) {
	return repositories.DeleteCustomer(id)
}
