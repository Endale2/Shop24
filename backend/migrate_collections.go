package main

import (
	"context"
	"log"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Connect to MongoDB
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Get the shop_categories collection
	shopCategoriesCollection := config.GetCollection("DRPS", "shop_categories")

	// Define default categories
	defaultCategories := []models.ShopCategory{
		{
			Name:        "Fashion & Apparel",
			Slug:        "fashion-apparel",
			Description: "Clothing, shoes, accessories, and fashion items",
		},
		{
			Name:        "Electronics",
			Slug:        "electronics",
			Description: "Computers, phones, gadgets, and electronic devices",
		},
		{
			Name:        "Home & Garden",
			Slug:        "home-garden",
			Description: "Furniture, decor, gardening, and home improvement",
		},
		{
			Name:        "Beauty & Personal Care",
			Slug:        "beauty-personal-care",
			Description: "Cosmetics, skincare, haircare, and personal hygiene",
		},
		{
			Name:        "Sports & Outdoors",
			Slug:        "sports-outdoors",
			Description: "Sports equipment, outdoor gear, and fitness items",
		},
		{
			Name:        "Books & Media",
			Slug:        "books-media",
			Description: "Books, magazines, music, movies, and digital media",
		},
		{
			Name:        "Toys & Games",
			Slug:        "toys-games",
			Description: "Toys, games, puzzles, and entertainment items",
		},
		{
			Name:        "Food & Beverages",
			Slug:        "food-beverages",
			Description: "Food, drinks, snacks, and culinary items",
		},
		{
			Name:        "Health & Wellness",
			Slug:        "health-wellness",
			Description: "Health products, supplements, and wellness items",
		},
		{
			Name:        "Automotive",
			Slug:        "automotive",
			Description: "Car parts, accessories, and automotive products",
		},
		{
			Name:        "Other",
			Slug:        "other",
			Description: "Miscellaneous items and general merchandise",
		},
	}

	// Check if categories already exist
	count, err := shopCategoriesCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("Failed to count categories: %v", err)
	}

	if count > 0 {
		log.Println("Shop categories already exist, skipping migration")
		return
	}

	// Insert default categories
	var categoriesToInsert []interface{}
	for _, category := range defaultCategories {
		category.ID = primitive.NewObjectID()
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()
		categoriesToInsert = append(categoriesToInsert, category)
	}

	result, err := shopCategoriesCollection.InsertMany(context.Background(), categoriesToInsert)
	if err != nil {
		log.Fatalf("Failed to insert categories: %v", err)
	}

	log.Printf("Successfully inserted %d shop categories", len(result.InsertedIDs))
}
