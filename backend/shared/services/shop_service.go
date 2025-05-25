package services

import (
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateShopService creates a new Shop.
func CreateShopService(shop *models.Shop) (*mongo.InsertOneResult, error) {
	// You can add business logic or validations here if needed.
	return repositories.CreateShop(shop)
}

// GetShopByIDService retrieves a Shop by its ID.
func GetShopByIDService(id string) (*models.Shop, error) {
	return repositories.GetShopByID(id)
}

// GetAllShopsService returns all Shops.
func GetAllShopsService() ([]models.Shop, error) {
	return repositories.GetAllShops()
}

// UpdateShopService updates fields of a Shop identified by its ID.
func UpdateShopService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	// Add any service-level validations or business rules here.
	return repositories.UpdateShop(id, updatedData)
}

// DeleteShopService removes a Shop by its ID.
func DeleteShopService(id string) (*mongo.DeleteResult, error) {
	return repositories.DeleteShop(id)
}


// GetShopBySlugService retrieves a Shop by its slug.
func GetShopBySlugService(slug string) (*models.Shop, error) {
    return repositories.GetShopBySlug(slug)
}
