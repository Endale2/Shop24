package services

import (
	"context"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWishlistService(shopID, customerID primitive.ObjectID) (*models.Wishlist, error) {
	return repositories.GetWishlist(context.Background(), shopID, customerID)
}

func AddProductToWishlistService(shopID, customerID, productID primitive.ObjectID) (*models.Wishlist, error) {
	return repositories.CreateOrUpdateWishlist(context.Background(), shopID, customerID, productID)
}

func RemoveProductFromWishlistService(shopID, customerID, productID primitive.ObjectID) error {
	return repositories.RemoveProductFromWishlist(context.Background(), shopID, customerID, productID)
}
