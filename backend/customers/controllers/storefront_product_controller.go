package controllers

import (
	"net/http"

	shopService "github.com/Endale2/DRPS/shared/services"
	productServices "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
)

// GetProductsByShop returns all products for a shop by slug.
// GET /shops/:slug/products
func GetProductsByShop(c *gin.Context) {
	slug := c.Param("slug")
	shop, err := shopService.GetShopBySlugService(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	products, err := productServices.GetProductsByShopSlugService(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}
	if products == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no products found"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductDetail returns one product by its slug.
// GET /shops/:slug/products/:productslug
func GetProductDetail(c *gin.Context) {
	// extract shopSlug to ensure nested route, though not used to fetch product
	// you could verify that product belongs to shop if needed
	productSlug := c.Param("productslug")
	product, err := productServices.GetProductBySlugService(productSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve product"})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
