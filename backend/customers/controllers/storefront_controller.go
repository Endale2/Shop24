// customers/controllers/storefront_controller.go
package controllers

import (
	"net/http"

	shopServices "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
)

// GetStorefront handles GET /shops/:shopSlug
func GetStorefront(c *gin.Context) {
	shopSlug := c.Param("shopSlug")

	shop, err := shopServices.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	c.JSON(http.StatusOK, shop)
}
