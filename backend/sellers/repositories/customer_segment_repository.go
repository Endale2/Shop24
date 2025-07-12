package repositories

import (
	"context"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/sellers/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var customerSegmentColl *mongo.Collection = config.GetCollection("DRPS", "customer_segments")

// CreateCustomerSegment inserts a new CustomerSegment document
func CreateCustomerSegment(segment *models.CustomerSegment) (*mongo.InsertOneResult, error) {
	segment.ID = primitive.NewObjectID()
	segment.CreatedAt = time.Now()
	segment.UpdatedAt = time.Now()
	return customerSegmentColl.InsertOne(context.Background(), segment)
}

// GetCustomerSegmentsByShop returns all customer segments for a given shop
func GetCustomerSegmentsByShop(shopID primitive.ObjectID) ([]models.CustomerSegment, error) {
	filter := bson.M{"shop_id": shopID}
	cursor, err := customerSegmentColl.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var result []models.CustomerSegment
	for cursor.Next(context.Background()) {
		var segment models.CustomerSegment
		if err := cursor.Decode(&segment); err != nil {
			return nil, err
		}
		result = append(result, segment)
	}
	return result, nil
}

// GetCustomerSegmentByID looks up a CustomerSegment by its ObjectID
func GetCustomerSegmentByID(id primitive.ObjectID) (*models.CustomerSegment, error) {
	var segment models.CustomerSegment
	err := customerSegmentColl.FindOne(context.Background(), bson.M{"_id": id}).Decode(&segment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &segment, nil
}

// UpdateCustomerSegment updates fields of a CustomerSegment by its ObjectID
func UpdateCustomerSegment(id primitive.ObjectID, updates bson.M) (*mongo.UpdateResult, error) {
	updates["updated_at"] = time.Now()
	return customerSegmentColl.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
}

// DeleteCustomerSegment removes a CustomerSegment by its ObjectID
func DeleteCustomerSegment(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return customerSegmentColl.DeleteOne(context.Background(), bson.M{"_id": id})
}

// AddCustomerToSegment adds a customer to a segment
func AddCustomerToSegment(segmentID primitive.ObjectID, customerID primitive.ObjectID) (*mongo.UpdateResult, error) {
	return customerSegmentColl.UpdateOne(
		context.Background(),
		bson.M{"_id": segmentID},
		bson.M{
			"$addToSet": bson.M{"customer_ids": customerID},
			"$set":      bson.M{"updated_at": time.Now()},
		},
	)
}

// RemoveCustomerFromSegment removes a customer from a segment
func RemoveCustomerFromSegment(segmentID primitive.ObjectID, customerID primitive.ObjectID) (*mongo.UpdateResult, error) {
	return customerSegmentColl.UpdateOne(
		context.Background(),
		bson.M{"_id": segmentID},
		bson.M{
			"$pull": bson.M{"customer_ids": customerID},
			"$set":  bson.M{"updated_at": time.Now()},
		},
	)
}

// GetSegmentsByCustomer returns all segments that contain a specific customer
func GetSegmentsByCustomer(shopID primitive.ObjectID, customerID primitive.ObjectID) ([]models.CustomerSegment, error) {
	filter := bson.M{
		"shop_id":      shopID,
		"customer_ids": customerID,
	}
	cursor, err := customerSegmentColl.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var result []models.CustomerSegment
	for cursor.Next(context.Background()) {
		var segment models.CustomerSegment
		if err := cursor.Decode(&segment); err != nil {
			return nil, err
		}
		result = append(result, segment)
	}
	return result, nil
}

// GetCustomersInSegment returns all customers in a specific segment
func GetCustomersInSegment(segmentID primitive.ObjectID) ([]primitive.ObjectID, error) {
	var segment models.CustomerSegment
	err := customerSegmentColl.FindOne(context.Background(), bson.M{"_id": segmentID}).Decode(&segment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []primitive.ObjectID{}, nil
		}
		return nil, err
	}
	return segment.CustomerIDs, nil
}
