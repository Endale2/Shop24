package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateProduct handles the creation of a new product.(for  admin)
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

// GetProduct retrieves a single product by its ID.(for admin)
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



// GetProductCount handles GET /admin/products/count?shop_id=xxx
func GetProductCount(c *gin.Context) {
    // Build filter from query parameters
    filter := bson.M{}
    if shopID := c.Query("shop_id"); shopID != "" {
        if obj, err := primitive.ObjectIDFromHex(shopID); err == nil {
            filter["shop_id"] = obj
        }
    }
    count, err := services.CountProductsService(filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting products"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"count": count})
}

// GetProductsByCategoryCount handles GET /admin/products/count-by-category
func GetProductsByCategoryCount(c *gin.Context) {
    data, err := services.CountByCategoryService()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error computing analytics"})
        return
    }
    c.JSON(http.StatusOK, data)
}