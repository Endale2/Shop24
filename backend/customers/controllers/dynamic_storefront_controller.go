package controllers

import (
	"fmt"
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	sharedRepositories "github.com/Endale2/DRPS/shared/repositories"
	shopServices "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetDynamicStorefront serves the complete theme-driven storefront configuration
// GET /shops/:shopSlug/storefront
func GetDynamicStorefront(c *gin.Context) {
	shopSlug := c.Param("shopSlug")

	// Get shop details
	shop, err := shopServices.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Get shop theme configuration from scalable theme collection
	theme, err := shopServices.GetOrCreateShopTheme(shop.ID, shop.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load theme configuration"})
		return
	}

	// Get shop analytics for enhanced storefront
	customers, _ := sharedRepositories.GetShopCustomerLinks(shop.ID)

	// Generate dynamic CSS from theme
	dynamicCSS := generateStorefrontCSS(theme)

	// Get theme performance metrics (placeholder for future implementation)
	var themeMetrics map[string]interface{}

	// Prepare complete storefront configuration
	storefrontConfig := gin.H{
		"shop": gin.H{
			"id":          shop.ID,
			"name":        shop.Name,
			"slug":        shop.Slug,
			"description": shop.Description,
			"image":       shop.Image,
			"banner":      shop.Banner,
			"email":       shop.Email,
			"phone":       shop.Phone,
			"address":     shop.Address,
			"currency":    shop.Currency,
			"status":      shop.Status,
			"isVerified":  shop.IsVerified,
			"rating":      shop.Rating,
			"reviewCount": shop.ReviewCount,
		},
		"theme": gin.H{
			"id":           theme.ID,
			"name":         theme.Name,
			"version":      theme.Version,
			"colors":       theme.Colors,
			"fonts":        theme.Fonts,
			"layout":       theme.Layout,
			"customCSS":    theme.CustomCSS,
			"seo":          theme.SEO,
			"mobile":       theme.Mobile,
			"dynamicCSS":   dynamicCSS,
			"performance":  themeMetrics,
			"lastModified": theme.UpdatedAt,
			"previewImage": theme.PreviewImage,
			"compiledCSS":  theme.CompiledCSS,
		},
		"layout":     generateLayoutConfig(theme),
		"components": generateComponentConfig(theme),
		"navigation": generateNavigationConfig(theme),
		"footer":     generateFooterConfig(theme),
		"seo":        generateSEOConfig(shop, theme),
		"analytics": gin.H{
			"customerCount": len(customers),
			"productCount":  shop.ProductCount,
			"revenue":       shop.Revenue,
			"shopAge":       shop.CreatedAt,
		},
		"scalable": true,
		"dynamic":  true,
	}

	c.JSON(http.StatusOK, storefrontConfig)
}

// GetStorefrontTheme serves just the theme configuration for real-time updates
// GET /shops/:shopSlug/theme
func GetStorefrontTheme(c *gin.Context) {
	shopSlug := c.Param("shopSlug")

	// Get shop details
	shop, err := shopServices.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Get theme configuration
	theme, err := shopServices.GetOrCreateShopTheme(shop.ID, shop.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load theme"})
		return
	}

	// Generate fresh CSS
	dynamicCSS := generateStorefrontCSS(theme)

	themeResponse := gin.H{
		"colors":       theme.Colors,
		"fonts":        theme.Fonts,
		"layout":       theme.Layout,
		"customCSS":    theme.GetCustomCSSCompiled(),
		"seo":          theme.SEO,
		"mobile":       theme.Mobile,
		"gradients":    theme.Gradients,
		"shadows":      theme.Shadows,
		"animations":   theme.Animations,
		"spacing":      theme.Spacing,
		"dynamicCSS":   dynamicCSS,
		"version":      theme.Version,
		"lastModified": theme.UpdatedAt,
		"cacheKey":     theme.CacheKey,
	}

	c.JSON(http.StatusOK, themeResponse)
}

// GetShopComponents returns component data for a specific shop
// GET /shops/:shopSlug/components/:componentType
func GetShopComponents(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	componentType := c.Param("componentType")

	// Get shop by slug
	shop, err := sharedRepositories.GetShopBySlug(shopSlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Get shop theme
	theme, err := shopServices.GetOrCreateShopTheme(shop.ID, shop.OwnerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "theme not found"})
		return
	}

	// Get component instances for this shop and component type
	components, err := getComponentsByType(shop.ID, componentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load components"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"components": components,
		"theme": gin.H{
			"colors": theme.Colors,
			"fonts":  theme.Fonts,
		},
	})
}

// GetShopPageLayout returns layout configuration for a specific page
// GET /shops/:shopSlug/pages/:pageType
func GetShopPageLayout(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	pageType := c.Param("pageType")

	// Get shop by slug
	shop, err := sharedRepositories.GetShopBySlug(shopSlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Get page layout
	layout, err := getPageLayout(shop.ID, pageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load page layout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"layout":  layout,
	})
}

// =============== Theme Configuration Generators ===============

// generateStorefrontCSS creates dynamic CSS from theme configuration
func generateStorefrontCSS(theme *models.ShopTheme) string {
	if theme == nil {
		return getDefaultCSS()
	}

	// Generate CSS variables from theme data
	css := fmt.Sprintf(`:root {
		/* Dynamic Color Variables */
		--color-primary: %s;
		--color-secondary: %s;
		--color-background: %s;
		--color-heading: %s;
		--color-body-text: %s;
		--color-border: %s;
		
		/* Dynamic Font Variables */
		--font-heading: '%s', sans-serif;
		--font-body: '%s', sans-serif;
		
		/* Dynamic Layout Variables */
		--container-width: %s;
		--header-height: %s;
		--grid-columns: %s;
		--border-radius: %s;
	}`,
		getColorValue(theme.Colors, "primary", "#10B981"),
		getColorValue(theme.Colors, "secondary", "#F59E0B"),
		getColorValue(theme.Colors, "background", "#FFFFFF"),
		getColorValue(theme.Colors, "heading", "#1F2937"),
		getColorValue(theme.Colors, "bodyText", "#6B7280"),
		getColorValue(theme.Colors, "border", "#E5E7EB"),
		getFontValue(theme.Fonts, "heading", "Inter"),
		getFontValue(theme.Fonts, "body", "Inter"),
		getLayoutWidth(theme.GetLayoutValue("containerWidth", "boxed")),
		getHeaderHeight(theme.GetLayoutValue("headerStyle", "classic")),
		theme.GetLayoutValue("gridColumns", "3"),
		getBorderRadius(theme.GetLayoutValue("borderStyle", "rounded")),
	)

	// Add the CSS styles
	css += `
	/* Dynamic Base Styles */
	body {
		font-family: var(--font-body);
		background-color: var(--color-background);
		color: var(--color-body-text);
	}

	h1, h2, h3, h4, h5, h6 {
		font-family: var(--font-heading);
		color: var(--color-heading);
	}

	/* Dynamic Button Styles */
	.btn-primary {
		background-color: var(--color-primary);
		border: 1px solid var(--color-primary);
		border-radius: var(--border-radius);
		transition: all 0.2s ease;
	}

	.btn-primary:hover {
		opacity: 0.9;
		transform: translateY(-1px);
	}

	.btn-secondary {
		background-color: var(--color-secondary);
		border: 1px solid var(--color-secondary);
		border-radius: var(--border-radius);
	}

	/* Dynamic Header Styles */
	.dynamic-header {
		background-color: var(--color-background);
		border-bottom: 1px solid var(--color-border);
		height: var(--header-height);
	}

	/* Dynamic Container */
	.dynamic-container {
		max-width: var(--container-width);
		margin: 0 auto;
		padding: 0 1rem;
	}

	/* Dynamic Grid */
	.dynamic-grid {
		display: grid;
		grid-template-columns: repeat(var(--grid-columns), 1fr);
		gap: 2rem;
	}

	/* Dynamic Product Card */
	.product-card {
		background-color: var(--color-background);
		border: 1px solid var(--color-border);
		border-radius: var(--border-radius);
		transition: all 0.2s ease;
	}

	.product-card:hover {
		border-color: var(--color-primary);
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
	}

	/* Responsive Grid */
	@media (max-width: 768px) {
		.dynamic-grid {
			grid-template-columns: repeat(2, 1fr);
		}
	}

	@media (max-width: 640px) {
		.dynamic-grid {
			grid-template-columns: 1fr;
		}
	}
	`

	// Add gradients
	if theme.Gradients != nil && len(theme.Gradients) > 0 {
		css += "\n\n/* Gradients */\n"
		for name, gradient := range theme.Gradients {
			css += fmt.Sprintf(".gradient-%s { background: %s; }\n", name, gradient)
		}
	}

	// Add shadows
	if theme.Shadows != nil && len(theme.Shadows) > 0 {
		css += "\n\n/* Shadows */\n"
		for name, shadow := range theme.Shadows {
			css += fmt.Sprintf(".shadow-%s { box-shadow: %s; }\n", name, shadow)
		}
	}

	// Add animations
	if theme.Animations != nil && len(theme.Animations) > 0 {
		css += "\n\n/* Animations */\n"
		for name, animation := range theme.Animations {
			css += fmt.Sprintf(".animate-%s { animation: %s; }\n", name, animation)
		}
	}

	// Add spacing utilities
	if theme.Spacing != nil && len(theme.Spacing) > 0 {
		css += "\n\n/* Spacing */\n"
		for name, spacing := range theme.Spacing {
			css += fmt.Sprintf(".spacing-%s { margin: %s; }\n", name, spacing)
			css += fmt.Sprintf(".padding-%s { padding: %s; }\n", name, spacing)
		}
	}

	// Add custom CSS if available
	customCSS := theme.GetCustomCSSCompiled()
	if customCSS == "" {
		customCSS = theme.GetCustomCSSRaw()
	}
	if customCSS != "" {
		css += "\n\n/* Custom CSS */\n" + customCSS
	}

	return css
}

// generateLayoutConfig creates layout configuration based on theme
func generateLayoutConfig(theme *models.ShopTheme) gin.H {
	return gin.H{
		"containerWidth": theme.GetLayoutValue("containerWidth", "boxed"),
		"headerStyle":    theme.GetLayoutValue("headerStyle", "classic"),
		"footerStyle":    theme.GetLayoutValue("footerStyle", "minimal"),
		"sidebarEnabled": theme.GetLayoutValue("sidebarPosition", "none") != "none",
		"gridColumns":    theme.GetLayoutValue("gridColumns", "3"),
		"spacing":        theme.GetLayoutValue("spacing", "normal"),
		"borderStyle":    theme.GetLayoutValue("borderStyle", "rounded"),
	}
}

// generateComponentConfig creates component configuration
func generateComponentConfig(theme *models.ShopTheme) gin.H {
	return gin.H{
		"productCard": gin.H{
			"style":       "dynamic",
			"showRating":  true,
			"showBadges":  true,
			"imageRatio":  "square",
			"hoverEffect": "lift",
		},
		"buttons": gin.H{
			"style":     theme.GetLayoutValue("buttonStyle", "rounded"),
			"size":      "medium",
			"animation": "subtle",
		},
		"navigation": gin.H{
			"style":       theme.GetLayoutValue("navStyle", "horizontal"),
			"position":    "top",
			"transparent": false,
		},
	}
}

// generateNavigationConfig creates navigation based on theme
func generateNavigationConfig(theme *models.ShopTheme) gin.H {
	return gin.H{
		"items": []gin.H{
			{"label": "Home", "path": "/", "icon": "home"},
			{"label": "Products", "path": "/products", "icon": "shopping-bag"},
			{"label": "Collections", "path": "/collections", "icon": "collection"},
		},
		"style":     getLayoutValue(theme.Layout, "navStyle", "horizontal"),
		"showIcons": getLayoutValue(theme.Layout, "showNavIcons", "true") == "true",
		"sticky":    getLayoutValue(theme.Layout, "stickyHeader", "true") == "true",
	}
}

// generateFooterConfig creates footer configuration
func generateFooterConfig(theme *models.ShopTheme) gin.H {
	return gin.H{
		"style":          getLayoutValue(theme.Layout, "footerStyle", "minimal"),
		"showSocial":     true,
		"showNewsletter": true,
		"columns":        3,
	}
}

// generateSEOConfig creates SEO configuration
func generateSEOConfig(shop *models.Shop, theme *models.ShopTheme) gin.H {
	return gin.H{
		"title":       getSEOValue(theme.SEO, "metaTitle", shop.Name),
		"description": getSEOValue(theme.SEO, "metaDescription", shop.Description),
		"keywords":    getSEOValue(theme.SEO, "keywords", ""),
		"ogImage":     shop.Image,
		"themeColor":  getColorValue(theme.Colors, "primary", "#10B981"),
	}
}

// =============== Helper Functions ===============

// Color helpers
func getColorValue(colors map[string]string, key string, defaultValue string) string {
	if colors == nil {
		return defaultValue
	}
	if value, exists := colors[key]; exists && value != "" {
		return value
	}
	return defaultValue
}

// Font helpers
func getFontValue(fonts map[string]string, key string, defaultValue string) string {
	if fonts == nil {
		return defaultValue
	}
	if value, exists := fonts[key]; exists && value != "" {
		return value
	}
	return defaultValue
}

// Layout helpers - updated to work with interface{} types
func getLayoutValue(layout map[string]interface{}, key string, defaultValue string) string {
	if layout == nil {
		return defaultValue
	}
	if value, exists := layout[key]; exists {
		if strValue, ok := value.(string); ok && strValue != "" {
			return strValue
		}
	}
	return defaultValue
}

// SEO helpers - updated to work with interface{} types
func getSEOValue(seo map[string]interface{}, key string, defaultValue string) string {
	if seo == nil {
		return defaultValue
	}

	// Handle nested structure for new SEO format
	if meta, ok := seo["meta"].(map[string]interface{}); ok {
		if value, exists := meta[key]; exists {
			if strValue, ok := value.(string); ok && strValue != "" {
				return strValue
			}
		}
	}

	// Fallback to direct key access for backward compatibility
	if value, exists := seo[key]; exists {
		if strValue, ok := value.(string); ok && strValue != "" {
			return strValue
		}
	}

	return defaultValue
}

func getLayoutWidth(containerType string) string {
	switch containerType {
	case "full":
		return "100%"
	case "wide":
		return "1400px"
	case "boxed":
		return "1200px"
	default:
		return "1200px"
	}
}

func getHeaderHeight(headerStyle string) string {
	switch headerStyle {
	case "compact":
		return "60px"
	case "large":
		return "100px"
	default:
		return "80px"
	}
}

func getBorderRadius(borderStyle string) string {
	switch borderStyle {
	case "sharp":
		return "0px"
	case "rounded":
		return "8px"
	case "pill":
		return "50px"
	default:
		return "8px"
	}
}

func getDefaultCSS() string {
	return `:root {
		--color-primary: #10B981;
		--color-secondary: #F59E0B;
		--color-background: #FFFFFF;
		--color-heading: #1F2937;
		--color-body-text: #6B7280;
		--font-heading: 'Inter', sans-serif;
		--font-body: 'Inter', sans-serif;
		--container-width: 1200px;
	}`
}

// Helper functions for component and layout management

// getComponentsByType retrieves component instances by type for a shop
func getComponentsByType(shopID primitive.ObjectID, componentType string) ([]interface{}, error) {
	// TODO: Implement component retrieval from database
	// This would query the component_instances collection
	return []interface{}{}, nil
}

// getPageLayout retrieves page layout configuration for a shop
func getPageLayout(shopID primitive.ObjectID, pageType string) (interface{}, error) {
	// TODO: Implement page layout retrieval from database
	// This would query the page_layouts collection
	return map[string]interface{}{
		"pageType":   pageType,
		"components": []interface{}{},
		"settings":   map[string]interface{}{},
	}, nil
}
