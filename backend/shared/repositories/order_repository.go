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

// orderCollection is the MongoDB collection for orders.
var orderCollection *mongo.Collection = config.GetCollection("DRPS", "orders")

// CreateOrder inserts a new order into the collection.
func CreateOrder(order *models.Order) (*mongo.InsertOneResult, error) {
	order.ID = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	return orderCollection.InsertOne(context.Background(), order)
}

// GetOrderByID retrieves an order by its hex ID.
func GetOrderByID(id string) (*models.Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	var ord models.Order
	err = orderCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&ord)
	if err != nil {
		return nil, err
	}
	return &ord, nil
}

// GetAllOrders returns all orders in the collection.
func GetAllOrders() ([]models.Order, error) {
	cursor, err := orderCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []models.Order
	for cursor.Next(context.Background()) {
		var ord models.Order
		if err := cursor.Decode(&ord); err != nil {
			return nil, err
		}
		orders = append(orders, ord)
	}
	return orders, nil
}

// UpdateOrder updates fields of an order by its hex ID.
func UpdateOrder(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	updatedData["updated_at"] = time.Now()

	return orderCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": updatedData},
	)
}

// DeleteOrder removes an order from the collection by its hex ID.
func DeleteOrder(id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	return orderCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
}