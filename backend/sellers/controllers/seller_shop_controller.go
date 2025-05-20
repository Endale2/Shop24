package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
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
	shop.OwnerID = sellerID

	// 3) Delegate to service
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

	var mine []models.Shop
	for _, s := range all {
		if s.OwnerID == sellerID {
			mine = append(mine, s)
		}
	}
	c.JSON(http.StatusOK, mine)
}

// GetShop returns one shop if owned by the seller.
// GET /seller/shops/:id
func GetShop(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopID := c.Param("id")
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
// PATCH /seller/shops/:id
func UpdateShop(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	shopID := c.Param("id")

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
	shopID := c.Param("id")

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
