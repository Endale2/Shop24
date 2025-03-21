package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateShop handles the creation of a new Shop.
func CreateShop(c *gin.Context) {
	var Shop models.Shop
	if err := c.ShouldBindJSON(&Shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := services.CreateShopService(&Shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating Shop"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetShop retrieves a single Shop by its ID.
func GetShop(c *gin.Context) {
	id := c.Param("id")

	shop, err := services.GetShopByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shop not found"})
		return
	}
	c.JSON(http.StatusOK, shop)
}

// GetShops retrieves all Shops.
func GetShops(c *gin.Context) {
	shops, err := services.GetAllShopsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Shops"})
		return
	}
	c.JSON(http.StatusOK, shops)
}

// UpdateShop updates an existing Shop identified by its ID.
func UpdateShop(c *gin.Context) {
	id := c.Param("id")

	// Decode the updated fields from the request body.
	var updatedData map[string]interface{}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := services.UpdateShopService(id, bson.M(updatedData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating Shop"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteShop deletes a Shop by its ID.
func DeleteShop(c *gin.Context) {
	id := c.Param("id")

	result, err := services.DeleteShopService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Shop"})
		return
	}
	c.JSON(http.StatusOK, result)
}
