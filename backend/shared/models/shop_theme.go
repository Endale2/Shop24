package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// =============== Theme Foundation Models ===============

// SectionConfig defines a single theme section (hero, product-list, footer, etc.)
type SectionConfig struct {
	ID       string            `bson:"id" json:"id"`                                 // unique section identifier (e.g., "hero", "product-grid")
	Type     string            `bson:"type" json:"type"`                             // section type (hero, product-list, banner, etc.)
	Settings map[string]string `bson:"settings,omitempty" json:"settings,omitempty"` // key-value pairs for section-specific settings
}

// ThemeConfig contains the customizable theme configuration (legacy compatibility)
type ThemeConfig struct {
	// Color scheme settings
	Colors map[string]string `bson:"colors,omitempty" json:"colors,omitempty"` // e.g., {"primary": "#007bff", "secondary": "#6c757d"}

	// Typography settings
	Fonts map[string]string `bson:"fonts,omitempty" json:"fonts,omitempty"` // e.g., {"heading": "Inter", "body": "Open Sans"}

	// Layout configurations
	Layouts map[string]string `bson:"layouts,omitempty" json:"layouts,omitempty"` // e.g., {"header_style": "minimal"}

	// Section configurations
	Sections []SectionConfig `bson:"sections,omitempty" json:"sections,omitempty"` // array of configured sections
}

// Theme represents a base theme definition that can be prebuilt or custom (marketplace themes)
type Theme struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name" json:"name"`       // theme display name
	Author  string             `bson:"author" json:"author"`   // theme creator/author
	Version string             `bson:"version" json:"version"` // semantic version (e.g., "1.0.0")

	// Theme availability
	IsPublic bool `bson:"isPublic" json:"isPublic"` // true for marketplace themes, false for custom/private themes

	// Theme configuration
	Config ThemeConfig `bson:"config" json:"config"` // default theme configuration

	// Optional metadata
	Description string   `bson:"description,omitempty" json:"description,omitempty"` // theme description
	Tags        []string `bson:"tags,omitempty" json:"tags,omitempty"`               // theme tags for categorization
	PreviewURL  string   `bson:"previewUrl,omitempty" json:"previewUrl,omitempty"`   // preview image/demo URL

	// Creator information (for custom themes)
	CreatedBy primitive.ObjectID `bson:"createdBy,omitempty" json:"createdBy,omitempty"` // seller who created this theme

	// Timestamps
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// Note: ThemePreset is now defined in theme_preset.go to avoid duplication

// =============== Scalable Shop Theme Models ===============

// ShopTheme represents a shop's theme configuration in a separate collection
// This provides better scalability and performance for theme operations
type ShopTheme struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID primitive.ObjectID `bson:"shopId" json:"shopId"` // Reference to shop

	// Theme Configuration - Optimized for frontend usage
	Colors    map[string]string      `bson:"colors,omitempty" json:"colors,omitempty"`       // Color palette (primary, secondary, background, etc.)
	Fonts     map[string]string      `bson:"fonts,omitempty" json:"fonts,omitempty"`         // Font configuration (heading, body)
	Layout    map[string]interface{} `bson:"layout,omitempty" json:"layout,omitempty"`       // Enhanced layout settings with breakpoints and spacing
	CustomCSS map[string]interface{} `bson:"customCSS,omitempty" json:"customCSS,omitempty"` // Enhanced CSS with validation and compilation
	SEO       map[string]interface{} `bson:"seo,omitempty" json:"seo,omitempty"`             // Enhanced SEO with structured data
	Mobile    map[string]interface{} `bson:"mobile,omitempty" json:"mobile,omitempty"`       // Enhanced mobile configuration

	// Advanced Styling Options
	Gradients  map[string]string `bson:"gradients,omitempty" json:"gradients,omitempty"`   // Gradient definitions (hero, button, card backgrounds)
	Shadows    map[string]string `bson:"shadows,omitempty" json:"shadows,omitempty"`       // Shadow configurations (card, button, image shadows)
	Animations map[string]string `bson:"animations,omitempty" json:"animations,omitempty"` // Animation settings (hover, transition, loading)
	Spacing    map[string]string `bson:"spacing,omitempty" json:"spacing,omitempty"`       // Spacing configurations (padding, margin, gaps)

	// Component-Specific Styling
	Components map[string]interface{} `bson:"components,omitempty" json:"components,omitempty"` // Component variants and styling options

	// Theme Metadata
	Name        string `bson:"name,omitempty" json:"name,omitempty"`               // Theme name/label
	Description string `bson:"description,omitempty" json:"description,omitempty"` // Theme description
	Version     string `bson:"version" json:"version"`                             // Theme version (e.g., "1.2.3")
	IsActive    bool   `bson:"isActive" json:"isActive"`                           // Whether this theme is currently active
	IsDefault   bool   `bson:"isDefault" json:"isDefault"`                         // Whether this is a default theme

	// Advanced Theme Features
	ParentThemeID primitive.ObjectID `bson:"parentThemeId,omitempty" json:"parentThemeId,omitempty"` // Reference to base theme (for inheritance)
	Tags          []string           `bson:"tags,omitempty" json:"tags,omitempty"`                   // Theme tags for categorization
	Backup        *ThemeBackup       `bson:"backup,omitempty" json:"backup,omitempty"`               // Previous version backup

	// Performance & Caching
	CompiledCSS    string    `bson:"compiledCSS,omitempty" json:"compiledCSS,omitempty"`       // Pre-compiled CSS for performance
	PreviewImage   string    `bson:"previewImage,omitempty" json:"previewImage,omitempty"`     // Theme preview image URL
	CacheKey       string    `bson:"cacheKey,omitempty" json:"cacheKey,omitempty"`             // Cache invalidation key
	LastCompiled   time.Time `bson:"lastCompiled,omitempty" json:"lastCompiled,omitempty"`     // When CSS was last compiled
	ConfigChecksum string    `bson:"configChecksum,omitempty" json:"configChecksum,omitempty"` // Checksum for change detection

	// Audit & Tracking
	CreatedBy  primitive.ObjectID `bson:"createdBy" json:"createdBy"`                       // Seller who created this theme
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`                       // Creation timestamp
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`                       // Last update timestamp
	LastUsedAt time.Time          `bson:"lastUsedAt,omitempty" json:"lastUsedAt,omitempty"` // When theme was last applied
	UsageCount int                `bson:"usageCount,omitempty" json:"usageCount,omitempty"` // How many times theme was applied

	// Analytics & Performance Metrics
	LoadTime    float64 `bson:"loadTime,omitempty" json:"loadTime,omitempty"`       // Average page load time with this theme
	Conversions int     `bson:"conversions,omitempty" json:"conversions,omitempty"` // Sales conversions with this theme
	Views       int     `bson:"views,omitempty" json:"views,omitempty"`             // Page views with this theme
}

// ThemeBackup stores previous theme configuration for rollback
type ThemeBackup struct {
	Version   string                 `bson:"version" json:"version"`
	Colors    map[string]string      `bson:"colors,omitempty" json:"colors,omitempty"`
	Fonts     map[string]string      `bson:"fonts,omitempty" json:"fonts,omitempty"`
	Layout    map[string]interface{} `bson:"layout,omitempty" json:"layout,omitempty"`
	CustomCSS map[string]interface{} `bson:"customCSS,omitempty" json:"customCSS,omitempty"`
	BackupAt  time.Time              `bson:"backupAt" json:"backupAt"`
}

// ShopThemeConfig represents the aggregated theme configuration for API responses
type ShopThemeConfig struct {
	Colors     map[string]string      `json:"colors"`
	Fonts      map[string]string      `json:"fonts"`
	Layout     map[string]interface{} `json:"layout"`
	CustomCSS  map[string]interface{} `json:"customCSS"`
	SEO        map[string]interface{} `json:"seo"`
	Mobile     map[string]interface{} `json:"mobile"`
	Gradients  map[string]string      `json:"gradients,omitempty"`
	Shadows    map[string]string      `json:"shadows,omitempty"`
	Animations map[string]string      `json:"animations,omitempty"`
	Spacing    map[string]string      `json:"spacing,omitempty"`
	Components map[string]interface{} `json:"components,omitempty"`
	Metadata   ThemeMetadata          `json:"metadata"`
}

// ThemeMetadata contains theme information for frontend
type ThemeMetadata struct {
	ID               primitive.ObjectID `json:"id"`
	Name             string             `json:"name"`
	Version          string             `json:"version"`
	IsActive         bool               `json:"isActive"`
	LastModified     time.Time          `json:"lastModified"`
	PreviewImage     string             `json:"previewImage,omitempty"`
	LoadTime         float64            `json:"loadTime,omitempty"`
	PerformanceScore int                `json:"performanceScore,omitempty"`
}

// GetDefaultShopTheme returns a default theme configuration
func GetDefaultShopTheme(shopID primitive.ObjectID, createdBy primitive.ObjectID) *ShopTheme {
	return &ShopTheme{
		ID:          primitive.NewObjectID(),
		ShopID:      shopID,
		CreatedBy:   createdBy,
		Name:        "Default Theme",
		Description: "Default shop theme with modern styling",
		Version:     "1.0.0",
		IsActive:    true,
		IsDefault:   true,
		Colors: map[string]string{
			"primary":    "#10B981",
			"secondary":  "#F59E0B",
			"background": "#FFFFFF",
			"heading":    "#1F2937",
			"bodyText":   "#6B7280",
			"border":     "#E5E7EB",
		},
		Fonts: map[string]string{
			"heading": "Inter",
			"body":    "Inter",
		},
		Layout: map[string]interface{}{
			"headerStyle":     "classic",
			"containerWidth":  "boxed",
			"sidebarPosition": "none",
			"gridColumns":     "3",
			"borderStyle":     "rounded",
			"spacing":         "normal",
			"breakpoints": map[string]interface{}{
				"mobile": map[string]interface{}{
					"gridColumns":    "1",
					"containerWidth": "full",
					"spacing":        "compact",
				},
				"tablet": map[string]interface{}{
					"gridColumns":    "2",
					"containerWidth": "boxed",
					"spacing":        "normal",
				},
				"desktop": map[string]interface{}{
					"gridColumns":    "3",
					"containerWidth": "boxed",
					"spacing":        "normal",
				},
			},
			"spacingConfig": map[string]interface{}{
				"containerPadding": "16px",
				"sectionGap":       "32px",
				"componentGap":     "16px",
				"cardPadding":      "24px",
			},
		},

		// Enhanced styling defaults
		Gradients: map[string]string{
			"hero":       "linear-gradient(135deg, #10B981 0%, #059669 100%)",
			"button":     "linear-gradient(135deg, #10B981 0%, #047857 100%)",
			"card":       "linear-gradient(135deg, #F9FAFB 0%, #F3F4F6 100%)",
			"background": "linear-gradient(135deg, #FFFFFF 0%, #F9FAFB 100%)",
		},

		Shadows: map[string]string{
			"card":   "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)",
			"button": "0 2px 4px -1px rgba(0, 0, 0, 0.1), 0 1px 2px -1px rgba(0, 0, 0, 0.06)",
			"image":  "0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)",
			"hover":  "0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)",
		},

		Animations: map[string]string{
			"transition": "all 0.3s ease",
			"hover":      "transform 0.2s ease, box-shadow 0.2s ease",
			"loading":    "opacity 0.5s ease",
			"fadeIn":     "opacity 0.3s ease-in-out",
		},

		Spacing: map[string]string{
			"xs":  "0.25rem",
			"sm":  "0.5rem",
			"md":  "1rem",
			"lg":  "1.5rem",
			"xl":  "2rem",
			"2xl": "3rem",
		},

		Components: map[string]interface{}{
			"productCard": map[string]interface{}{
				"style":       "modern",
				"hoverEffect": "lift",
				"imageRatio":  "square",
				"showBadges":  true,
				"showRating":  true,
			},
			"buttons": map[string]interface{}{
				"style":     "rounded",
				"size":      "medium",
				"animation": "subtle",
			},
			"navigation": map[string]interface{}{
				"style":       "horizontal",
				"position":    "top",
				"transparent": false,
				"sticky":      true,
			},
		},

		SEO: map[string]interface{}{
			"meta": map[string]interface{}{
				"title":       "",
				"description": "",
				"keywords":    []string{},
			},
			"openGraph": map[string]interface{}{
				"title":       "",
				"description": "",
				"image":       "",
				"type":        "website",
			},
			"structuredData": map[string]interface{}{
				"type":        "Store",
				"name":        "",
				"description": "",
			},
		},
		Mobile: map[string]interface{}{
			"navigation": map[string]interface{}{
				"style":    "hamburger",
				"position": "left",
				"overlay":  true,
			},
			"performance": map[string]interface{}{
				"imageQuality":    80,
				"lazyLoading":     true,
				"preloadCritical": true,
			},
			"interactions": map[string]interface{}{
				"touchFeedback":   true,
				"swipeNavigation": true,
				"pullToRefresh":   false,
				"stickyHeader":    true,
			},
		},
		CustomCSS: map[string]interface{}{
			"raw":       "",
			"compiled":  "",
			"variables": map[string]interface{}{},
			"validation": map[string]interface{}{
				"isValid":  true,
				"errors":   []string{},
				"warnings": []string{},
			},
		},
		Tags:       []string{"default", "modern", "responsive"},
		UsageCount: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// ToConfig converts ShopTheme to ShopThemeConfig for API responses
func (st *ShopTheme) ToConfig() ShopThemeConfig {
	return ShopThemeConfig{
		Colors:     st.Colors,
		Fonts:      st.Fonts,
		Layout:     st.Layout,
		CustomCSS:  st.CustomCSS,
		SEO:        st.SEO,
		Mobile:     st.Mobile,
		Gradients:  st.Gradients,
		Shadows:    st.Shadows,
		Animations: st.Animations,
		Spacing:    st.Spacing,
		Components: st.Components,
		Metadata: ThemeMetadata{
			ID:           st.ID,
			Name:         st.Name,
			Version:      st.Version,
			IsActive:     st.IsActive,
			LastModified: st.UpdatedAt,
			PreviewImage: st.PreviewImage,
			LoadTime:     st.LoadTime,
		},
	}
}

// CreateBackup creates a backup of current theme configuration
func (st *ShopTheme) CreateBackup() {
	st.Backup = &ThemeBackup{
		Version:   st.Version,
		Colors:    st.Colors,
		Fonts:     st.Fonts,
		Layout:    st.Layout,
		CustomCSS: st.CustomCSS,
		BackupAt:  time.Now(),
	}
}

// RestoreFromBackup restores theme from backup
func (st *ShopTheme) RestoreFromBackup() bool {
	if st.Backup == nil {
		return false
	}

	st.Colors = st.Backup.Colors
	st.Fonts = st.Backup.Fonts
	st.Layout = st.Backup.Layout
	st.CustomCSS = st.Backup.CustomCSS
	st.UpdatedAt = time.Now()

	return true
}

// =============== Helper Functions for Interface{} Types ===============

// GetLayoutValue safely extracts string values from layout map
func (st *ShopTheme) GetLayoutValue(key, defaultValue string) string {
	if st.Layout == nil {
		return defaultValue
	}
	if value, ok := st.Layout[key]; ok {
		if strValue, ok := value.(string); ok {
			return strValue
		}
	}
	return defaultValue
}

// GetCustomCSSRaw safely extracts raw CSS string
func (st *ShopTheme) GetCustomCSSRaw() string {
	if st.CustomCSS == nil {
		return ""
	}
	if raw, ok := st.CustomCSS["raw"]; ok {
		if strValue, ok := raw.(string); ok {
			return strValue
		}
	}
	return ""
}

// GetCustomCSSCompiled safely extracts compiled CSS string
func (st *ShopTheme) GetCustomCSSCompiled() string {
	if st.CustomCSS == nil {
		return ""
	}
	if compiled, ok := st.CustomCSS["compiled"]; ok {
		if strValue, ok := compiled.(string); ok {
			return strValue
		}
	}
	return ""
}

// SetCustomCSSRaw safely sets raw CSS string
func (st *ShopTheme) SetCustomCSSRaw(css string) {
	if st.CustomCSS == nil {
		st.CustomCSS = make(map[string]interface{})
	}
	st.CustomCSS["raw"] = css
}

// SetCustomCSSCompiled safely sets compiled CSS string
func (st *ShopTheme) SetCustomCSSCompiled(css string) {
	if st.CustomCSS == nil {
		st.CustomCSS = make(map[string]interface{})
	}
	st.CustomCSS["compiled"] = css
}

// IncrementVersion increments the theme version
func (st *ShopTheme) IncrementVersion() {
	// Simple version increment (1.0.0 -> 1.0.1)
	switch st.Version {
	case "":
		st.Version = "1.0.0"
	default:
		// For simplicity, just append timestamp
		st.Version = time.Now().Format("1.0.20060102150405")
	}
}
