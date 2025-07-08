package repositories

import (
	"context"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var wishlistCol = config.GetCollection("DRPS", "wishlists")

func GetWishlist(ctx context.Context, shopID, customerID primitive.ObjectID) (*models.Wishlist, error) {
	var w models.Wishlist
	err := wishlistCol.FindOne(ctx, bson.M{"shop_id": shopID, "customer_id": customerID}).Decode(&w)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func CreateOrUpdateWishlist(ctx context.Context, shopID, customerID primitive.ObjectID, productID primitive.ObjectID) (*models.Wishlist, error) {
	now := time.Now()
	filter := bson.M{"shop_id": shopID, "customer_id": customerID}
	update := bson.M{
		"$setOnInsert": bson.M{
			"created_at": now,
		},
		"$set": bson.M{
			"updated_at": now,
		},
		"$addToSet": bson.M{
			"product_ids": productID,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	var w models.Wishlist
	err := wishlistCol.FindOneAndUpdate(ctx, filter, update, opts).Decode(&w)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func RemoveProductFromWishlist(ctx context.Context, shopID, customerID, productID primitive.ObjectID) error {
	filter := bson.M{"shop_id": shopID, "customer_id": customerID}
	update := bson.M{
		"$pull": bson.M{"product_ids": productID},
		"$set":  bson.M{"updated_at": time.Now()},
	}
	_, err := wishlistCol.UpdateOne(ctx, filter, update)
	return err
}
