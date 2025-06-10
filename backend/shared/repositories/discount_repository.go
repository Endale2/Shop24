package repositories

import (
	"context"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var discountCollection *mongo.Collection = config.GetCollection("DRPS", "discounts")

// CreateDiscount inserts a new Discount document, assigning ID and timestamps.
func CreateDiscount(d *models.Discount) (*mongo.InsertOneResult, error) {
	// assign ID and timestamps here if not already set
	if d.ID.IsZero() {
		d.ID = primitive.NewObjectID()
	}
	now := time.Now()
	if d.CreatedAt.IsZero() {
		d.CreatedAt = now
	}
	d.UpdatedAt = now
	return discountCollection.InsertOne(context.Background(), d)
}

// GetDiscountByID looks up a Discount document by its ObjectID.
func GetDiscountByID(id primitive.ObjectID) (*models.Discount, error) {
	var d models.Discount
	err := discountCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

// GetDiscountByCouponCode retrieves a single Discount by coupon_code (case‐sensitive).
func GetDiscountByCouponCode(code string) (*models.Discount, error) {
	var d models.Discount
	filter := bson.M{"coupon_code": code}
	err := discountCollection.FindOne(context.Background(), filter).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

// ListDiscountsByShop returns all discounts for a given shop.
func ListDiscountsByShop(shopID primitive.ObjectID) ([]models.Discount, error) {
	cursor, err := discountCollection.Find(context.Background(), bson.M{"shop_id": shopID})
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

// UpdateDiscount updates fields of a Discount by its ObjectID.
func UpdateDiscount(id primitive.ObjectID, upd bson.M) (*mongo.UpdateResult, error) {
	// ensure updated_at is set
	upd["updated_at"] = time.Now()
	return discountCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": upd},
	)
}

// DeleteDiscount removes a Discount document by its ObjectID.
func DeleteDiscount(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return discountCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// GetActiveDiscountsForShop retrieves all “active” discounts for a given shop, whose date range includes now.
func GetActiveDiscountsForShop(shopID primitive.ObjectID) ([]models.Discount, error) {
	now := time.Now()
	filter := bson.M{
		"shop_id":  shopID,
		"active":   true,
		"start_at": bson.M{"$lte": now},
		"end_at":   bson.M{"$gte": now},
	}
	cursor, err := discountCollection.Find(context.Background(), filter)
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

// GetActiveDiscountsForProduct retrieves all active “product‐category” discounts that apply to this product or its variants.
func GetActiveDiscountsForProduct(
	shopID, productID, variantID, customerID primitive.ObjectID,
) ([]models.Discount, error) {
	now := time.Now()
	filter := bson.M{
		"shop_id":  shopID,
		"active":   true,
		"start_at": bson.M{"$lte": now},
		"end_at":   bson.M{"$gte": now},
		"category": models.DiscountCategoryProduct,
		"$or": []bson.M{
			{"applies_to_products": productID},
			{"applies_to_variants": variantID},
		},
	}

	cursor, err := discountCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.Discount
	for cursor.Next(context.Background()) {
		var d models.Discount
		if err := cursor.Decode(&d); err != nil {
			return nil, err
		}
		// If AllowedCustomerIDs is non-empty, ensure customerID is listed
		if len(d.AllowedCustomerIDs) > 0 {
			allowed := false
			for _, cid := range d.AllowedCustomerIDs {
				if cid == customerID {
					allowed = true
					break
				}
			}
			if !allowed {
				continue
			}
		}
		results = append(results, d)
	}
	return results, nil
}

// GetActiveOrderDiscountsForShop retrieves all “order‐category” discounts for this shop.
func GetActiveOrderDiscountsForShop(shopID primitive.ObjectID) ([]models.Discount, error) {
	now := time.Now()
	filter := bson.M{
		"shop_id":  shopID,
		"active":   true,
		"start_at": bson.M{"$lte": now},
		"end_at":   bson.M{"$gte": now},
		"category": models.DiscountCategoryOrder,
	}
	cursor, err := discountCollection.Find(context.Background(), filter)
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

// GetActiveShippingDiscountsForShop retrieves all “shipping‐category” discounts for this shop.
func GetActiveShippingDiscountsForShop(shopID primitive.ObjectID) ([]models.Discount, error) {
	now := time.Now()
	filter := bson.M{
		"shop_id":  shopID,
		"active":   true,
		"start_at": bson.M{"$lte": now},
		"end_at":   bson.M{"$gte": now},
		"category": models.DiscountCategoryShipping,
	}
	cursor, err := discountCollection.Find(context.Background(), filter)
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

