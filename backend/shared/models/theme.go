package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SectionConfig defines a single theme section (hero, product-list, footer, etc.)
type SectionConfig struct {
	ID       string            `bson:"id" json:"id"`                           // unique section identifier (e.g., "hero", "product-grid")
	Type     string            `bson:"type" json:"type"`                       // section type (hero, product-list, banner, etc.)
	Settings map[string]string `bson:"settings,omitempty" json:"settings,omitempty"` // key-value pairs for section-specific settings
}

// ThemeConfig contains the customizable theme configuration
type ThemeConfig struct {
	// Color scheme settings
	Colors map[string]string `bson:"colors,omitempty" json:"colors,omitempty"` // e.g., {"primary": "#007bff", "secondary": "#6c757d", "background": "#ffffff"}
	
	// Typography settings
	Fonts map[string]string `bson:"fonts,omitempty" json:"fonts,omitempty"` // e.g., {"heading": "Inter", "body": "Open Sans", "accent": "Playfair Display"}
	
	// Layout configurations
	Layouts map[string]string `bson:"layouts,omitempty" json:"layouts,omitempty"` // e.g., {"header_style": "minimal", "footer_style": "full", "product_grid": "3-column"}
	
	// Section configurations
	Sections []SectionConfig `bson:"sections,omitempty" json:"sections,omitempty"` // array of configured sections
}

// Theme represents a base theme definition that can be prebuilt or custom
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

// ShopTheme links a shop to a theme with customizations
type ShopTheme struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID  primitive.ObjectID `bson:"shopId" json:"shopId"`   // reference to Shop
	ThemeID primitive.ObjectID `bson:"themeId" json:"themeId"` // reference to Theme
	
	// Theme customizations/overrides
	Overrides ThemeConfig `bson:"overrides,omitempty" json:"overrides,omitempty"` // shop-specific customizations that override theme defaults
	
	// Publication status
	Published bool `bson:"published" json:"published"` // whether this theme configuration is active/published
	
	// Optional metadata
	CustomName string `bson:"customName,omitempty" json:"customName,omitempty"` // shop owner's custom name for this theme setup
	
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
