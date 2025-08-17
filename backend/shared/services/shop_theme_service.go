package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	shopRepo "github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// =============== Theme Management Services ===============

// GetOrCreateShopTheme gets the active theme for a shop, or creates a default one if none exists
func GetOrCreateShopTheme(shopID, sellerID primitive.ObjectID) (*models.ShopTheme, error) {
	// Try to get from cache first
	cache := GetCacheService()
	cacheKey := ThemeCacheKey(shopID.Hex())
	if cachedData, found := cache.Get(cacheKey); found {
		if theme, ok := cachedData.(*models.ShopTheme); ok {
			return theme, nil
		}
	}

	// Try to get existing active theme
	theme, err := repositories.GetActiveShopTheme(shopID)
	if err != nil {
		return nil, err
	}

	// If theme exists, cache it and return
	if theme != nil {
		cache := GetCacheService()
		cache.Set(ThemeCacheKey(shopID.Hex()), theme, CacheMedium)
		return theme, nil
	}

	// If no theme exists, create a default one
	defaultTheme := models.GetDefaultShopTheme(shopID, sellerID)
	_, err = repositories.CreateShopTheme(defaultTheme)
	if err != nil {
		return nil, err
	}

	// Update shop reference to this theme
	err = updateShopThemeReference(shopID, defaultTheme.ID, defaultTheme.Colors["primary"], defaultTheme.Version)
	if err != nil {
		return nil, err
	}

	// Cache the new theme
	cache.Set(ThemeCacheKey(shopID.Hex()), defaultTheme, CacheMedium)
	return defaultTheme, nil
}

// UpdateShopThemeColors updates only the colors for a shop's active theme
func UpdateShopThemeColors(shopID primitive.ObjectID, colors map[string]string) (*models.ShopTheme, error) {
	theme, err := repositories.UpdateThemeColors(shopID, colors)
	if err != nil {
		return nil, err
	}

	// Update shop's cached primary color
	if primaryColor, exists := colors["primary"]; exists {
		err = updateShopThemeReference(shopID, theme.ID, primaryColor, theme.Version)
		if err != nil {
			return nil, err
		}
	}

	return theme, nil
}

// UpdateShopThemeFonts updates only the fonts for a shop's active theme
func UpdateShopThemeFonts(shopID primitive.ObjectID, fonts map[string]string) (*models.ShopTheme, error) {
	theme, err := repositories.UpdateThemeFonts(shopID, fonts)
	if err != nil {
		return nil, err
	}

	// Update shop's theme reference
	err = updateShopThemeReference(shopID, theme.ID, "", theme.Version)
	if err != nil {
		return nil, err
	}

	return theme, nil
}

// UpdateShopThemeLayout updates only the layout for a shop's active theme
func UpdateShopThemeLayout(shopID primitive.ObjectID, layout map[string]string) (*models.ShopTheme, error) {
	theme, err := repositories.UpdateThemeLayout(shopID, layout)
	if err != nil {
		return nil, err
	}

	// Update shop's theme reference
	err = updateShopThemeReference(shopID, theme.ID, "", theme.Version)
	if err != nil {
		return nil, err
	}

	return theme, nil
}

// UpdateShopThemeComplete updates multiple sections of a shop's theme
func UpdateShopThemeComplete(shopID primitive.ObjectID, updateData map[string]interface{}) (*models.ShopTheme, error) {
	// Get current active theme
	theme, err := repositories.GetActiveShopTheme(shopID)
	if err != nil {
		return nil, err
	}
	if theme == nil {
		return nil, errors.New("no active theme found")
	}

	// Create backup before update
	updateBSON := bson.M{}
	for key, value := range updateData {
		updateBSON[key] = value
	}

	err = repositories.CreateBackupAndUpdate(theme.ID, updateBSON)
	if err != nil {
		return nil, err
	}

	// Get updated theme
	updatedTheme, err := repositories.GetShopThemeByID(theme.ID)
	if err != nil {
		return nil, err
	}

	// Update shop reference if colors changed
	if colors, exists := updateData["colors"]; exists {
		if colorMap, ok := colors.(map[string]string); ok {
			if primaryColor, exists := colorMap["primary"]; exists {
				err = updateShopThemeReference(shopID, theme.ID, primaryColor, updatedTheme.Version)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return updatedTheme, nil
}

// ResetShopTheme resets a shop's theme to default configuration
func ResetShopTheme(shopID, sellerID primitive.ObjectID) (*models.ShopTheme, error) {
	theme, err := repositories.ResetThemeToDefault(shopID, sellerID)
	if err != nil {
		return nil, err
	}

	// Update shop reference
	err = updateShopThemeReference(shopID, theme.ID, theme.Colors["primary"], theme.Version)
	if err != nil {
		return nil, err
	}

	return theme, nil
}

// GetShopThemeConfig returns the complete theme configuration for API responses
func GetShopThemeConfig(shopID primitive.ObjectID) (*models.ShopThemeConfig, error) {
	theme, err := repositories.GetActiveShopTheme(shopID)
	if err != nil {
		return nil, err
	}
	if theme == nil {
		return nil, errors.New("no active theme found")
	}

	config := theme.ToConfig()
	return &config, nil
}

// CloneShopTheme creates a copy of an existing theme
func CloneShopTheme(sourceThemeID primitive.ObjectID, newShopID, sellerID primitive.ObjectID, newName string) (*models.ShopTheme, error) {
	// Get source theme
	sourceTheme, err := repositories.GetShopThemeByID(sourceThemeID)
	if err != nil {
		return nil, err
	}
	if sourceTheme == nil {
		return nil, errors.New("source theme not found")
	}

	// Create new theme based on source
	newTheme := &models.ShopTheme{
		ID:          primitive.NewObjectID(),
		ShopID:      newShopID,
		CreatedBy:   sellerID,
		Name:        newName,
		Description: "Cloned from " + sourceTheme.Name,
		Version:     "1.0.0",
		IsActive:    true,
		IsDefault:   false,
		Colors:      sourceTheme.Colors,
		Fonts:       sourceTheme.Fonts,
		Layout:      sourceTheme.Layout,
		CustomCSS:   sourceTheme.CustomCSS,
		SEO:         sourceTheme.SEO,
		Mobile:      sourceTheme.Mobile,
		Tags:        append(sourceTheme.Tags, "cloned"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Deactivate existing themes for the shop
	err = repositories.SetActiveTheme(newShopID, newTheme.ID)
	if err != nil {
		return nil, err
	}

	// Create the new theme
	_, err = repositories.CreateShopTheme(newTheme)
	if err != nil {
		return nil, err
	}

	// Update shop reference
	err = updateShopThemeReference(newShopID, newTheme.ID, newTheme.Colors["primary"], newTheme.Version)
	if err != nil {
		return nil, err
	}

	return newTheme, nil
}

// =============== Theme Analytics Services ===============

// RecordThemeUsage records theme usage for analytics
func RecordThemeUsage(themeID primitive.ObjectID, loadTime float64) error {
	return repositories.UpdateThemeAnalytics(themeID, loadTime, 0, 1)
}

// RecordThemeConversion records a conversion for theme analytics
func RecordThemeConversion(themeID primitive.ObjectID) error {
	return repositories.UpdateThemeAnalytics(themeID, 0, 1, 0)
}

// GetThemePerformanceMetrics returns performance metrics for a theme
func GetThemePerformanceMetrics(shopID primitive.ObjectID) (map[string]interface{}, error) {
	theme, err := repositories.GetActiveShopTheme(shopID)
	if err != nil {
		return nil, err
	}
	if theme == nil {
		return nil, errors.New("no active theme found")
	}

	// Calculate performance score (simple algorithm)
	performanceScore := 100
	if theme.LoadTime > 3.0 {
		performanceScore -= 20
	}
	if theme.LoadTime > 5.0 {
		performanceScore -= 30
	}

	metrics := map[string]interface{}{
		"loadTime":         theme.LoadTime,
		"conversions":      theme.Conversions,
		"views":            theme.Views,
		"usageCount":       theme.UsageCount,
		"performanceScore": performanceScore,
		"conversionRate":   calculateConversionRate(theme.Conversions, theme.Views),
		"lastUsed":         theme.LastUsedAt,
		"version":          theme.Version,
	}

	return metrics, nil
}

// =============== Theme Validation Services ===============

// ValidateThemeConfiguration validates a theme configuration
func ValidateThemeConfiguration(config map[string]interface{}) error {
	// Validate colors
	if colors, exists := config["colors"]; exists {
		if colorMap, ok := colors.(map[string]string); ok {
			for key, value := range colorMap {
				if key == "" || value == "" {
					return errors.New("color configuration cannot have empty keys or values")
				}
				if !isValidHexColor(value) {
					return errors.New("invalid color format: " + value)
				}
			}
		}
	}

	// Validate fonts
	if fonts, exists := config["fonts"]; exists {
		if fontMap, ok := fonts.(map[string]string); ok {
			for key, value := range fontMap {
				if key == "" || value == "" {
					return errors.New("font configuration cannot have empty keys or values")
				}
			}
		}
	}

	// Validate custom CSS size
	if customCSS, exists := config["customCSS"]; exists {
		if cssStr, ok := customCSS.(string); ok {
			if len(cssStr) > 100000 { // 100KB limit
				return errors.New("custom CSS too large (max 100KB)")
			}
		}
	}

	return nil
}

// =============== Helper Functions ===============

// updateShopThemeReference updates the shop's theme reference fields for quick access
func updateShopThemeReference(shopID, themeID primitive.ObjectID, primaryColor, version string) error {
	updateData := bson.M{
		"activeThemeId":     themeID,
		"themeLastModified": time.Now(),
	}

	if primaryColor != "" {
		updateData["themeColor"] = primaryColor
	}
	if version != "" {
		updateData["themeVersion"] = version
	}

	_, err := shopRepo.UpdateShop(shopID.Hex(), updateData)
	return err
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

// calculateConversionRate calculates conversion rate from conversions and views
func calculateConversionRate(conversions, views int) float64 {
	if views == 0 {
		return 0.0
	}
	return float64(conversions) / float64(views) * 100.0
}

// =============== Theme Cache Services ===============

// GenerateThemeCSS generates and caches compiled CSS for a theme
func GenerateThemeCSS(themeID primitive.ObjectID) (string, error) {
	theme, err := repositories.GetShopThemeByID(themeID)
	if err != nil {
		return "", err
	}
	if theme == nil {
		return "", errors.New("theme not found")
	}

	// Generate CSS from theme configuration
	// This is a simplified version - in production you'd use a proper CSS compiler
	css := generateCSSFromTheme(theme)

	// Generate cache key
	cacheKey := generateCacheKey(theme)

	// Update theme cache
	err = repositories.UpdateThemeCache(themeID, css, cacheKey)
	if err != nil {
		return "", err
	}

	return css, nil
}

// generateCSSFromTheme generates CSS from theme configuration
func generateCSSFromTheme(theme *models.ShopTheme) string {
	css := ":root {\n"

	// Add color variables
	for key, value := range theme.Colors {
		css += "  --color-" + key + ": " + value + ";\n"
	}

	// Add font variables
	for key, value := range theme.Fonts {
		css += "  --font-" + key + ": '" + value + "';\n"
	}

	css += "}\n\n"

	// Add custom CSS
	customCSS := theme.GetCustomCSSCompiled()
	if customCSS == "" {
		customCSS = theme.GetCustomCSSRaw()
	}
	if customCSS != "" {
		css += customCSS + "\n"
	}

	return css
}

// generateCacheKey generates a cache key for theme
func generateCacheKey(theme *models.ShopTheme) string {
	return theme.ID.Hex() + "_" + theme.Version + "_" + theme.UpdatedAt.Format("20060102150405")
}

// =============== Advanced Styling Service Functions ===============

// UpdateThemeGradients updates gradient configurations for a shop theme
func UpdateThemeGradients(shopID primitive.ObjectID, gradients map[string]string) error {
	err := repositories.UpdateThemeGradients(shopID, gradients)
	if err == nil {
		cache := GetCacheService()
		cache.Delete(ThemeCacheKey(shopID.Hex()))
	}
	return err
}

// UpdateThemeShadows updates shadow configurations for a shop theme
func UpdateThemeShadows(shopID primitive.ObjectID, shadows map[string]string) error {
	err := repositories.UpdateThemeShadows(shopID, shadows)
	if err == nil {
		cache := GetCacheService()
		cache.Delete(ThemeCacheKey(shopID.Hex()))
	}
	return err
}

// UpdateThemeAnimations updates animation configurations for a shop theme
func UpdateThemeAnimations(shopID primitive.ObjectID, animations map[string]string) error {
	err := repositories.UpdateThemeAnimations(shopID, animations)
	if err == nil {
		cache := GetCacheService()
		cache.Delete(ThemeCacheKey(shopID.Hex()))
	}
	return err
}

// UpdateThemeMobileConfig updates mobile configuration for a shop theme
func UpdateThemeMobileConfig(shopID primitive.ObjectID, mobileConfig map[string]interface{}) error {
	err := repositories.UpdateThemeMobileConfig(shopID, mobileConfig)
	if err == nil {
		cache := GetCacheService()
		cache.Delete(ThemeCacheKey(shopID.Hex()))
	}
	return err
}

// UpdateThemeSEOConfig updates SEO configuration for a shop theme
func UpdateThemeSEOConfig(shopID primitive.ObjectID, seoConfig map[string]interface{}) error {
	err := repositories.UpdateThemeSEOConfig(shopID, seoConfig)
	if err == nil {
		cache := GetCacheService()
		cache.Delete(ThemeCacheKey(shopID.Hex()))
	}
	return err
}
