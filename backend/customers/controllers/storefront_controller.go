
package controllers

import (
	"net/http"

	shopServices "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
)

// GetStorefront handles GET /storefront/:slug
func GetStorefront(c *gin.Context) {
    slug := c.Param("slug")

    shop, err := shopServices.GetShopBySlugService(slug)
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