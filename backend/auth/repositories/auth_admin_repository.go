package repositories

import (
    "context"

    "github.com/Endale2/DRPS/auth/models"
    "github.com/Endale2/DRPS/config"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var authAdminColl *mongo.Collection = config.GetCollection("DRPS", "authadmins")

func CreateAuthAdmin(a *models.AuthAdmin) (*mongo.InsertOneResult, error) {
    return authAdminColl.InsertOne(context.Background(), a)
}

func FindAuthAdminByEmail(email string) (*models.AuthAdmin, error) {
    var a models.AuthAdmin
    err := authAdminColl.FindOne(context.Background(), bson.M{"email": email, "provider": "local"}).Decode(&a)
    if err != nil {
        return nil, err
    }
    return &a, nil
}

func FindAuthAdminByProvider(provider, providerID string) (*models.AuthAdmin, error) {
    var a models.AuthAdmin
    err := authAdminColl.FindOne(context.Background(), bson.M{
        "provider":    provider,
        "provider_id": providerID,
    }).Decode(&a)
    if err != nil {
        return nil, err
    }
    return &a, nil
}
