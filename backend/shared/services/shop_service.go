package services

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// generateSlug creates a URL-friendly slug from a shop name
func generateSlug(name string) string {
	// Convert to lowercase and replace spaces with dashes
	slug := strings.ToLower(name)
	slug = regexp.MustCompile(`[^a-z0-9\s-]`).ReplaceAllString(slug, "")
	slug = regexp.MustCompile(`[\s-]+`).ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}

// generateUniqueSlug ensures the slug is unique by appending a number if needed
func generateUniqueSlug(name string) (string, error) {
	baseSlug := generateSlug(name)
	slug := baseSlug
	counter := 1

	// Check if slug exists, if so append a number
	for {
		existingShop, err := GetShopBySlugService(slug)
		if err != nil {
			return "", err
		}
		if existingShop == nil {
			break // Slug is unique
		}
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}

	return slug, nil
}

// CreateShopService creates a new Shop with a unique slug.
func CreateShopService(shop *models.Shop) (*mongo.InsertOneResult, error) {
	// Generate unique slug
	slug, err := generateUniqueSlug(shop.Name)
	if err != nil {
		return nil, err
	}
	shop.Slug = slug

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

// Notes:

// Correct and simple.

// Suggestion: If you want to enforce "unique slug" at creation, check GetShopBySlug first.
