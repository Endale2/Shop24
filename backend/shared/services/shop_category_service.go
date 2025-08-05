package services

import (
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateShopCategoryService creates a new ShopCategory.
func CreateShopCategoryService(category *models.ShopCategory) (*mongo.InsertOneResult, error) {
	return repositories.CreateShopCategory(category)
}

// GetShopCategoryByIDService retrieves a ShopCategory by its ID.
func GetShopCategoryByIDService(id string) (*models.ShopCategory, error) {
	return repositories.GetShopCategoryByID(id)
}

// GetAllShopCategoriesService returns all ShopCategories.
func GetAllShopCategoriesService() ([]models.ShopCategory, error) {
	return repositories.GetAllShopCategories()
}

// GetShopCategoryBySlugService retrieves a ShopCategory by its slug.
func GetShopCategoryBySlugService(slug string) (*models.ShopCategory, error) {
	return repositories.GetShopCategoryBySlug(slug)
}

// UpdateShopCategoryService updates fields of a ShopCategory identified by its ID.
func UpdateShopCategoryService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	return repositories.UpdateShopCategory(id, updatedData)
}

// DeleteShopCategoryService removes a ShopCategory by its ID.
func DeleteShopCategoryService(id string) (*mongo.DeleteResult, error) {
	return repositories.DeleteShopCategory(id)
}
