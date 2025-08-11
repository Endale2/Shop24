package controllers

import (
	"net/http"
	"strconv"

	"github.com/Endale2/DRPS/sellers/services"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// =============== Theme Base Controllers ===============

// CreateTheme creates a new theme
// POST /seller/themes
func CreateTheme(c *gin.Context) {
	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Bind request data
	var theme models.Theme
	if err := c.ShouldBindJSON(&theme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Create theme
	result, err := services.CreateThemeService(&theme, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "theme created successfully",
		"theme_id": result.InsertedID,
	})
}

// GetTheme retrieves a specific theme
// GET /seller/themes/:themeId
func GetTheme(c *gin.Context) {
	themeID := c.Param("themeId")

	theme, err := services.GetThemeByIDService(themeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"theme": theme})
}

// GetPublicThemes retrieves all marketplace themes
// GET /seller/themes/marketplace
func GetPublicThemes(c *gin.Context) {
	themes, err := services.GetPublicThemesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch themes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"themes": themes})
}

// GetMyThemes retrieves themes created by the current seller
// GET /seller/themes/my-themes
func GetMyThemes(c *gin.Context) {
	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	themes, err := services.GetThemesByCreatorService(sellerHex.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch themes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"themes": themes})
}

// UpdateTheme updates a theme
// PATCH /seller/themes/:themeId
func UpdateTheme(c *gin.Context) {
	themeID := c.Param("themeId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Bind update data
	var updateData bson.M
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Update theme
	result, err := services.UpdateThemeService(themeID, updateData, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "theme not found or no changes made"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "theme updated successfully"})
}

// DeleteTheme deletes a theme
// DELETE /seller/themes/:themeId
func DeleteTheme(c *gin.Context) {
	themeID := c.Param("themeId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Delete theme
	result, err := services.DeleteThemeService(themeID, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "theme not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "theme deleted successfully"})
}

// SearchThemes searches themes by query
// GET /seller/themes/search?q=query&public=true
func SearchThemes(c *gin.Context) {
	query := c.Query("q")
	publicStr := c.DefaultQuery("public", "true")
	isPublic, _ := strconv.ParseBool(publicStr)

	themes, err := services.SearchThemesService(query, isPublic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search themes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"themes": themes})
}

// =============== Shop Theme Controllers ===============

// CreateShopTheme creates a new shop theme configuration
// POST /seller/shops/:shopId/themes
func CreateShopTheme(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Bind request data
	var shopTheme models.ShopTheme
	if err := c.ShouldBindJSON(&shopTheme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Set shop ID
	shopObjectID, err := primitive.ObjectIDFromHex(shopID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}
	shopTheme.ShopID = shopObjectID

	// Create shop theme
	result, err := services.CreateShopThemeService(&shopTheme, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":        "shop theme created successfully",
		"shop_theme_id": result.InsertedID,
	})
}

// GetShopActiveTheme retrieves the active theme for a shop
// GET /seller/shops/:shopId/themes/active
func GetShopActiveTheme(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Get active theme
	shopTheme, theme, err := services.GetShopActiveThemeService(shopID, sellerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shop_theme": shopTheme,
		"theme":      theme,
	})
}

// GetShopThemes retrieves all theme configurations for a shop
// GET /seller/shops/:shopId/themes
func GetShopThemes(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Get shop themes
	shopThemes, err := services.GetShopThemesService(shopID, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shop_themes": shopThemes})
}

// UpdateShopTheme updates a shop theme configuration
// PATCH /seller/shops/:shopId/themes/:shopThemeId
func UpdateShopTheme(c *gin.Context) {
	shopThemeID := c.Param("shopThemeId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Bind update data
	var updateData bson.M
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Update shop theme
	result, err := services.UpdateShopThemeService(shopThemeID, updateData, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop theme not found or no changes made"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "shop theme updated successfully"})
}

// PublishShopTheme activates a theme for a shop
// POST /seller/shops/:shopId/themes/:shopThemeId/publish
func PublishShopTheme(c *gin.Context) {
	shopID := c.Param("shopId")
	shopThemeID := c.Param("shopThemeId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Publish theme
	err := services.PublishShopThemeService(shopID, shopThemeID, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "theme published successfully"})
}

// DeleteShopTheme deletes a shop theme configuration
// DELETE /seller/shops/:shopId/themes/:shopThemeId
func DeleteShopTheme(c *gin.Context) {
	shopThemeID := c.Param("shopThemeId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Delete shop theme
	result, err := services.DeleteShopThemeService(shopThemeID, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop theme not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "shop theme deleted successfully"})
}

// =============== Theme Preset Controllers ===============

// GetThemePresets retrieves presets for a theme
// GET /seller/themes/:themeId/presets
func GetThemePresets(c *gin.Context) {
	themeID := c.Param("themeId")

	presets, err := services.GetThemePresetsService(themeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"presets": presets})
}

// ApplyThemePreset applies a preset to a shop theme
// POST /seller/shops/:shopId/themes/:shopThemeId/presets/:presetId/apply
func ApplyThemePreset(c *gin.Context) {
	shopThemeID := c.Param("shopThemeId")
	presetID := c.Param("presetId")

	// Get seller ID from context
	sellerHex, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// Apply preset
	result, err := services.ApplyThemePresetService(shopThemeID, presetID, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "preset not applied"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "preset applied successfully"})
}

// =============== Customization API Controllers ===============

// GetCustomizationData returns all customization data for a shop
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

	// Get active theme
	shopTheme, theme, err := services.GetShopActiveThemeService(shopID, sellerID)
	if err != nil {
		// If no active theme, return default data
		c.JSON(http.StatusOK, gin.H{
			"customization": gin.H{
				"colors": gin.H{
					"primary":     "#10B981",
					"secondary":   "#F59E0B",
					"background":  "#FFFFFF",
					"heading":     "#1F2937",
					"bodyText":    "#6B7280",
				},
				"fonts": gin.H{
					"heading": "Inter",
					"body":    "Inter",
				},
				"layout": gin.H{
					"headerStyle":      "classic",
					"containerWidth":   "boxed",
					"sidebarPosition":  "none",
					"gridColumns":      "3",
				},
			},
			"message": "using default theme",
		})
		return
	}

	// Merge base theme config with shop overrides
	customization := mergeThemeConfig(theme.Config, shopTheme.Overrides)

	c.JSON(http.StatusOK, gin.H{
		"customization": customization,
		"shop_theme":    shopTheme,
		"base_theme":    theme,
	})
}

// UpdateCustomization updates theme customization for a shop
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

	// Bind customization data
	var customizationData struct {
		Colors map[string]string `json:"colors,omitempty"`
		Fonts  map[string]string `json:"fonts,omitempty"`
		Layout map[string]string `json:"layout,omitempty"`
	}
	if err := c.ShouldBindJSON(&customizationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Get or create shop theme
	shopObjectID, err := primitive.ObjectIDFromHex(shopID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	shopTheme, _, err := services.GetShopActiveThemeService(shopID, sellerID)
	if err != nil {
		// Create default shop theme if none exists
		defaultTheme := &models.ShopTheme{
			ShopID:  shopObjectID,
			ThemeID: primitive.NewObjectID(), // Default theme ID
			Overrides: models.ThemeConfig{
				Colors: customizationData.Colors,
				Fonts:  customizationData.Fonts,
				Layouts: customizationData.Layout,
			},
			Published: true,
		}

		_, err := services.CreateShopThemeService(defaultTheme, sellerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create theme configuration"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "customization saved successfully"})
		return
	}

	// Update existing shop theme
	updateData := bson.M{
		"overrides": models.ThemeConfig{
			Colors:  customizationData.Colors,
			Fonts:   customizationData.Fonts,
			Layouts: customizationData.Layout,
		},
	}

	_, err = services.UpdateShopThemeService(shopTheme.ID.Hex(), updateData, sellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customization updated successfully"})
}

// =============== Helper Functions ===============

// mergeThemeConfig merges base theme config with overrides
func mergeThemeConfig(base, overrides models.ThemeConfig) map[string]interface{} {
	result := make(map[string]interface{})

	// Merge colors
	colors := make(map[string]string)
	for k, v := range base.Colors {
		colors[k] = v
	}
	for k, v := range overrides.Colors {
		colors[k] = v
	}
	result["colors"] = colors

	// Merge fonts
	fonts := make(map[string]string)
	for k, v := range base.Fonts {
		fonts[k] = v
	}
	for k, v := range overrides.Fonts {
		fonts[k] = v
	}
	result["fonts"] = fonts

	// Merge layouts
	layouts := make(map[string]string)
	for k, v := range base.Layouts {
		layouts[k] = v
	}
	for k, v := range overrides.Layouts {
		layouts[k] = v
	}
	result["layout"] = layouts

	return result
}
