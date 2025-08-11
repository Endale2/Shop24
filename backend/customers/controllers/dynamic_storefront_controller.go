package controllers

import (
	"net/http"
	"fmt"

	shopServices "github.com/Endale2/DRPS/shared/services"
	themeServices "github.com/Endale2/DRPS/shared/services"
	sharedRepositories "github.com/Endale2/DRPS/shared/repositories"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/gin-gonic/gin"
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
	theme, err := themeServices.GetOrCreateShopTheme(shop.ID, shop.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load theme configuration"})
		return
	}

	// Get shop analytics for enhanced storefront
	customers, _ := sharedRepositories.GetShopCustomerLinks(shop.ID)

	// Generate dynamic CSS from theme
	dynamicCSS := generateStorefrontCSS(theme)

	// Get theme performance metrics
	themeMetrics, _ := themeServices.GetThemePerformanceMetrics(shop.ID)

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
			"id":               theme.ID,
			"name":             theme.Name,
			"version":          theme.Version,
			"colors":           theme.Colors,
			"fonts":            theme.Fonts,
			"layout":           theme.Layout,
			"customCSS":        theme.CustomCSS,
			"seo":              theme.SEO,
			"mobile":           theme.Mobile,
			"dynamicCSS":       dynamicCSS,
			"performance":      themeMetrics,
			"lastModified":     theme.UpdatedAt,
			"previewImage":     theme.PreviewImage,
			"compiledCSS":      theme.CompiledCSS,
		},
		"layout": generateLayoutConfig(theme),
		"components": generateComponentConfig(theme),
		"navigation": generateNavigationConfig(theme),
		"footer": generateFooterConfig(theme),
		"seo": generateSEOConfig(shop, theme),
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
	theme, err := themeServices.GetOrCreateShopTheme(shop.ID, shop.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load theme"})
		return
	}

	// Generate fresh CSS
	dynamicCSS := generateStorefrontCSS(theme)

	themeResponse := gin.H{
		"colors":      theme.Colors,
		"fonts":       theme.Fonts,
		"layout":      theme.Layout,
		"customCSS":   theme.CustomCSS,
		"seo":         theme.SEO,
		"mobile":      theme.Mobile,
		"dynamicCSS":  dynamicCSS,
		"version":     theme.Version,
		"lastModified": theme.UpdatedAt,
		"cacheKey":    theme.CacheKey,
	}

	c.JSON(http.StatusOK, themeResponse)
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
		getLayoutWidth(getLayoutValue(theme.Layout, "containerWidth", "boxed")),
		getHeaderHeight(getLayoutValue(theme.Layout, "headerStyle", "classic")),
		getLayoutValue(theme.Layout, "gridColumns", "3"),
		getBorderRadius(getLayoutValue(theme.Layout, "borderStyle", "rounded")),
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

	// Add custom CSS if available
	if theme.CustomCSS != "" {
		css += "\n\n/* Custom CSS */\n" + theme.CustomCSS
	}

	return css
}

// generateLayoutConfig creates layout configuration based on theme
func generateLayoutConfig(theme *models.ShopTheme) gin.H {
	return gin.H{
		"containerWidth":  getLayoutValue(theme.Layout, "containerWidth", "boxed"),
		"headerStyle":     getLayoutValue(theme.Layout, "headerStyle", "classic"),
		"footerStyle":     getLayoutValue(theme.Layout, "footerStyle", "minimal"),
		"sidebarEnabled":  getLayoutValue(theme.Layout, "sidebarPosition", "none") != "none",
		"gridColumns":     getLayoutValue(theme.Layout, "gridColumns", "3"),
		"spacing":         getLayoutValue(theme.Layout, "spacing", "normal"),
		"borderStyle":     getLayoutValue(theme.Layout, "borderStyle", "rounded"),
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
			"style":       getLayoutValue(theme.Layout, "buttonStyle", "rounded"),
			"size":        "medium",
			"animation":   "subtle",
		},
		"navigation": gin.H{
			"style":       getLayoutValue(theme.Layout, "navStyle", "horizontal"),
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
		"style": getLayoutValue(theme.Layout, "navStyle", "horizontal"),
		"showIcons": getLayoutValue(theme.Layout, "showNavIcons", "true") == "true",
		"sticky": getLayoutValue(theme.Layout, "stickyHeader", "true") == "true",
	}
}

// generateFooterConfig creates footer configuration
func generateFooterConfig(theme *models.ShopTheme) gin.H {
	return gin.H{
		"style": getLayoutValue(theme.Layout, "footerStyle", "minimal"),
		"showSocial": true,
		"showNewsletter": true,
		"columns": 3,
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

// Layout helpers
func getLayoutValue(layout map[string]string, key string, defaultValue string) string {
	if layout == nil {
		return defaultValue
	}
	if value, exists := layout[key]; exists && value != "" {
		return value
	}
	return defaultValue
}

// SEO helpers
func getSEOValue(seo map[string]string, key string, defaultValue string) string {
	if seo == nil {
		return defaultValue
	}
	if value, exists := seo[key]; exists && value != "" {
		return value
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
