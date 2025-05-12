package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Endale2/DRPS/sellers/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var sellerCollection *mongo.Collection = config.GetCollection("your_database", "sellers")

// CreateSeller inserts a new Seller document and sets timestamps.
func CreateSeller(seller *models.Seller) (*mongo.InsertOneResult, error) {
	seller.ID = primitive.NewObjectID()
	seller.CreatedAt = time.Now()
	seller.UpdatedAt = time.Now()

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

// GetAllSellers returns all Seller documents in the collection.
func GetAllSellers() ([]models.Seller, error) {
	cursor, err := sellerCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var sellers []models.Seller
	for cursor.Next(context.Background()) {
		var s models.Seller
		if err := cursor.Decode(&s); err != nil {
			return nil, err
		}
		sellers = append(sellers, s)
	}
	return sellers, nil
}

// UpdateSeller updates fields of a Seller identified by its ObjectID.
func UpdateSeller(id primitive.ObjectID, updatedData bson.M) (*mongo.UpdateResult, error) {
	if id.IsZero() {
		return nil, errors.New("invalid seller ID")
	}
	// set updated timestamp
	updatedData["updated_at"] = time.Now()

	return sellerCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
}

// DeleteSeller removes a Seller document by its ObjectID.
func DeleteSeller(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	if id.IsZero() {
		return nil, errors.New("invalid seller ID")
	}

	return sellerCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}