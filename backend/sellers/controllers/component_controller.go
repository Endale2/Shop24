package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Endale2/DRPS/shared/models"
)

// GetComponents retrieves available component templates
// GET /seller/components
func GetComponents(c *gin.Context) {
	// Return default component templates for now
	components := models.GetDefaultComponentTemplates()

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"components": components,
		"pagination": gin.H{
			"page":       1,
			"limit":      50,
			"total":      len(components),
			"totalPages": 1,
		},
	})
}

// GetComponent retrieves a specific component template
// GET /seller/components/:id
func GetComponent(c *gin.Context) {
	componentID := c.Param("id")

	// For now, return a default component
	// In a full implementation, this would query the database
	component := &models.ComponentTemplate{
		ID:          primitive.NewObjectID(),
		Type:        "hero",
		Name:        "Hero Banner",
		Description: "A customizable hero banner component",
		Category:    "layout",
		Settings: map[string]interface{}{
			"title":      "Welcome to Our Store",
			"subtitle":   "Discover amazing products",
			"buttonText": "Shop Now",
			"buttonLink": "/products",
			"image":      "",
			"alignment":  "center",
			"height":     "400px",
			"background": "#f8f9fa",
		},
		StyleConfig: map[string]interface{}{
			"className": "hero-banner",
			"style": map[string]interface{}{
				"minHeight":      "400px",
				"display":        "flex",
				"alignItems":     "center",
				"justifyContent": "center",
			},
		},
		IsPublic: true,
		Tags:     []string{"hero", "banner", "layout"},
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"component":   component,
		"componentId": componentID,
	})
}

// CreateComponent creates a new custom component template
// POST /seller/components
func CreateComponent(c *gin.Context) {
	var req struct {
		Name        string                 `json:"name" binding:"required"`
		Type        string                 `json:"type" binding:"required"`
		Description string                 `json:"description"`
		Category    string                 `json:"category" binding:"required"`
		Settings    map[string]interface{} `json:"settings"`
		StyleConfig map[string]interface{} `json:"styleConfig"`
		Content     map[string]interface{} `json:"content"`
		Tags        []string               `json:"tags"`
		IsPublic    bool                   `json:"isPublic"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Create component template
	component := models.ComponentTemplate{
		ID:          primitive.NewObjectID(),
		Type:        req.Type,
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Settings:    req.Settings,
		StyleConfig: req.StyleConfig,
		Content:     req.Content,
		IsPublic:    req.IsPublic,
		Tags:        req.Tags,
		UsageCount:  0,
	}

	// In a full implementation, this would save to database
	// For now, just return success
	c.JSON(http.StatusCreated, gin.H{
		"success":   true,
		"component": component,
		"message":   "Component created successfully",
	})
}

// UpdateComponent updates an existing component template
// PUT /seller/components/:id
func UpdateComponent(c *gin.Context) {
	componentID := c.Param("id")

	var req struct {
		Name        string                 `json:"name"`
		Description string                 `json:"description"`
		Category    string                 `json:"category"`
		Settings    map[string]interface{} `json:"settings"`
		StyleConfig map[string]interface{} `json:"styleConfig"`
		Content     map[string]interface{} `json:"content"`
		Tags        []string               `json:"tags"`
		IsPublic    bool                   `json:"isPublic"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// In a full implementation, this would update the database
	// For now, just return success
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     "Component updated successfully",
		"componentId": componentID,
	})
}

// DeleteComponent deletes a component template
// DELETE /seller/components/:id
func DeleteComponent(c *gin.Context) {
	componentID := c.Param("id")

	// In a full implementation, this would delete from database
	// For now, just return success
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     "Component deleted successfully",
		"componentId": componentID,
	})
}

// GetComponentCategories returns available component categories
// GET /seller/components/categories
func GetComponentCategories(c *gin.Context) {
	categories := []map[string]interface{}{
		{"id": "layout", "name": "Layout", "description": "Headers, footers, sections"},
		{"id": "content", "name": "Content", "description": "Text, images, videos"},
		{"id": "interactive", "name": "Interactive", "description": "Forms, buttons, modals"},
		{"id": "marketing", "name": "Marketing", "description": "Testimonials, CTAs, banners"},
		{"id": "ecommerce", "name": "E-commerce", "description": "Product grids, carts, checkout"},
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"categories": categories,
	})
}

// GetComponentsByCategory returns components filtered by category
// GET /seller/components/category/:category
func GetComponentsByCategory(c *gin.Context) {
	category := c.Param("category")

	// Get all components and filter by category
	allComponents := models.GetDefaultComponentTemplates()
	var filteredComponents []models.ComponentTemplate

	for _, component := range allComponents {
		if component.Category == category {
			filteredComponents = append(filteredComponents, component)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"components": filteredComponents,
		"category":   category,
		"total":      len(filteredComponents),
	})
}

// SearchComponents searches components by name or tags
// GET /seller/components/search?q=query
func SearchComponents(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	// Get all components and filter by search query
	allComponents := models.GetDefaultComponentTemplates()
	var results []models.ComponentTemplate

	queryLower := strings.ToLower(query)
	for _, component := range allComponents {
		// Simple search in name, description, and tags
		if strings.Contains(strings.ToLower(component.Name), queryLower) ||
			strings.Contains(strings.ToLower(component.Description), queryLower) ||
			containsInTags(component.Tags, queryLower) {
			results = append(results, component)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"components": results,
		"query":      query,
		"total":      len(results),
	})
}

// Helper function
func containsInTags(tags []string, query string) bool {
	for _, tag := range tags {
		if strings.Contains(strings.ToLower(tag), query) {
			return true
		}
	}
	return false
}
