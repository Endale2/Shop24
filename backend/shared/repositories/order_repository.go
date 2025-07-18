// shared/repositories/order_repository.go
package repositories

import (
	"context"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCol *mongo.Collection = config.GetCollection("DRPS", "orders")

func CreateOrder(ctx context.Context, o *models.Order) (*mongo.InsertOneResult, error) {
	return orderCol.InsertOne(ctx, o)
}

func GetOrderByID(ctx context.Context, hexID string) (*models.Order, error) {
	var o models.Order
	id, _ := primitive.ObjectIDFromHex(hexID)
	err := orderCol.FindOne(ctx, bson.M{"_id": id}).Decode(&o)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &o, err
}

func ListOrders(ctx context.Context, filter bson.M) ([]models.Order, error) {
	cur, err := orderCol.Find(ctx, filter, options.Find().SetSort(bson.D{{"created_at", -1}}))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []models.Order
	for cur.Next(ctx) {
		var o models.Order
		if err := cur.Decode(&o); err != nil {
			continue
		}
		out = append(out, o)
	}
	return out, err
}

// ListOrdersPaginated returns paginated orders with total count
func ListOrdersPaginated(ctx context.Context, filter bson.M, page, limit int) ([]models.Order, int64, error) {
	// Calculate skip value
	skip := (page - 1) * limit

	// Get total count
	total, err := orderCol.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	findOptions := options.Find().
		SetSort(bson.D{{"created_at", -1}}).
		SetSkip(int64(skip)).
		SetLimit(int64(limit))

	cur, err := orderCol.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []models.Order
	for cur.Next(ctx) {
		var o models.Order
		if err := cur.Decode(&o); err != nil {
			continue
		}
		out = append(out, o)
	}

	return out, total, err
}

func UpdateOrder(ctx context.Context, hexID string, upd bson.M) (*mongo.UpdateResult, error) {
	id, _ := primitive.ObjectIDFromHex(hexID)
	return orderCol.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": upd})
}

func DeleteOrder(ctx context.Context, hexID string) (*mongo.DeleteResult, error) {
	id, _ := primitive.ObjectIDFromHex(hexID)
	return orderCol.DeleteOne(ctx, bson.M{"_id": id})
}

func CountOrders(ctx context.Context, filter bson.M) (int64, error) {
	return orderCol.CountDocuments(ctx, filter)
}
