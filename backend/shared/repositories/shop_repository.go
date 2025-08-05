package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var shopCollection *mongo.Collection = config.GetCollection("DRPS", "shops")

// Createshop inserts a new shop into the collection.
func CreateShop(shop *models.Shop) (*mongo.InsertOneResult, error) {
	shop.ID = primitive.NewObjectID()
	shop.CreatedAt = time.Now()
	shop.UpdatedAt = time.Now()

	return shopCollection.InsertOne(context.Background(), shop)
}

// GetshopByID retrieves a shop by its ID.
func GetShopByID(id string) (*models.Shop, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}

	var shop models.Shop
	err = shopCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&shop)
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

// GetAllshops returns all shops in the collection.
func GetAllShops() ([]models.Shop, error) {
	cursor, err := shopCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var shops []models.Shop
	for cursor.Next(context.Background()) {
		var shop models.Shop
		if err := cursor.Decode(&shop); err != nil {
			return nil, err
		}
		shops = append(shops, shop)
	}
	return shops, nil
}

// Updateshop updates fields of a shop identified by its ID.
func UpdateShop(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}

	updatedData["updatedAt"] = time.Now()

	return shopCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": updatedData},
	)
}

// Deleteshop removes a shop from the collection.
func DeleteShop(id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}

	return shopCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
}

// GetShopBySlug retrieves a shop by its slug.
func GetShopBySlug(slug string) (*models.Shop, error) {
	var shop models.Shop
	err := shopCollection.FindOne(context.Background(), bson.M{"slug": slug}).Decode(&shop)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &shop, nil
}
