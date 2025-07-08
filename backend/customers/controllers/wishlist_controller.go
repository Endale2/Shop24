package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET /shops/:shopSlug/wishlist
func GetWishlist(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	cidVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	customerID, err := primitive.ObjectIDFromHex(cidVal.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid customer ID"})
		return
	}
	wishlist, err := services.GetWishlistService(shop.ID, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if wishlist == nil {
		c.JSON(http.StatusOK, gin.H{"product_ids": []string{}})
		return
	}
	ids := make([]string, len(wishlist.ProductIDs))
	for i, id := range wishlist.ProductIDs {
		ids[i] = id.Hex()
	}
	c.JSON(http.StatusOK, gin.H{"product_ids": ids})
}

// POST /shops/:shopSlug/wishlist { product_id }
func AddToWishlist(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	cidVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	customerID, err := primitive.ObjectIDFromHex(cidVal.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid customer ID"})
		return
	}
	var req struct {
		ProductID string `json:"product_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	wishlist, err := services.AddProductToWishlistService(shop.ID, customerID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ids := make([]string, len(wishlist.ProductIDs))
	for i, id := range wishlist.ProductIDs {
		ids[i] = id.Hex()
	}
	c.JSON(http.StatusOK, gin.H{"product_ids": ids})
}

// DELETE /shops/:shopSlug/wishlist/:productId
func RemoveFromWishlist(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	cidVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	customerID, err := primitive.ObjectIDFromHex(cidVal.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid customer ID"})
		return
	}
	productID, err := primitive.ObjectIDFromHex(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	err = services.RemoveProductFromWishlistService(shop.ID, customerID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
