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

var orderCollection *mongo.Collection = config.GetCollection("yourDatabaseName", "orders")

// CreateOrder inserts a new order and returns the insert result.
func CreateOrder(order *models.Order) (*mongo.InsertOneResult, error) {
    order.ID = primitive.NewObjectID()
    order.CreatedAt = time.Now()
    order.UpdatedAt = time.Now()
    return orderCollection.InsertOne(context.Background(), order)
}

// GetOrderByID retrieves an order by hex ID.
func GetOrderByID(id string) (*models.Order, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, errors.New("invalid order ID")
    }
    var ord models.Order
    if err = orderCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&ord); err != nil {
        return nil, err
    }
    return &ord, nil
}

// GetOrdersByUser returns all orders for a given user.
func GetOrdersByUser(userID string) ([]models.Order, error) {
    objID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return nil, errors.New("invalid user ID")
    }
    cursor, err := orderCollection.Find(context.Background(), bson.M{"user_id": objID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var orders []models.Order
    for cursor.Next(context.Background()) {
        var o models.Order
        if err := cursor.Decode(&o); err != nil {
            return nil, err
        }
        orders = append(orders, o)
    }
    return orders, nil
}

// UpdateOrderStatus updates payment or shipping status.
func UpdateOrderStatus(id string, update bson.M) (*mongo.UpdateResult, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, errors.New("invalid order ID")
    }
    update["updated_at"] = time.Now()
    return orderCollection.UpdateOne(
        context.Background(),
        bson.M{"_id": objID},
        bson.M{"$set": update},
    )
}