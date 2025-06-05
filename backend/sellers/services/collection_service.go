// seller/services/collection_service.go
package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/sellers/models"
	sharedShopService "github.com/Endale2/DRPS/shared/services"
	"github.com/Endale2/DRPS/sellers/repositories"
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

// UpdateCollectionService updates a collection’s fields (title, description, handle, image).
// Only allowed if sellerID matches collection.ShopID owner (checked in controller).
func (svc *CollectionService) UpdateCollectionService(
	collID, sellerID primitive.ObjectID,
	updates bson.M,
) error {
	// 1) Fetch collection
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	// 2) Verify shop’s owner
	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop.OwnerID != sellerID {
		return errors.New("not authorized to update this collection")
	}

	// 3) Perform update
	_, err = repositories.UpdateCollection(collID, updates)
	return err
}

// DeleteCollectionService deletes an existing collection.
func (svc *CollectionService) DeleteCollectionService(collID, sellerID primitive.ObjectID) error {
	// 1) Fetch and verify
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop.OwnerID != sellerID {
		return errors.New("not authorized to delete this collection")
	}

	// 2) Delete
	_, err = repositories.DeleteCollection(collID)
	return err
}

// AddProductToCollectionService adds a product to a collection.
func (svc *CollectionService) AddProductToCollectionService(
	collID, prodID, sellerID primitive.ObjectID,
) error {
	// 1) Fetch collection
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	// 2) Verify shop’s owner
	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop.OwnerID != sellerID {
		return errors.New("not authorized to modify this collection")
	}

	// 3) Add product
	_, err = repositories.AddProductToCollection(collID, prodID)
	return err
}

// RemoveProductFromCollectionService removes a product from a collection.
func (svc *CollectionService) RemoveProductFromCollectionService(
	collID, prodID, sellerID primitive.ObjectID,
) error {
	// 1) Fetch collection
	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		return err
	}
	if coll == nil {
		return errors.New("collection not found")
	}

	// 2) Verify shop’s owner
	shop, err := sharedShopService.GetShopByIDService(coll.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop.OwnerID != sellerID {
		return errors.New("not authorized to modify this collection")
	}

	// 3) Remove product
	_, err = repositories.RemoveProductFromCollection(collID, prodID)
	return err
}
