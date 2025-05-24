
package controllers

import (
	"net/http"

	shopServices "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
)

// GetStorefront handles public GET /storefront/:id requests to retrieve a shop by ID.
func GetStorefront(c *gin.Context) {
	// Extract shop ID from path parameter
	id := c.Param("id")

	// Call service layer to get shop
	shop, err := shopServices.GetShopByIDService(id)
	if err != nil {
		// On error (e.g., invalid ID or DB error), return 400 or 500 accordingly
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if shop == nil {
		// No shop found for the given ID
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Return shop data
	c.JSON(http.StatusOK, shop)
}