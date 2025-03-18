package repositories

import (
	"context"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/config" // Assume this provides the MongoDB connection
	"go.mongodb.org/mongo-driver/bson"
)

// FindAdminByEmail searches for an admin by email.
func FindAdminByEmail(email string) (*models.Admin, error) {
	collection := config.GetCollection("your_database", "admins")
	var admin models.Admin
	filter := bson.M{"email": email}
	err := collection.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// InsertAdmin inserts a new admin record.
func InsertAdmin(admin *models.Admin) error {
	collection := config.GetCollection("your_database", "admins")
	_, err := collection.InsertOne(context.Background(), admin)
	return err
}
