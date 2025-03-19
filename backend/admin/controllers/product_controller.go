package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateProduct handles the creation of a new product.
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := services.CreateProductService(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating product"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetProduct retrieves a single product by its ID.
func GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := services.GetProductByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// GetProducts retrieves all products.
func GetProducts(c *gin.Context) {
	products, err := services.GetAllProductsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// UpdateProduct updates an existing product identified by its ID.
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	// Decode the updated fields from the request body.
	var updatedData map[string]interface{}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := services.UpdateProductService(id, bson.M(updatedData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteProduct deletes a product by its ID.
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	result, err := services.DeleteProductService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting product"})
		return
	}
	c.JSON(http.StatusOK, result)
}
