package controllers

import (
	"errors"
	"net/http"
	"time"

	shopServices "github.com/Endale2/DRPS/shared/services"
	themeServices "github.com/Endale2/DRPS/shared/services"
	sharedRepositories "github.com/Endale2/DRPS/shared/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// =============== Scalable Theme Controllers ===============
// This file implements the new scalable theme architecture using separate collections

// =============== Customization API Controllers ===============

// GetCustomizationData returns all customization data for a shop using scalable theme collection
// GET /seller/shops/:shopId/customization
func GetCustomizationData(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Get shop details using existing shop service
	shop, err := shopServices.GetShopByIDService(shopID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Verify shop ownership
	if shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized: you don't own this shop"})
		return
	}

	// Get shop analytics for context
	shopObjectID, _ := primitive.ObjectIDFromHex(shopID)
	customers, _ := sharedRepositories.GetShopCustomerLinks(shopObjectID)
	
	// Get or create theme from separate collection (scalable approach)
	theme, err := themeServices.GetOrCreateShopTheme(shopObjectID, sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load theme configuration"})
		return
	}

	// Get theme performance metrics
	themeMetrics, _ := themeServices.GetThemePerformanceMetrics(shopObjectID)

	// Prepare theme customization response matching frontend expectations
	customization := gin.H{
		"colors":    theme.Colors,
		"fonts":     theme.Fonts,
		"layout":    theme.Layout,
		"customCSS": theme.CustomCSS,
		"seo":       theme.SEO,
		"mobile":    theme.Mobile,
	}

	// Prepare enhanced shop context data with theme analytics
	shopContext := gin.H{
		"shop": gin.H{
			"id":                shop.ID,
			"name":              shop.Name,
			"slug":              shop.Slug,
			"image":             shop.Image,
			"banner":            shop.Banner,
			"description":       shop.Description,
			"email":             shop.Email,
			"phone":             shop.Phone,
			"address":           shop.Address,
			"currency":          shop.Currency,
			"status":            shop.Status,
			"isVerified":        shop.IsVerified,
			"themeColor":        shop.ThemeColor,
			"themeVersion":      shop.ThemeVersion,
			"themeLastModified": shop.ThemeLastModified,
			"activeThemeId":     shop.ActiveThemeID,
			"customerCount":     len(customers),
			"productCount":      shop.ProductCount,
			"revenue":           shop.Revenue,
			"rating":            shop.Rating,
			"reviewCount":       shop.ReviewCount,
		},
		"analytics": gin.H{
			"totalCustomers": len(customers),
			"hasProducts":    shop.ProductCount > 0,
			"hasOrders":      shop.Revenue > 0,
			"shopAge":        shop.CreatedAt,
			"themeUpdated":   shop.ThemeLastModified,
		},
		"theme": gin.H{
			"id":               theme.ID,
			"name":             theme.Name,
			"version":          theme.Version,
			"isActive":         theme.IsActive,
			"usageCount":       theme.UsageCount,
			"lastUsed":         theme.LastUsedAt,
			"performance":      themeMetrics,
			"hasBackup":        theme.Backup != nil,
			"tags":             theme.Tags,
		},
	}

	response := gin.H{
		"customization": customization,
		"shop_context":  shopContext,
		"success":       true,
		"scalable":      true, // Indicates using scalable theme collection
	}

	c.JSON(http.StatusOK, response)
}

// UpdateCustomization updates theme customization for a shop using scalable theme collection
// PATCH /seller/shops/:shopId/customization
func UpdateCustomization(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Verify shop ownership first
	shop, err := shopServices.GetShopByIDService(shopID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	if shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized: you don't own this shop"})
		return
	}

	// Bind customization data - support all theme sections
	var customizationData struct {
		Colors    map[string]string `json:"colors,omitempty"`
		Fonts     map[string]string `json:"fonts,omitempty"`
		Layout    map[string]string `json:"layout,omitempty"`
		CustomCSS string            `json:"customCSS,omitempty"`
		SEO       map[string]string `json:"seo,omitempty"`
		Mobile    map[string]string `json:"mobile,omitempty"`
	}
	if err := c.ShouldBindJSON(&customizationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Validate customization data
	if err := validateEnhancedCustomizationData(customizationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update theme using scalable theme collection
	shopObjectID, _ := primitive.ObjectIDFromHex(shopID)
	updateData := make(map[string]interface{})
	
	if customizationData.Colors != nil {
		updateData["colors"] = customizationData.Colors
	}
	if customizationData.Fonts != nil {
		updateData["fonts"] = customizationData.Fonts
	}
	if customizationData.Layout != nil {
		updateData["layout"] = customizationData.Layout
	}
	if customizationData.CustomCSS != "" {
		updateData["customCSS"] = customizationData.CustomCSS
	}
	if customizationData.SEO != nil {
		updateData["seo"] = customizationData.SEO
	}
	if customizationData.Mobile != nil {
		updateData["mobile"] = customizationData.Mobile
	}

	theme, err := themeServices.UpdateShopThemeComplete(shopObjectID, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save customization"})
		return
	}

	// Get updated customer count for response
	customers, _ := sharedRepositories.GetShopCustomerLinks(shopObjectID)
	
	// Prepare response matching frontend expectations
	response := gin.H{
		"message": "customization saved successfully",
		"success": true,
		"customization": gin.H{
			"colors":    theme.Colors,
			"fonts":     theme.Fonts,
			"layout":    theme.Layout,
			"customCSS": theme.CustomCSS,
			"seo":       theme.SEO,
			"mobile":    theme.Mobile,
		},
		"theme": theme.ToConfig().Metadata,
		"shop_context": gin.H{
			"updated_at":     time.Now(),
			"customer_count": len(customers),
			"theme_modified": true,
		},
		"scalable": true,
	}

	c.JSON(http.StatusOK, response)
}

// =============== Granular Theme API Endpoints ===============

// SaveThemeColors saves only the color configuration using scalable theme collection
// PATCH /seller/shops/:shopId/customization/colors
func SaveThemeColors(c *gin.Context) {
	shopID := c.Param("shopId")
	
	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	
	var colors map[string]string
	if err := c.ShouldBindJSON(&colors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid colors data"})
		return
	}

	// Validate colors
	for _, value := range colors {
		if !isValidHexColor(value) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid color format: " + value})
			return
		}
	}

	// Verify shop ownership
	shop, err := shopServices.GetShopByIDService(shopID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	if shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized: you don't own this shop"})
		return
	}

	// Update colors in theme collection (scalable)
	shopObjectID, _ := primitive.ObjectIDFromHex(shopID)
	theme, err := themeServices.UpdateShopThemeColors(shopObjectID, colors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save colors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "colors saved successfully",
		"colors":      colors,
		"theme":       theme.ToConfig().Metadata,
		"success":     true,
		"scalable":    true,
	})
}

// SaveTypography saves only the font configuration using scalable theme collection
// PATCH /seller/shops/:shopId/customization/fonts
func SaveTypography(c *gin.Context) {
	shopID := c.Param("shopId")
	
	// Get seller ID from context and verify ownership
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	
	var fonts map[string]string
	if err := c.ShouldBindJSON(&fonts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fonts data"})
		return
	}

	// Verify shop ownership
	shop, err := shopServices.GetShopByIDService(shopID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	if shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized: you don't own this shop"})
		return
	}

	// Update fonts in theme collection (scalable)
	shopObjectID, _ := primitive.ObjectIDFromHex(shopID)
	theme, err := themeServices.UpdateShopThemeFonts(shopObjectID, fonts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save fonts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "typography saved successfully",
		"fonts":    fonts,
		"theme":    theme.ToConfig().Metadata,
		"success":  true,
		"scalable": true,
	})
}

// SaveLayout saves only the layout configuration using scalable theme collection
// PATCH /seller/shops/:shopId/customization/layout
func SaveLayout(c *gin.Context) {
	shopID := c.Param("shopId")
	
	// Get seller ID from context and verify ownership
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	
	var layout map[string]string
	if err := c.ShouldBindJSON(&layout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid layout data"})
		return
	}

	// Verify shop ownership
	shop, err := shopServices.GetShopByIDService(shopID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	if shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized: you don't own this shop"})
		return
	}

	// Update layout in theme collection (scalable)
	shopObjectID, _ := primitive.ObjectIDFromHex(shopID)
	theme, err := themeServices.UpdateShopThemeLayout(shopObjectID, layout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save layout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "layout saved successfully",
		"layout":   layout,
		"theme":    theme.ToConfig().Metadata,
		"success":  true,
		"scalable": true,
	})
}

// ResetCustomization resets all theme settings to default using scalable theme collection
// POST /seller/shops/:shopId/customization/reset
func ResetCustomization(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context and verify ownership
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Verify shop ownership
	shop, err := shopServices.GetShopByIDService(shopID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	if shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized: you don't own this shop"})
		return
	}

	// Reset theme in theme collection (scalable)
	shopObjectID, _ := primitive.ObjectIDFromHex(shopID)
	theme, err := themeServices.ResetShopTheme(shopObjectID, sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reset customization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "customization reset to default",
		"customization": gin.H{
			"colors":    theme.Colors,
			"fonts":     theme.Fonts,
			"layout":    theme.Layout,
			"customCSS": theme.CustomCSS,
			"seo":       theme.SEO,
			"mobile":    theme.Mobile,
		},
		"theme":    theme.ToConfig().Metadata,
		"success":  true,
		"scalable": true,
	})
}

// =============== Legacy Theme Controllers (Placeholders) ===============

// CreateTheme placeholder for marketplace theme creation
func CreateTheme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "marketplace theme creation disabled in scalable architecture",
		"note":    "focus on shop-specific themes for better performance",
	})
}

// GetPublicThemes placeholder for marketplace themes
func GetPublicThemes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"themes": []interface{}{}, "scalable": true})
}

// GetMyThemes placeholder for user themes
func GetMyThemes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"themes": []interface{}{}, "scalable": true})
}

// SearchThemes placeholder for theme search
func SearchThemes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"themes": []interface{}{}, "scalable": true})
}

// GetTheme placeholder for theme details
func GetTheme(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "theme not found in scalable architecture"})
}

// UpdateTheme placeholder for theme updates
func UpdateTheme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "theme update disabled in scalable architecture"})
}

// DeleteTheme placeholder for theme deletion
func DeleteTheme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "theme deletion disabled in scalable architecture"})
}

// GetThemePresets placeholder for theme presets
func GetThemePresets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"presets": []interface{}{}, "scalable": true})
}

// CreateShopTheme placeholder for shop theme creation
func CreateShopTheme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "using scalable theme collection"})
}

// GetShopThemes placeholder for shop themes
func GetShopThemes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"shop_themes": []interface{}{}, "scalable": true})
}

// GetShopActiveTheme placeholder for active shop theme
func GetShopActiveTheme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "using scalable theme collection"})
}

// UpdateShopTheme placeholder for shop theme updates
func UpdateShopTheme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "using scalable theme collection"})
}

// PublishShopTheme placeholder for theme publishing
func PublishShopTheme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "using scalable theme collection"})
}

// DeleteShopTheme placeholder for shop theme deletion
func DeleteShopTheme(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "using scalable theme collection"})
}

// ApplyThemePreset placeholder for applying presets
func ApplyThemePreset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "using scalable theme collection"})
}

// SeedDefaultThemes placeholder for theme seeding
func SeedDefaultThemes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "using scalable theme collection"})
}

// =============== Helper Functions ===============

// validateEnhancedCustomizationData validates the enhanced customization input
func validateEnhancedCustomizationData(data struct {
	Colors    map[string]string `json:"colors,omitempty"`
	Fonts     map[string]string `json:"fonts,omitempty"`
	Layout    map[string]string `json:"layout,omitempty"`
	CustomCSS string            `json:"customCSS,omitempty"`
	SEO       map[string]string `json:"seo,omitempty"`
	Mobile    map[string]string `json:"mobile,omitempty"`
}) error {
	// Validate colors
	if data.Colors != nil {
		for key, value := range data.Colors {
			if key == "" || value == "" {
				return errors.New("color configuration cannot have empty keys or values")
			}
			// Basic hex color validation
			if !isValidHexColor(value) {
				return errors.New("invalid color format: " + value)
			}
		}
	}

	// Validate fonts
	if data.Fonts != nil {
		for key, value := range data.Fonts {
			if key == "" || value == "" {
				return errors.New("font configuration cannot have empty keys or values")
			}
		}
	}

	// Validate layout
	if data.Layout != nil {
		for key, value := range data.Layout {
			if key == "" || value == "" {
				return errors.New("layout configuration cannot have empty keys or values")
			}
		}
	}

	// Validate custom CSS length
	if len(data.CustomCSS) > 100000 { // 100KB limit
		return errors.New("custom CSS too large (max 100KB)")
	}

	return nil
}

// isValidHexColor validates hex color format
func isValidHexColor(color string) bool {
	if len(color) != 7 || color[0] != '#' {
		return false
	}
	for i := 1; i < len(color); i++ {
		c := color[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}
