package services

import (
	"errors"

	"github.com/Endale2/DRPS/sellers/models"
	"github.com/Endale2/DRPS/sellers/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateSellerService creates a new seller.
func CreateSellerService(seller *models.Seller) (*mongo.InsertOneResult, error) {
	// Place for business validations (e.g., email format, uniqueness checks)
	return repositories.CreateSeller(seller)
}

// GetSellerByIDService retrieves a seller by its hex ID string.
func GetSellerByIDService(idStr string) (*models.Seller, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid seller ID")
	}
	return repositories.GetSellerByID(id)
}

// GetAllSellersService returns all sellers.
func GetAllSellersService() ([]models.Seller, error) {
	return repositories.GetAllSellers()
}

// UpdateSellerService updates fields of a seller identified by its hex ID string.
func UpdateSellerService(idStr string, updatedData bson.M) (*mongo.UpdateResult, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid seller ID")
	}
	// Place for service-level validations (e.g., allowed fields)
	return repositories.UpdateSeller(id, updatedData)
}

// DeleteSellerService removes a seller by its hex ID string.
func DeleteSellerService(idStr string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid seller ID")
	}
	return repositories.DeleteSeller(id)
}
