package repositories

import (
	"context"

	"github.com/Endale2/DRPS/sellers/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var sellerCollection *mongo.Collection = config.GetCollection("your_database", "sellers")

// CreateSeller inserts a new Seller document.
func CreateSeller(seller *models.Seller) (*mongo.InsertOneResult, error) {
	return sellerCollection.InsertOne(context.Background(), seller)
}

// GetSellerByID retrieves a Seller document by its ObjectID.
func GetSellerByID(id primitive.ObjectID) (*models.Seller, error) {
	var seller models.Seller
	err := sellerCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&seller)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}
