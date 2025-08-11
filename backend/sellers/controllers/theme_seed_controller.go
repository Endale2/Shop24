package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/sellers/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SeedDefaultThemes creates default marketplace themes
// POST /seller/themes/seed (for development/admin use)
func SeedDefaultThemes(c *gin.Context) {
	// Default themes to seed
	defaultThemes := []models.Theme{
		{
			Name:        "Classic Commerce",
			Author:      "DRPS Team",
			Version:     "1.0.0",
			IsPublic:    true,
			Description: "A classic, professional theme for any type of business",
			Tags:        []string{"classic", "professional", "versatile"},
			Config: models.ThemeConfig{
				Colors: map[string]string{
					"primary":     "#3B82F6",
					"secondary":   "#06B6D4",
					"background":  "#FFFFFF",
					"heading":     "#1E3A8A",
					"bodyText":    "#475569",
				},
				Fonts: map[string]string{
					"heading": "Inter",
					"body":    "Inter",
				},
				Layouts: map[string]string{
					"headerStyle":      "classic",
					"containerWidth":   "boxed",
					"sidebarPosition":  "none",
					"gridColumns":      "3",
				},
			},
		},
		{
			Name:        "Modern Minimalist",
			Author:      "DRPS Team",
			Version:     "1.0.0",
			IsPublic:    true,
			Description: "Clean, modern design with minimal elements",
			Tags:        []string{"modern", "minimal", "clean"},
			Config: models.ThemeConfig{
				Colors: map[string]string{
					"primary":     "#10B981",
					"secondary":   "#34D399",
					"background":  "#FFFFFF",
					"heading":     "#064E3B",
					"bodyText":    "#374151",
				},
				Fonts: map[string]string{
					"heading": "Poppins",
					"body":    "Inter",
				},
				Layouts: map[string]string{
					"headerStyle":      "minimal",
					"containerWidth":   "full",
					"sidebarPosition":  "none",
					"gridColumns":      "4",
				},
			},
		},
		{
			Name:        "Bold & Vibrant",
			Author:      "DRPS Team",
			Version:     "1.0.0",
			IsPublic:    true,
			Description: "Eye-catching design with vibrant colors",
			Tags:        []string{"bold", "vibrant", "colorful"},
			Config: models.ThemeConfig{
				Colors: map[string]string{
					"primary":     "#F59E0B",
					"secondary":   "#FB923C",
					"background":  "#FFFFFF",
					"heading":     "#92400E",
					"bodyText":    "#6B7280",
				},
				Fonts: map[string]string{
					"heading": "Montserrat",
					"body":    "Open Sans",
				},
				Layouts: map[string]string{
					"headerStyle":      "centered",
					"containerWidth":   "boxed",
					"sidebarPosition":  "right",
					"gridColumns":      "3",
				},
			},
		},
		{
			Name:        "Dark Professional",
			Author:      "DRPS Team",
			Version:     "1.0.0",
			IsPublic:    true,
			Description: "Professional dark theme for modern brands",
			Tags:        []string{"dark", "professional", "modern"},
			Config: models.ThemeConfig{
				Colors: map[string]string{
					"primary":     "#6366F1",
					"secondary":   "#F59E0B",
					"background":  "#111827",
					"heading":     "#F9FAFB",
					"bodyText":    "#D1D5DB",
				},
				Fonts: map[string]string{
					"heading": "Inter",
					"body":    "Inter",
				},
				Layouts: map[string]string{
					"headerStyle":      "classic",
					"containerWidth":   "boxed",
					"sidebarPosition":  "none",
					"gridColumns":      "3",
				},
			},
		},
		{
			Name:        "Elegant Fashion",
			Author:      "DRPS Team",
			Version:     "1.0.0",
			IsPublic:    true,
			Description: "Sophisticated theme perfect for fashion and lifestyle brands",
			Tags:        []string{"elegant", "fashion", "sophisticated"},
			Config: models.ThemeConfig{
				Colors: map[string]string{
					"primary":     "#EC4899",
					"secondary":   "#F472B6",
					"background":  "#FFFFFF",
					"heading":     "#BE185D",
					"bodyText":    "#6B7280",
				},
				Fonts: map[string]string{
					"heading": "Playfair Display",
					"body":    "Source Sans Pro",
				},
				Layouts: map[string]string{
					"headerStyle":      "centered",
					"containerWidth":   "narrow",
					"sidebarPosition":  "none",
					"gridColumns":      "2",
				},
			},
		},
	}

	var createdCount = 0
	var errors []string

	// Create each theme
	for _, theme := range defaultThemes {
		_, err := services.CreateThemeService(&theme, primitive.NilObjectID)
		if err != nil {
			errors = append(errors, "Failed to create theme '"+theme.Name+"': "+err.Error())
		} else {
			createdCount++
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusPartialContent, gin.H{
			"message":      "Some themes created successfully",
			"created":      createdCount,
			"total":        len(defaultThemes),
			"errors":       errors,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "All default themes created successfully",
		"created": createdCount,
		"total":   len(defaultThemes),
	})
}
