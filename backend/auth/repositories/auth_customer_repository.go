package repositories

import (
	"context"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authCustomerCollection *mongo.Collection = config.GetCollection("DRPS", "authcustomers")

// CreateAuthCustomer inserts a new AuthCustomer document.
func CreateAuthCustomer(authCustomer *models.AuthCustomer) (*mongo.InsertOneResult, error) {
	return authCustomerCollection.InsertOne(context.Background(), authCustomer)
}

// FindAuthCustomerByEmail retrieves an AuthCustomer by email.
func FindAuthCustomerByEmail(email string) (*models.AuthCustomer, error) {
	var authCustomer models.AuthCustomer
	err := authCustomerCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&authCustomer)
	if err != nil {
		if err == mongo.ErrNoDocuments { // IMPORTANT: Handle no documents explicitly
			return nil, nil
		}
		return nil, err
	}
	return &authCustomer, nil
}

// FindAuthCustomerByProvider retrieves an AuthCustomer by provider and provider ID.
func FindAuthCustomerByProvider(provider, providerID string) (*models.AuthCustomer, error) {
	var rec models.AuthCustomer
	filter := bson.M{"provider": provider, "provider_id": providerID}
	err := authCustomerCollection.FindOne(context.Background(), filter).Decode(&rec)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil, nil if no document is found
		}
		return nil, err
	}
	return &rec, nil
}