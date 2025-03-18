package repositories

import (
	"context"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authAdminCollection *mongo.Collection = config.GetCollection("your_database", "authadmins")

// CreateAuthAdmin inserts a new AuthAdmin document.
func CreateAuthAdmin(authAdmin *models.AuthAdmin) (*mongo.InsertOneResult, error) {
	return authAdminCollection.InsertOne(context.Background(), authAdmin)
}

// FindAuthAdminByEmail retrieves an AuthAdmin document by email.
func FindAuthAdminByEmail(email string) (*models.AuthAdmin, error) {
	var authAdmin models.AuthAdmin
	err := authAdminCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&authAdmin)
	if err != nil {
		return nil, err
	}
	return &authAdmin, nil
}
