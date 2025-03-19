package repositories

import (
	"context"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authCustomerCollection *mongo.Collection = config.GetCollection("your_database", "authcustomers")

// CreateAuthCustomer inserts a new AuthCustomer document.
func CreateAuthCustomer(authCustomer *models.AuthCustomer) (*mongo.InsertOneResult, error) {
	return authCustomerCollection.InsertOne(context.Background(), authCustomer)
}

// FindAuthCustomerByEmail retrieves an AuthCustomer by email.
func FindAuthCustomerByEmail(email string) (*models.AuthCustomer, error) {
	var authCustomer models.AuthCustomer
	err := authCustomerCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&authCustomer)
	if err != nil {
		return nil, err
	}
	return &authCustomer, nil
}
