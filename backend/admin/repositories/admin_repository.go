package repositories

import (
	"context"

	"github.com/Endale2/DRPS/admin/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// ListAdmins retrieves all Admin documents.
func ListAdmins() ([]models.Admin, error) {
	var admins []models.Admin
	cursor, err := adminCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var admin models.Admin
		if err := cursor.Decode(&admin); err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}

// UpdateAdmin updates an Admin document by its ObjectID.
func UpdateAdmin(id primitive.ObjectID, updates bson.M) error {
	_, err := adminCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": updates})
	return err
}

// DeleteAdmin deletes an Admin document by its ObjectID.
func DeleteAdmin(id primitive.ObjectID) error {
	_, err := adminCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
