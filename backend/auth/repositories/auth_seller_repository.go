package repositories

import (
	"context"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authSellerCollection = config.GetCollection("DRPS", "authsellers")

func CreateAuthSeller(a *models.AuthSeller) (*mongo.InsertOneResult, error) {
    return authSellerCollection.InsertOne(context.Background(), a)
}

func FindAuthSellerByEmail(email string) (*models.AuthSeller, error) {
    var a models.AuthSeller
    err := authSellerCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&a)
    if err == mongo.ErrNoDocuments {
        return nil, nil
    }
    return &a, err
}

func FindAuthSellerByProvider(provider, providerID string) (*models.AuthSeller, error) {
	var rec models.AuthSeller
	filter := bson.M{"provider": provider, "provider_id": providerID}
	err := authSellerCollection.FindOne(context.Background(), filter).Decode(&rec)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &rec, nil
}