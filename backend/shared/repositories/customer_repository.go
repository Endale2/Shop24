package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var customerCollection *mongo.Collection = config.GetCollection("yourDatabaseName", "customers")

// CreateCustomer inserts a new customer into the collection.
func CreateCustomer(customer *models.Customer) (*mongo.InsertOneResult, error) {
	customer.ID = primitive.NewObjectID()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	return customerCollection.InsertOne(context.Background(), customer)
}

// GetCustomerByID retrieves a customer by its ID.
func GetCustomerByID(id string) (*models.Customer, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid customer ID")
	}

	var customer models.Customer
	err = customerCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetAllCustomers returns all customers in the collection.
func GetAllCustomers() ([]models.Customer, error) {
	cursor, err := customerCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var customers []models.Customer
	for cursor.Next(context.Background()) {
		var customer models.Customer
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


// GetCustomersWithFilter returns customers matching a filter (e.g. shopId).
func GetCustomersWithFilter(filter bson.M) ([]models.Customer, error) {
    cursor, err := customerCollection.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var list []models.Customer
    for cursor.Next(context.Background()) {
        var c models.Customer
        if err := cursor.Decode(&c); err != nil {
            return nil, err
        }
        list = append(list, c)
    }
    return list, nil
}
