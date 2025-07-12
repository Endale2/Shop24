package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var discountColl *mongo.Collection = config.GetCollection("DRPS", "discounts")

func CreateDiscount(d *models.Discount) (*mongo.InsertOneResult, error) {
	// ID and timestamps stamped by service or repo; choose one. If here:
	if d.ID.IsZero() {
		d.ID = primitive.NewObjectID()
	}
	now := time.Now()
	if d.CreatedAt.IsZero() {
		d.CreatedAt = now
	}
	d.UpdatedAt = now
	return discountColl.InsertOne(context.Background(), d)
}

func GetDiscountByID(id primitive.ObjectID) (*models.Discount, error) {
	var d models.Discount
	err := discountColl.FindOne(context.Background(), bson.M{"_id": id}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func ListDiscountsByShop(shopID primitive.ObjectID) ([]models.Discount, error) {
	cursor, err := discountColl.Find(context.Background(), bson.M{"shop_id": shopID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var list []models.Discount
	for cursor.Next(context.Background()) {
		var d models.Discount
		if err := cursor.Decode(&d); err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func UpdateDiscount(id primitive.ObjectID, upd bson.M) (*mongo.UpdateResult, error) {
	upd["updated_at"] = time.Now()
	return discountColl.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": upd})
}

// UpdateDiscountWithOperators allows using different MongoDB update operators
func UpdateDiscountWithOperators(id primitive.ObjectID, update bson.M) (*mongo.UpdateResult, error) {
	// Ensure updated_at is always set
	if _, exists := update["$set"]; exists {
		update["$set"].(bson.M)["updated_at"] = time.Now()
	} else {
		update["$set"] = bson.M{"updated_at": time.Now()}
	}
	return discountColl.UpdateOne(context.Background(), bson.M{"_id": id}, update)
}

func DeleteDiscount(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return discountColl.DeleteOne(context.Background(), bson.M{"_id": id})
}

// Active queries: product-level only (collections removed)
func GetActiveDiscountsForProduct(shopID, productID, variantID primitive.ObjectID, collectionIDs []primitive.ObjectID) ([]models.Discount, error) {
	now := time.Now()
	orClauses := []bson.M{
		{"applies_to_products": productID},
	}

	// Only add variant clause if variantID is not zero
	if !variantID.IsZero() {
		orClauses = append(orClauses, bson.M{"applies_to_variants": variantID})
	}

	// Build time-based filters that handle zero times properly
	timeFilters := []bson.M{
		{"active": true},
		{"category": models.DiscountCategoryProduct},
	}

	// Add start time filter (if start_at is zero, it means "started")
	timeFilters = append(timeFilters, bson.M{
		"$or": []bson.M{
			{"start_at": bson.M{"$lte": now}},
			{"start_at": time.Time{}}, // Zero time means "started"
		},
	})

	// Add end time filter (if end_at is zero, it means "no end")
	timeFilters = append(timeFilters, bson.M{
		"$or": []bson.M{
			{"end_at": bson.M{"$gte": now}},
			{"end_at": time.Time{}}, // Zero time means "no end"
		},
	})

	filter := bson.M{
		"shop_id": shopID,
		"$and":    timeFilters,
		"$or":     orClauses,
	}

	fmt.Printf("Discount Query - ShopID: %s, ProductID: %s, VariantID: %s\n",
		shopID.Hex(), productID.Hex(), variantID.Hex())
	fmt.Printf("Discount Query - Filter: %+v\n", filter)
	fmt.Printf("Discount Query - Time filters: %+v\n", timeFilters)
	fmt.Printf("Discount Query - OR clauses: %+v\n", orClauses)

	cursor, err := discountColl.Find(context.Background(), filter)
	if err != nil {
		fmt.Printf("Discount Query - Error: %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	var results []models.Discount
	for cursor.Next(context.Background()) {
		var d models.Discount
		if err := cursor.Decode(&d); err != nil {
			fmt.Printf("Discount Query - Decode error: %v\n", err)
			return nil, err
		}
		// Additional validation can be done in the service layer
		results = append(results, d)
		fmt.Printf("Discount Query - Found discount: %s (Type: %s, Value: %.2f)\n",
			d.Name, d.Type, d.Value)
	}

	fmt.Printf("Discount Query - Total results: %d\n", len(results))
	return results, nil
}
