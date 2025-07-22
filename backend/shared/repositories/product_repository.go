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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var productCollection *mongo.Collection = config.GetCollection("DRPS", "products")

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

// CountProducts returns the total number of products matching the given filter.
func CountProducts(filter bson.M) (int64, error) {
	count, err := productCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetProductsByFilter retrieves products matching an arbitrary filter.
func GetProductsByFilter(filter bson.M) ([]models.Product, error) {
	cursor, err := productCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	for cursor.Next(context.Background()) {
		var p models.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// GetProductsByShopSlug retrieves all products for the shop with the given slug.
func GetProductsByShopSlug(slug string) ([]models.Product, error) {
	// 1) Lookup the shop by slug
	shop, err := GetShopBySlug(slug)
	if err != nil {
		return nil, err
	}
	if shop == nil {
		// no such shop
		return nil, nil
	}
	// 2) Use the existing filter helper
	return GetProductsByFilter(bson.M{"shop_id": shop.ID})
}

// GetProductBySlug retrieves a product by its slug.
func GetProductBySlug(slug string) (*models.Product, error) {
	var p models.Product
	err := productCollection.FindOne(
		context.Background(),
		bson.M{"slug": slug},
	).Decode(&p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// GetProductsByFilterPaginated retrieves products matching a filter, paginated, and returns total count.
// NOTE: Products are sorted by 'createdAt' descending (newest first) for performance and UX.
// Ensure there is an index on 'createdAt' for optimal performance.
func GetProductsByFilterPaginated(filter bson.M, page, limit int) ([]models.Product, int64, error) {
	skip := (page - 1) * limit

	total, err := productCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"createdAt", -1}}) // newest first
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	cursor, err := productCollection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	for cursor.Next(context.Background()) {
		var p models.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}
	return products, total, nil
}
