package services

import (
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateShopCategoryService(category *models.ShopCategory) (*mongo.InsertOneResult, error) {
	return repositories.CreateShopCategory(category)
}

func GetShopCategoryByIDService(id string) (*models.ShopCategory, error) {
	return repositories.GetShopCategoryByID(id)
}

func GetAllShopCategoriesService() ([]models.ShopCategory, error) {
	return repositories.GetAllShopCategories()
}

func UpdateShopCategoryService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	return repositories.UpdateShopCategory(id, updatedData)
}

func DeleteShopCategoryService(id string) (*mongo.DeleteResult, error) {
	return repositories.DeleteShopCategory(id)
}

func GetShopCategoriesByIDsService(ids []primitive.ObjectID) ([]models.ShopCategory, error) {
	return repositories.GetShopCategoriesByIDs(ids)
}
