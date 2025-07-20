package repositories

import (
	"context"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/customers/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// GetCustomerByEmail retrieves a Customer by its email.
func GetCustomerByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	err := customerCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

// UpdateCustomerByID updates a customer by ObjectID with the given fields.
func UpdateCustomerByID(id primitive.ObjectID, update bson.M) error {
	_, err := customerCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": update},
	)
	return err
}
