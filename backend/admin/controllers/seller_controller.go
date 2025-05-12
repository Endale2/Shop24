package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/sellers/models"
	"github.com/Endale2/DRPS/sellers/services"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateSeller handles creation of a new Seller (admin).
func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := services.CreateSellerService(&seller)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating seller"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetSeller retrieves a single Seller by its ID.
func GetSeller(c *gin.Context) {
	id := c.Param("id")

	seller, err := services.GetSellerByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	c.JSON(http.StatusOK, seller)
}

// GetSellers retrieves all Sellers.
func GetSellers(c *gin.Context) {
	sellers, err := services.GetAllSellersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving sellers"})
		return
	}
	c.JSON(http.StatusOK, sellers)
}

// UpdateSeller updates an existing Seller by its ID.
func UpdateSeller(c *gin.Context) {
	id := c.Param("id")

	var updatedData map[string]interface{}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := services.UpdateSellerService(id, bson.M(updatedData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating seller"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteSeller deletes a Seller by its ID.
func DeleteSeller(c *gin.Context) {
	id := c.Param("id")

	result, err := services.DeleteSellerService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting seller"})
		return
	}
	c.JSON(http.StatusOK, result)
}
