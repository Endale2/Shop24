package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProduct creates a product under a seller’s shop.
// POST /seller/shops/:shopId/products
// variantInput matches your payload's structure:
// {
//   "name": "Size",
//   "options": ["S","M","L"],
//   "prices": [49.99,49.99,49.99]
// }
type variantInput struct {
	Name    string    `json:"name" binding:"required"`
	Options []string  `json:"options" binding:"required"`
	Prices  []float64 `json:"prices" binding:"required"`
}

// createProductInput matches the rest of your payload
type createProductInput struct {
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Images      []string       `json:"images" binding:"required"`
	Category    string         `json:"category" binding:"required"`
	Price       float64        `json:"price" binding:"required"`
	CreatedBy   string         `json:"createdBy" binding:"required"`
	Variants    []variantInput `json:"variants"`
}

// CreateProduct handles POST /seller/shops/:shopId/products
func CreateProduct(c *gin.Context) {
	// 1) get seller ID from context
	sellerHex, _ := c.Get("user_id")
	sellerID, err := primitive.ObjectIDFromHex(sellerHex.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

	// 2) parse shop ID from URL
	shopHex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shopId"})
		return
	}

	// 3) verify shop ownership
	shop, err := services.GetShopByIDService(shopHex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	// 4) bind JSON into our input struct
	var input createProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	// 5) map input → models.Product
	prod := &models.Product{
		ShopID:      shopID,
		UserID:      sellerID,
		Name:        input.Name,
		Description: input.Description,
		Images:      input.Images,
		Category:    input.Category,
		Price:       input.Price,
		CreatedBy:   sellerID,
		
	}

	// 6) build Variant slice
	for _, v := range input.Variants {
		// ensure options/prices lengths match
		if len(v.Options) != len(v.Prices) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "variant '" + v.Name + "' options/prices length mismatch"})
			return
		}
		for i, opt := range v.Options {
			// each entry becomes its own Variant, keyed by variant name
			prod.Variants = append(prod.Variants, models.Variant{
				Options: map[string]string{v.Name: opt},
				Price:   v.Prices[i],
				Stock:   0,
				Image:   "",
			})
		}
	}

	// 7) persist
	res, err := services.CreateProductService(prod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creation failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProducts lists products for one of the seller’s shops.
// GET /seller/shops/:shopId/products
func GetProducts(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopID := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	all, err := services.GetAllProductsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fetch failed"})
		return
	}

	var filtered []models.Product
	for _, pr := range all {
		if pr.ShopID == shop.ID {
			filtered = append(filtered, pr)
		}
	}
	c.JSON(http.StatusOK, filtered)
}

// GetProduct retrieves one product in a seller’s shop.
// GET /seller/shops/:shopId/products/:productId
func GetProduct(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopID := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	productId := c.Param("productId")
	pr, err := services.GetProductByIDService(productId)
	if err != nil || pr.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, pr)
}

// UpdateProduct updates a product in the seller’s shop.
// PATCH /seller/shops/:shopId/products/:productId
func UpdateProduct(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopID := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	productId := c.Param("productId")
	// Optional: verify ownership by loading product:
	// pr, _ := services.GetProductByIDService(productId)
	// if pr.ShopID != shop.ID { ... }

	var updates bson.M
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	updates["shop_id"] = shop.ID

	res, err := services.UpdateProductService(productId, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteProduct deletes a product from a seller’s shop.
// DELETE /seller/shops/:shopId/products/:productId
func DeleteProduct(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopID := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	productId := c.Param("productId")
	res, err := services.DeleteProductService(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}
