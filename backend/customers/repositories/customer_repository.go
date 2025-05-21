package repositories

import (
	"context"

	"github.com/Endale2/DRPS/customers/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var customerCollection *mongo.Collection = config.GetCollection("DRPS", "customers")

// CreateCustomer inserts a new Customer document.
func CreateCustomer(customer *models.Customer) (*mongo.InsertOneResult, error) {
	return customerCollection.InsertOne(context.Background(), customer)
}

// GetCustomerByID retrieves a Customer by its ObjectID.
func GetCustomerByID(id primitive.ObjectID) (*models.Customer, error) {
	var customer models.Customer
	err := customerCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
