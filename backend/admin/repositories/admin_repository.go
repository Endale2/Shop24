package repositories

import (
	"context"

	"github.com/Endale2/DRPS/admin/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var adminCollection *mongo.Collection = config.GetCollection("DRPS", "admins")

// CreateAdmin inserts a new Admin document.
func CreateAdmin(admin *models.Admin) (*mongo.InsertOneResult, error) {
	return adminCollection.InsertOne(context.Background(), admin)
}

// GetAdminByID retrieves an Admin document by its ObjectID.
func GetAdminByID(id primitive.ObjectID) (*models.Admin, error) {
	var admin models.Admin
	err := adminCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
