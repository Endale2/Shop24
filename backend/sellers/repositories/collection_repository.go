// seller/repositories/collection_repository.go
package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/sellers/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionColl *mongo.Collection = config.GetCollection("DRPS", "collections")

// CreateCollection inserts a new Collection document.
func CreateCollection(coll *models.Collection) (*mongo.InsertOneResult, error) {
	coll.ID = primitive.NewObjectID()
	coll.CreatedAt = time.Now()
	coll.UpdatedAt = time.Now()
	return collectionColl.InsertOne(context.Background(), coll)
}

// GetCollectionsByShop returns all collections for a given shop.
func GetCollectionsByShop(shopID primitive.ObjectID) ([]models.Collection, error) {
	filter := bson.M{"shop_id": shopID}
	cursor, err := collectionColl.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var result []models.Collection
	for cursor.Next(context.Background()) {
		var c models.Collection
		if err := cursor.Decode(&c); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

// GetCollectionByID looks up a Collection by its ObjectID.
func GetCollectionByID(id primitive.ObjectID) (*models.Collection, error) {
	var c models.Collection
	err := collectionColl.FindOne(context.Background(), bson.M{"_id": id}).Decode(&c)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

// UpdateCollection updates fields of a Collection by its ObjectID.
func UpdateCollection(id primitive.ObjectID, updates bson.M) (*mongo.UpdateResult, error) {
	updates["updated_at"] = time.Now()
	return collectionColl.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
}

// DeleteCollection removes a Collection by its ObjectID.
func DeleteCollection(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	if id.IsZero() {
		return nil, errors.New("invalid collection ID")
	}
	return collectionColl.DeleteOne(context.Background(), bson.M{"_id": id})
}

// AddProductToCollection pushes a productID into the collection's product_ids array.
func AddProductToCollection(collID, prodID primitive.ObjectID) (*mongo.UpdateResult, error) {
	return collectionColl.UpdateOne(
		context.Background(),
		bson.M{"_id": collID},
		bson.M{"$addToSet": bson.M{"product_ids": prodID}},
	)
}

// RemoveProductFromCollection pulls a productID out of the collection's product_ids array.
func RemoveProductFromCollection(collID, prodID primitive.ObjectID) (*mongo.UpdateResult, error) {
	return collectionColl.UpdateOne(
		context.Background(),
		bson.M{"_id": collID},
		bson.M{"$pull": bson.M{"product_ids": prodID}},
	)
}

// GetCollectionByHandle looks up a collection by its unique handle and shop ID.
func GetCollectionByHandle(shopID primitive.ObjectID, handle string) (*models.Collection, error) {
	var coll models.Collection
	filter := bson.M{"shop_id": shopID, "handle": handle}
	err := config.GetCollection("DRPS", "collections").
		FindOne(context.Background(), filter).
		Decode(&coll)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &coll, nil
}
