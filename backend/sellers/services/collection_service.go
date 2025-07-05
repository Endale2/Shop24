// seller/services/collection_service.go
package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/sellers/models"
	"github.com/Endale2/DRPS/sellers/repositories"
	sharedShopService "github.com/Endale2/DRPS/shared/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CollectionService encapsulates business logic for seller collections.
type CollectionService struct{}

// NewCollectionService returns a new service instance.
func NewCollectionService() *CollectionService {
	return &CollectionService{}
}

// CreateCollectionService creates a new collection for a given shop.
func (svc *CollectionService) CreateCollectionService(
	shopID primitive.ObjectID,
	title, description, handle, image string,
) (*models.Collection, error) {

	// Verify this shop exists
	shop, err := sharedShopService.GetShopByIDService(shopID.Hex())
	if err != nil {
		return nil, err
	}
	if shop == nil {
		return nil, errors.New("shop not found")
	}

	// Build the Collection model
	coll := &models.Collection{
		ShopID:      shopID,
		Title:       title,
		Description: description,
		Handle:      handle,
		Image:       image,
		ProductIDs:  []primitive.ObjectID{},
		Filters:     nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = repositories.CreateCollection(coll)
	if err != nil {
		return nil, err
	}
	return coll, nil
}

// GetCollectionsByShopService returns all collections for a given shop.
func (svc *CollectionService) GetCollectionsByShopService(shopID primitive.ObjectID) ([]models.Collection, error) {
	return repositories.GetCollectionsByShop(shopID)
}

// UpdateCollectionService updates a collection's fields (title, description, handle, image).
// Only allowed if sellerID matches collection.ShopID owner (checked in controller).
func (svc *CollectionService) UpdateCollectionService(collID primitive.ObjectID, sellerID primitive.ObjectID, updates bson.M) error {
	// Verify collection exists and belongs to seller
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to update this collection")
	}

	_, err = repositories.UpdateCollection(collID, updates)
	return err
}

// DeleteCollectionService removes a collection.
// Only allowed if sellerID matches collection.ShopID owner (checked in controller).
func (svc *CollectionService) DeleteCollectionService(collID primitive.ObjectID, sellerID primitive.ObjectID) error {
	// Verify collection exists and belongs to seller
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to delete this collection")
	}

	_, err = repositories.DeleteCollection(collID)
	return err
}

// AddProductToCollectionService adds a product to a collection.
// Only allowed if sellerID matches collection.ShopID owner (checked in controller).
func (svc *CollectionService) AddProductToCollectionService(collID primitive.ObjectID, productID primitive.ObjectID, sellerID primitive.ObjectID) error {
	// Verify collection exists and belongs to seller
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to modify this collection")
	}

	// Check if product is already in collection
	for _, existingID := range coll.ProductIDs {
		if existingID == productID {
			return errors.New("product already in collection")
		}
	}

	// Add product to collection
	updates := bson.M{
		"product_ids": append(coll.ProductIDs, productID),
		"updated_at":  time.Now(),
	}
	_, err = repositories.UpdateCollection(collID, updates)
	return err
}

// RemoveProductFromCollectionService removes a product from a collection.
// Only allowed if sellerID matches collection.ShopID owner (checked in controller).
func (svc *CollectionService) RemoveProductFromCollectionService(collID primitive.ObjectID, productID primitive.ObjectID, sellerID primitive.ObjectID) error {
	// Verify collection exists and belongs to seller
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to modify this collection")
	}

	// Remove product from collection
	var newProductIDs []primitive.ObjectID
	for _, existingID := range coll.ProductIDs {
		if existingID != productID {
			newProductIDs = append(newProductIDs, existingID)
		}
	}

	updates := bson.M{
		"product_ids": newProductIDs,
		"updated_at":  time.Now(),
	}
	_, err = repositories.UpdateCollection(collID, updates)
	return err
}

// ReplaceCollectionProductsService replaces all products in a collection with new ones.
// Only allowed if sellerID matches collection.ShopID owner (checked in controller).
func (svc *CollectionService) ReplaceCollectionProductsService(collID primitive.ObjectID, productIDStrings []string, shopID primitive.ObjectID, sellerID primitive.ObjectID) error {
	// Verify collection exists and belongs to seller
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to modify this collection")
	}

	// Convert string IDs to ObjectIDs and validate they belong to this shop
	var newProductIDs []primitive.ObjectID
	for _, productIDStr := range productIDStrings {
		productID, err := primitive.ObjectIDFromHex(productIDStr)
		if err != nil {
			continue // skip invalid product IDs
		}

		// Verify product belongs to this shop
		product, err := sharedShopService.GetProductByIDService(productID.Hex())
		if err == nil && product != nil && product.ShopID == shopID {
			newProductIDs = append(newProductIDs, productID)
		}
	}

	// Update collection with new product IDs
	updates := bson.M{
		"product_ids": newProductIDs,
		"updated_at":  time.Now(),
	}
	_, err = repositories.UpdateCollection(collID, updates)
	return err
}
