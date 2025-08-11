package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// =============== Theme Foundation Models ===============

// SectionConfig defines a single theme section (hero, product-list, footer, etc.)
type SectionConfig struct {
	ID       string            `bson:"id" json:"id"`                                     // unique section identifier (e.g., "hero", "product-grid")
	Type     string            `bson:"type" json:"type"`                                 // section type (hero, product-list, banner, etc.)
	Settings map[string]string `bson:"settings,omitempty" json:"settings,omitempty"`   // key-value pairs for section-specific settings
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
	Name    string             `bson:"name" json:"name"`                           // theme display name
	Author  string             `bson:"author" json:"author"`                       // theme creator/author
	Version string             `bson:"version" json:"version"`                     // semantic version (e.g., "1.0.0")
	
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

// ThemePreset represents predefined theme configurations for quick setup
type ThemePreset struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ThemeID     primitive.ObjectID `bson:"themeId" json:"themeId"`         // base theme this preset belongs to
	Name        string             `bson:"name" json:"name"`               // preset name (e.g., "Dark Mode", "Minimalist", "Bold Colors")
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	
	// Preset configuration
	Config ThemeConfig `bson:"config" json:"config"` // preconfigured settings
	
	// Preset metadata
	IsDefault   bool     `bson:"isDefault,omitempty" json:"isDefault,omitempty"`     // whether this is the default preset for the theme
	Tags        []string `bson:"tags,omitempty" json:"tags,omitempty"`               // preset tags
	PreviewURL  string   `bson:"previewUrl,omitempty" json:"previewUrl,omitempty"`   // preview image
	
	// Timestamps
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// =============== Scalable Shop Theme Models ===============

// ShopTheme represents a shop's theme configuration in a separate collection
// This provides better scalability and performance for theme operations
type ShopTheme struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID primitive.ObjectID `bson:"shopId" json:"shopId"` // Reference to shop

	// Theme Configuration - Optimized for frontend usage
	Colors    map[string]string `bson:"colors,omitempty" json:"colors,omitempty"`       // Color palette (primary, secondary, background, etc.)
	Fonts     map[string]string `bson:"fonts,omitempty" json:"fonts,omitempty"`         // Font configuration (heading, body)
	Layout    map[string]string `bson:"layout,omitempty" json:"layout,omitempty"`       // Layout settings (containerWidth, headerStyle, etc.)
	CustomCSS string            `bson:"customCSS,omitempty" json:"customCSS,omitempty"` // Custom CSS code
	SEO       map[string]string `bson:"seo,omitempty" json:"seo,omitempty"`             // SEO settings
	Mobile    map[string]string `bson:"mobile,omitempty" json:"mobile,omitempty"`       // Mobile-specific settings

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
	CompiledCSS    string            `bson:"compiledCSS,omitempty" json:"compiledCSS,omitempty"`       // Pre-compiled CSS for performance
	PreviewImage   string            `bson:"previewImage,omitempty" json:"previewImage,omitempty"`     // Theme preview image URL
	CacheKey       string            `bson:"cacheKey,omitempty" json:"cacheKey,omitempty"`             // Cache invalidation key
	LastCompiled   time.Time         `bson:"lastCompiled,omitempty" json:"lastCompiled,omitempty"`     // When CSS was last compiled
	ConfigChecksum string            `bson:"configChecksum,omitempty" json:"configChecksum,omitempty"` // Checksum for change detection

	// Audit & Tracking
	CreatedBy   primitive.ObjectID `bson:"createdBy" json:"createdBy"`     // Seller who created this theme
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`     // Creation timestamp
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`     // Last update timestamp
	LastUsedAt  time.Time          `bson:"lastUsedAt,omitempty" json:"lastUsedAt,omitempty"` // When theme was last applied
	UsageCount  int                `bson:"usageCount,omitempty" json:"usageCount,omitempty"` // How many times theme was applied
	
	// Analytics & Performance Metrics
	LoadTime    float64 `bson:"loadTime,omitempty" json:"loadTime,omitempty"`       // Average page load time with this theme
	Conversions int     `bson:"conversions,omitempty" json:"conversions,omitempty"` // Sales conversions with this theme
	Views       int     `bson:"views,omitempty" json:"views,omitempty"`             // Page views with this theme
}

// ThemeBackup stores previous theme configuration for rollback
type ThemeBackup struct {
	Version   string            `bson:"version" json:"version"`
	Colors    map[string]string `bson:"colors,omitempty" json:"colors,omitempty"`
	Fonts     map[string]string `bson:"fonts,omitempty" json:"fonts,omitempty"`
	Layout    map[string]string `bson:"layout,omitempty" json:"layout,omitempty"`
	CustomCSS string            `bson:"customCSS,omitempty" json:"customCSS,omitempty"`
	BackupAt  time.Time         `bson:"backupAt" json:"backupAt"`
}

// ShopThemeConfig represents the aggregated theme configuration for API responses
type ShopThemeConfig struct {
	Colors    map[string]string `json:"colors"`
	Fonts     map[string]string `json:"fonts"`
	Layout    map[string]string `json:"layout"`
	CustomCSS string            `json:"customCSS"`
	SEO       map[string]string `json:"seo"`
	Mobile    map[string]string `json:"mobile"`
	Metadata  ThemeMetadata     `json:"metadata"`
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
			"primary":     "#10B981",
			"secondary":   "#F59E0B",
			"background":  "#FFFFFF",
			"heading":     "#1F2937",
			"bodyText":    "#6B7280",
		},
		Fonts: map[string]string{
			"heading": "Inter",
			"body":    "Inter",
		},
		Layout: map[string]string{
			"headerStyle":     "classic",
			"containerWidth":  "boxed",
			"sidebarPosition": "none",
			"gridColumns":     "3",
		},
		SEO: map[string]string{
			"metaTitle":       "",
			"metaDescription": "",
			"keywords":        "",
		},
		Mobile: map[string]string{
			"enabled":        "true",
			"responsive":     "true",
			"touchOptimized": "true",
		},
		CustomCSS:    "",
		Tags:         []string{"default", "modern", "responsive"},
		UsageCount:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// ToConfig converts ShopTheme to ShopThemeConfig for API responses
func (st *ShopTheme) ToConfig() ShopThemeConfig {
	return ShopThemeConfig{
		Colors:    st.Colors,
		Fonts:     st.Fonts,
		Layout:    st.Layout,
		CustomCSS: st.CustomCSS,
		SEO:       st.SEO,
		Mobile:    st.Mobile,
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
