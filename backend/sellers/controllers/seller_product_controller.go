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
func CreateProduct(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopID := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	p.ShopID = shop.ID
	p.UserID = sellerID

	res, err := services.CreateProductService(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creation failed"})
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
