package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = config.GetCollection("yourDatabaseName", "products")

// CreateProduct inserts a new product into the collection.
func CreateProduct(product *models.Product) (*mongo.InsertOneResult, error) {
	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	return productCollection.InsertOne(context.Background(), product)
}

// GetProductByID retrieves a product by its ID.
func GetProductByID(id string) (*models.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	var product models.Product
	err = productCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAllProducts returns all products in the collection.
func GetAllProducts() ([]models.Product, error) {
	cursor, err := productCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	for cursor.Next(context.Background()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// UpdateProduct updates fields of a product identified by its ID.
func UpdateProduct(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	updatedData["updatedAt"] = time.Now()

	return productCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": updatedData},
	)
}

// DeleteProduct removes a product from the collection.
func DeleteProduct(id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	return productCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
}
