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

var shopCategoryCollection *mongo.Collection = config.GetCollection("DRPS", "shop_categories")

func CreateShopCategory(category *models.ShopCategory) (*mongo.InsertOneResult, error) {
	category.ID = primitive.NewObjectID()
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	return shopCategoryCollection.InsertOne(context.Background(), category)
}

func GetShopCategoryByID(id string) (*models.ShopCategory, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	var category models.ShopCategory
	err = shopCategoryCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func GetAllShopCategories() ([]models.ShopCategory, error) {
	cursor, err := shopCategoryCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var categories []models.ShopCategory
	for cursor.Next(context.Background()) {
		var category models.ShopCategory
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func UpdateShopCategory(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	updatedData["updatedAt"] = time.Now()
	return shopCategoryCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": updatedData},
	)
}

func DeleteShopCategory(id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	return shopCategoryCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
}

func GetShopCategoriesByIDs(ids []primitive.ObjectID) ([]models.ShopCategory, error) {
	cursor, err := shopCategoryCollection.Find(context.Background(), bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var categories []models.ShopCategory
	for cursor.Next(context.Background()) {
		var category models.ShopCategory
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
