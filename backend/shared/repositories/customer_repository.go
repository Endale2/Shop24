package repositories

import (
	"context"
	"errors"
	"time"

	custModel "github.com/Endale2/DRPS/customers/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var customerCollection *mongo.Collection = config.GetCollection("yourDatabaseName", "customers")

// CreateCustomer inserts a new customer into the collection.
func CreateCustomer(customer *custModel.Customer) (*mongo.InsertOneResult, error) {
	customer.ID = primitive.NewObjectID()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	return customerCollection.InsertOne(context.Background(), customer)
}

// GetCustomerByID retrieves a customer by its ID.
func GetCustomerByID(id string) (*custModel.Customer, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid customer ID")
	}

	var customer custModel.Customer
	err = customerCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetAllCustomers returns all customers in the collection.
func GetAllCustomers() ([]custModel.Customer, error) {
	cursor, err := customerCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var customers []custModel.Customer
	for cursor.Next(context.Background()) {
		var customer custModel.Customer
		if err := cursor.Decode(&customer); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

// UpdateCustomer updates fields of a customer identified by its ID.
func UpdateCustomer(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid customer ID")
	}

	updatedData["updatedAt"] = time.Now()

	return customerCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": updatedData},
	)
}

// DeleteCustomer removes a customer from the collection.
func DeleteCustomer(id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid customer ID")
	}

	return customerCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
}
