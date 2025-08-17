package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetLayouts retrieves page layouts for a shop
// GET /seller/shops/:shopId/layouts
func GetLayouts(c *gin.Context) {
	shopID := c.Param("shopId")

	// Return default layouts for now
	layouts := []map[string]interface{}{
		{
			"id":       primitive.NewObjectID().Hex(),
			"shopId":   shopID,
			"pageType": "home",
			"name":     "Homepage Layout",
			"components": []map[string]interface{}{
				{
					"id":       primitive.NewObjectID().Hex(),
					"type":     "hero",
					"position": 1,
					"config": map[string]interface{}{
						"title":    "Welcome to Our Store",
						"subtitle": "Discover amazing products",
					},
				},
				{
					"id":       primitive.NewObjectID().Hex(),
					"type":     "product-grid",
					"position": 2,
					"config": map[string]interface{}{
						"title": "Featured Products",
						"limit": 8,
					},
				},
			},
			"isActive": true,
		},
		{
			"id":       primitive.NewObjectID().Hex(),
			"shopId":   shopID,
			"pageType": "product",
			"name":     "Product Page Layout",
			"components": []map[string]interface{}{
				{
					"id":       primitive.NewObjectID().Hex(),
					"type":     "product-details",
					"position": 1,
					"config": map[string]interface{}{
						"showReviews": true,
						"showRelated": true,
					},
				},
			},
			"isActive": true,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"layouts": layouts,
		"total":   len(layouts),
	})
}

// GetLayout retrieves a specific page layout
// GET /seller/shops/:shopId/layouts/:layoutId
func GetLayout(c *gin.Context) {
	shopID := c.Param("shopId")
	layoutID := c.Param("layoutId")

	// Return a default layout for now
	layout := map[string]interface{}{
		"id":       layoutID,
		"shopId":   shopID,
		"pageType": "home",
		"name":     "Homepage Layout",
		"components": []map[string]interface{}{
			{
				"id":       primitive.NewObjectID().Hex(),
				"type":     "hero",
				"position": 1,
				"config": map[string]interface{}{
					"title":      "Welcome to Our Store",
					"subtitle":   "Discover amazing products",
					"buttonText": "Shop Now",
					"buttonLink": "/products",
				},
			},
			{
				"id":       primitive.NewObjectID().Hex(),
				"type":     "product-grid",
				"position": 2,
				"config": map[string]interface{}{
					"title":   "Featured Products",
					"limit":   8,
					"columns": 4,
				},
			},
		},
		"isActive": true,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"layout":  layout,
	})
}

// CreateLayout creates a new page layout
// POST /seller/shops/:shopId/layouts
func CreateLayout(c *gin.Context) {
	shopID := c.Param("shopId")

	var req struct {
		PageType   string                   `json:"pageType" binding:"required"`
		Name       string                   `json:"name" binding:"required"`
		Components []map[string]interface{} `json:"components"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Create layout
	layout := map[string]interface{}{
		"id":         primitive.NewObjectID().Hex(),
		"shopId":     shopID,
		"pageType":   req.PageType,
		"name":       req.Name,
		"components": req.Components,
		"isActive":   true,
	}

	// In a full implementation, this would save to database
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"layout":  layout,
		"message": "Layout created successfully",
	})
}

// UpdateLayout updates an existing page layout
// PUT /seller/shops/:shopId/layouts/:layoutId
func UpdateLayout(c *gin.Context) {
	shopID := c.Param("shopId")
	layoutID := c.Param("layoutId")

	var req struct {
		Name       string                   `json:"name"`
		Components []map[string]interface{} `json:"components"`
		IsActive   bool                     `json:"isActive"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// In a full implementation, this would update the database
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Layout updated successfully",
		"layoutId": layoutID,
		"shopId":   shopID,
	})
}

// DeleteLayout deletes a page layout
// DELETE /seller/shops/:shopId/layouts/:layoutId
func DeleteLayout(c *gin.Context) {
	shopID := c.Param("shopId")
	layoutID := c.Param("layoutId")

	// In a full implementation, this would delete from database
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Layout deleted successfully",
		"layoutId": layoutID,
		"shopId":   shopID,
	})
}

// SaveLayoutComponents saves component positions and configurations
// POST /seller/shops/:shopId/layouts/:layoutId/components
func SaveLayoutComponents(c *gin.Context) {
	shopID := c.Param("shopId")
	layoutID := c.Param("layoutId")

	var req struct {
		Components []map[string]interface{} `json:"components" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// In a full implementation, this would update component positions in database
	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"message":        "Layout components saved successfully",
		"layoutId":       layoutID,
		"shopId":         shopID,
		"componentCount": len(req.Components),
	})
}

// GetPageTypes returns available page types for layouts
// GET /seller/layouts/page-types
func GetPageTypes(c *gin.Context) {
	pageTypes := []map[string]interface{}{
		{"id": "home", "name": "Homepage", "description": "Main landing page"},
		{"id": "product", "name": "Product Page", "description": "Individual product details"},
		{"id": "collection", "name": "Collection Page", "description": "Product category/collection"},
		{"id": "about", "name": "About Page", "description": "About us page"},
		{"id": "contact", "name": "Contact Page", "description": "Contact information"},
		{"id": "blog", "name": "Blog Page", "description": "Blog listing page"},
		{"id": "cart", "name": "Cart Page", "description": "Shopping cart"},
		{"id": "checkout", "name": "Checkout Page", "description": "Checkout process"},
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"pageTypes": pageTypes,
	})
}

// PreviewLayout generates a preview of the layout
// GET /seller/shops/:shopId/layouts/:layoutId/preview
func PreviewLayout(c *gin.Context) {
	shopID := c.Param("shopId")
	layoutID := c.Param("layoutId")

	// Generate preview HTML/CSS for the layout
	preview := map[string]interface{}{
		"html": "<div class='layout-preview'>Layout preview content</div>",
		"css":  ".layout-preview { padding: 20px; background: #f5f5f5; }",
		"metadata": map[string]interface{}{
			"layoutId":    layoutID,
			"shopId":      shopID,
			"generatedAt": "2024-01-01T00:00:00Z",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"preview": preview,
	})
}
