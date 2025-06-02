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

var cartCollection *mongo.Collection = config.GetCollection("DRPS", "carts")

// CreateCart inserts a new Cart document.
func CreateCart(cart *models.Cart) (*mongo.InsertOneResult, error) {
	cart.ID = primitive.NewObjectID()
	cart.CreatedAt = time.Now()
	cart.LastUpdated = time.Now()
	return cartCollection.InsertOne(context.Background(), cart)
}

// GetCartByCustomerID finds a cart for a given customer (shop-specific).
func GetCartByCustomerID(shopID, customerID primitive.ObjectID) (*models.Cart, error) {
	var cart models.Cart
	filter := bson.M{
		"shop_id":     shopID,
		"customer_id": customerID,
	}
	err := cartCollection.FindOne(context.Background(), filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

// GetCartBySessionID finds a cart for an anonymous (guest) session.
func GetCartBySessionID(shopID primitive.ObjectID, sessionID string) (*models.Cart, error) {
	var cart models.Cart
	filter := bson.M{
		"shop_id":    shopID,
		"session_id": sessionID,
	}
	err := cartCollection.FindOne(context.Background(), filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

// UpdateCart replaces the cart document (used when items change).
func UpdateCart(cart *models.Cart) (*mongo.UpdateResult, error) {
	cart.LastUpdated = time.Now()
	update := bson.M{"$set": cart}
	return cartCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": cart.ID},
		update,
	)
}

// DeleteCart deletes a cart by its ID.
func DeleteCart(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return cartCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// AddOrUpdateCartItem inserts or updates a CartItem in the Cart’s Items array.
func AddOrUpdateCartItem(cartID primitive.ObjectID, item *models.CartItem) (*mongo.UpdateResult, error) {
	item.ID = primitive.NewObjectID()
	// If a CartItem for the same ProductID+VariantID exists, replace it; else push new.
	filter := bson.M{
		"_id":            cartID,
		"items.product_id": item.ProductID,
		"items.variant_id": item.VariantID,
	}
	// First, try to update an existing item
	updateResult, err := cartCollection.UpdateOne(
		context.Background(),
		filter,
		bson.M{
			"$set": bson.M{
				"items.$.quantity":        item.Quantity,
				"items.$.unit_price":      item.UnitPrice,
				"items.$.line_total":      item.LineTotal,
				"items.$.discount_amount": item.DiscountAmount,
				"items.$.final_line_total": item.FinalLineTotal,
				"items.$.applied_discount_ids": item.AppliedDiscountIDs,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	// If no existing item was updated, push the new one
	if updateResult.MatchedCount == 0 {
		return cartCollection.UpdateOne(
			context.Background(),
			bson.M{"_id": cartID},
			bson.M{"$push": bson.M{"items": item}},
		)
	}
	return updateResult, nil
}

// RemoveCartItem removes a CartItem by its CartItem ID.
func RemoveCartItem(cartID, itemID primitive.ObjectID) (*mongo.UpdateResult, error) {
	return cartCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": cartID},
		bson.M{"$pull": bson.M{"items": bson.M{"_id": itemID}}},
	)
}

// Helper to clear all items (e.g. after checkout or empty cart).
func ClearCartItems(cartID primitive.ObjectID) (*mongo.UpdateResult, error) {
	return cartCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": cartID},
		bson.M{"$set": bson.M{"items": []models.CartItem{}}},
	)
}

// Helper to fetch by cart ID
func GetCartByID(id primitive.ObjectID) (*models.Cart, error) {
	var cart models.Cart
	err := cartCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

// UpsertCart creates or updates a cart in one call. Useful if you want to set session/customer in one shot.
func UpsertCart(cart *models.Cart) (*mongo.UpdateResult, error) {
	cart.LastUpdated = time.Now()
	filter := bson.M{"shop_id": cart.ShopID}
	if cart.CustomerID != nil {
		filter["customer_id"] = cart.CustomerID
	} else if cart.SessionID != nil {
		filter["session_id"] = *cart.SessionID
	} else {
		return nil, errors.New("either customer_id or session_id must be set in Cart")
	}
	update := bson.M{"$set": cart}
	return cartCollection.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{Upsert: ptrBool(true)},
	)
}

// helper to get a pointer to bool for Upsert option
func ptrBool(b bool) *bool {
	return &b
}
