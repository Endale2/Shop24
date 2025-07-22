package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	// 1) Extract user_id from context
	uidVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	uidHex, ok := uidVal.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID in context"})
		return
	}
	userID, err := primitive.ObjectIDFromHex(uidHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed user ID"})
		return
	}

	// 2) Bind the rest of the product JSON
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// 3) Override the UserID field
	product.UserID = userID

	// 4) Delegate to service/repo
	result, err := services.CreateProductService(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
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

// GetProductsByCollectionCount handles GET /admin/products/count-by-category
func GetProductsByCollectionCount(c *gin.Context) {
	// Example: count products per collection
	pipeline := []bson.M{
		{"$group": bson.M{"_id": "$collection_id", "count": bson.M{"$sum": 1}}},
	}
	coll := config.GetCollection("DRPS", "products")
	cursor, err := coll.Aggregate(c, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error computing analytics"})
		return
	}
	var results []bson.M
	if err := cursor.All(c, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading results"})
		return
	}
	c.JSON(http.StatusOK, results)
}
