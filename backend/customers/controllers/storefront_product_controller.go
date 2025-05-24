package controllers

import (
	"net/http"

	productServices "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProductsByShop returns all products for a shop.
// GET /shops/:shopid/products
func GetProductsByShop(c *gin.Context) {
	// parse shop ID
	rawShop := c.Param("shopid")
	shopOID, err := primitive.ObjectIDFromHex(rawShop)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// fetch products
	products, err := productServices.GetProductsByShopIDService(shopOID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductDetail returns one product by ID.
// GET /shops/:shopid/products/:productid
func GetProductDetail(c *gin.Context) {
	productID := c.Param("productid")
	product, err := productServices.GetProductByIDService(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
