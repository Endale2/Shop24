package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Endale2/DRPS/shared/repositories"
	shopServices "github.com/Endale2/DRPS/shared/services"
)

// ValidateCSS validates custom CSS for security and syntax
// POST /seller/shops/:shopId/customization/css/validate
func ValidateCSS(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context
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

	var req struct {
		CSS string `json:"css" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data", "details": err.Error()})
		return
	}

	// Validate CSS
	validation := validateCSSContent(req.CSS)

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"validation": validation,
	})
}

// CompileCSS compiles and optimizes custom CSS
// POST /seller/shops/:shopId/customization/css/compile
func CompileCSS(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context
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

	var req struct {
		CSS string `json:"css" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data", "details": err.Error()})
		return
	}

	// Validate CSS first
	validation := validateCSSContent(req.CSS)
	if !validation["isValid"].(bool) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "CSS validation failed",
			"validation": validation,
		})
		return
	}

	// Compile CSS (minify and optimize)
	compiledCSS := compileCSSContent(req.CSS)

	// Update theme with compiled CSS
	shopObjectID, _ := primitive.ObjectIDFromHex(shopID)
	customCSS := map[string]interface{}{
		"raw":        req.CSS,
		"compiled":   compiledCSS,
		"validation": validation,
	}

	err = repositories.UpdateThemeCustomCSS(shopObjectID, customCSS)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save compiled CSS", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"compiled":   compiledCSS,
		"validation": validation,
	})
}

// GetCSSVariables returns available CSS variables for the theme
// GET /seller/shops/:shopId/customization/css/variables
func GetCSSVariables(c *gin.Context) {
	shopID := c.Param("shopId")

	// Get seller ID from context
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

	// Get theme to extract variables
	shopObjectID, _ := primitive.ObjectIDFromHex(shopID)
	theme, err := shopServices.GetOrCreateShopTheme(shopObjectID, primitive.NilObjectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "theme not found"})
		return
	}

	// Generate CSS variables from theme
	variables := generateCSSVariables(theme)

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"variables": variables,
	})
}

// Helper functions

// validateCSSContent validates CSS for security and syntax
func validateCSSContent(css string) map[string]interface{} {
	validation := map[string]interface{}{
		"isValid":  true,
		"errors":   []string{},
		"warnings": []string{},
	}

	errors := []string{}
	warnings := []string{}

	// Check for dangerous patterns
	dangerousPatterns := []string{
		`@import`,
		`javascript:`,
		`expression\(`,
		`behavior:`,
		`-moz-binding`,
		`@-webkit-keyframes`,
		`@keyframes.*{.*}.*{.*}`, // Complex keyframes
	}

	for _, pattern := range dangerousPatterns {
		if matched, _ := regexp.MatchString(pattern, css); matched {
			errors = append(errors, "Dangerous CSS pattern detected: "+pattern)
		}
	}

	// Check for syntax issues
	if strings.Count(css, "{") != strings.Count(css, "}") {
		errors = append(errors, "Mismatched braces in CSS")
	}

	// Check for performance warnings
	if strings.Contains(css, "*") {
		warnings = append(warnings, "Universal selector (*) may impact performance")
	}

	if len(strings.Split(css, "\n")) > 1000 {
		warnings = append(warnings, "CSS is very large, consider optimization")
	}

	validation["errors"] = errors
	validation["warnings"] = warnings
	validation["isValid"] = len(errors) == 0

	return validation
}

// compileCSSContent compiles and minifies CSS
func compileCSSContent(css string) string {
	// Basic CSS minification
	compiled := css

	// Remove comments
	commentRegex := regexp.MustCompile(`/\*.*?\*/`)
	compiled = commentRegex.ReplaceAllString(compiled, "")

	// Remove extra whitespace
	whitespaceRegex := regexp.MustCompile(`\s+`)
	compiled = whitespaceRegex.ReplaceAllString(compiled, " ")

	// Remove spaces around certain characters
	compiled = strings.ReplaceAll(compiled, " {", "{")
	compiled = strings.ReplaceAll(compiled, "{ ", "{")
	compiled = strings.ReplaceAll(compiled, " }", "}")
	compiled = strings.ReplaceAll(compiled, "} ", "}")
	compiled = strings.ReplaceAll(compiled, "; ", ";")
	compiled = strings.ReplaceAll(compiled, " ;", ";")
	compiled = strings.ReplaceAll(compiled, ": ", ":")
	compiled = strings.ReplaceAll(compiled, " :", ":")

	// Trim
	compiled = strings.TrimSpace(compiled)

	return compiled
}

// generateCSSVariables generates CSS variables from theme
func generateCSSVariables(theme interface{}) map[string]string {
	variables := map[string]string{
		"--primary-color":    "#2563EB",
		"--secondary-color":  "#64748B",
		"--background-color": "#FFFFFF",
		"--text-color":       "#1E293B",
		"--border-color":     "#E2E8F0",
		"--heading-font":     "Inter, sans-serif",
		"--body-font":        "Inter, sans-serif",
		"--border-radius":    "0.375rem",
		"--shadow-sm":        "0 1px 2px 0 rgba(0, 0, 0, 0.05)",
		"--shadow-md":        "0 4px 6px -1px rgba(0, 0, 0, 0.1)",
		"--shadow-lg":        "0 10px 15px -3px rgba(0, 0, 0, 0.1)",
		"--spacing-xs":       "0.25rem",
		"--spacing-sm":       "0.5rem",
		"--spacing-md":       "1rem",
		"--spacing-lg":       "1.5rem",
		"--spacing-xl":       "3rem",
	}

	// TODO: Extract actual values from theme when theme structure is available
	// This would involve parsing the theme object and mapping values to CSS variables

	return variables
}
