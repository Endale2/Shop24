package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/Endale2/DRPS/shared/models"
	categoryService "github.com/Endale2/DRPS/shared/services"
	shopService "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateShop lets a seller create a new shop owned by them.
// POST /seller/shops
func CreateShop(c *gin.Context) {
	// 1) Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// 2) Bind incoming JSON
	var shop models.Shop
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// 3) Set owner and validate required fields
	shop.OwnerID = sellerID

	// Validate and sanitize shop name
	if strings.TrimSpace(shop.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shop name is required"})
		return
	}
	shop.Name = strings.TrimSpace(shop.Name)
	if len(shop.Name) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shop name must be at least 3 characters"})
		return
	}
	if len(shop.Name) > 50 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shop name must be less than 50 characters"})
		return
	}

	// Validate and sanitize description
	if shop.Description != "" {
		shop.Description = strings.TrimSpace(shop.Description)
		if len(shop.Description) > 200 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "description must be less than 200 characters"})
			return
		}
	}

	// Validate category ID
	if shop.CategoryID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shop category is required"})
		return
	}

	// Verify that the category exists
	category, err := categoryService.GetShopCategoryByIDService(shop.CategoryID.Hex())
	if err != nil || category == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop category"})
		return
	}

	// Validate and sanitize email
	if strings.TrimSpace(shop.Email) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contact email is required"})
		return
	}
	shop.Email = strings.TrimSpace(shop.Email)
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if !emailRegex.MatchString(shop.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please enter a valid email address"})
		return
	}

	// Set default values
	if shop.Currency == "" {
		shop.Currency = "USD"
	}

	if shop.Status == "" {
		shop.Status = "inactive"
	}

	// Initialize analytics fields
	shop.CustomerCount = 0
	shop.ProductCount = 0
	shop.Revenue = 0
	shop.Rating = 0
	shop.ReviewCount = 0
	shop.IsVerified = false

	// 4) Delegate to service
	res, err := shopService.CreateShopService(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create shop"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetShops returns all shops owned by the current seller.
// GET /seller/shops
func GetShops(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	all, err := shopService.GetAllShopsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list shops"})
		return
	}

    // Ensure we return an empty array [] instead of null when no shops exist
    mine := make([]models.Shop, 0)
	for _, s := range all {
		if s.OwnerID == sellerID {
			// Add some default metadata for frontend display
			if s.Image == "" {
				s.Image = "" // Keep empty for frontend to show placeholder
			}
			mine = append(mine, s)
		}
	}
	c.JSON(http.StatusOK, mine)
}

// GetShop returns one shop if owned by the seller.
// GET /seller/shops/:shopId
func GetShop(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// use lowercase "shopId" exactly as in your route
	shopID := c.Param("shopId")
	if shopID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shopId is required"})
		return
	}

	shop, err := shopService.GetShopByIDService(shopID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	if shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your shop"})
		return
	}

	c.JSON(http.StatusOK, shop)
}

// UpdateShop allows a seller to update their own shop.
// PATCH /seller/shops/:shopId
func UpdateShop(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	shopID := c.Param("shopId")

	shop, err := shopService.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Validate updates if they contain required fields
	if name, exists := updates["name"]; exists {
		if strings.TrimSpace(name.(string)) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "shop name cannot be empty"})
			return
		}
	}

	if category, exists := updates["category"]; exists {
		if strings.TrimSpace(category.(string)) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "shop category cannot be empty"})
			return
		}
	}

	if email, exists := updates["email"]; exists {
		if strings.TrimSpace(email.(string)) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "contact email cannot be empty"})
			return
		}
	}

	res, err := shopService.UpdateShopService(shopID, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteShop deletes a shop owned by the seller.
// DELETE /seller/shops/:id
func DeleteShop(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	shopID := c.Param("shopId")

	shop, err := shopService.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	res, err := shopService.DeleteShopService(shopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetShopCategories returns all available shop categories.
// GET /seller/categories
func GetShopCategories(c *gin.Context) {
	categories, err := categoryService.GetAllShopCategoriesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}
