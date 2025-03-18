package repositories

import (
	"context"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authSellerCollection *mongo.Collection = config.GetCollection("your_database", "authsellers")

// CreateAuthSeller inserts a new AuthSeller document.
func CreateAuthSeller(authSeller *models.AuthSeller) (*mongo.InsertOneResult, error) {
	return authSellerCollection.InsertOne(context.Background(), authSeller)
}

// FindAuthSellerByEmail retrieves an AuthSeller document by email.
func FindAuthSellerByEmail(email string) (*models.AuthSeller, error) {
	var authSeller models.AuthSeller
	err := authSellerCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&authSeller)
	if err != nil {
		return nil, err
	}
	return &authSeller, nil
}
