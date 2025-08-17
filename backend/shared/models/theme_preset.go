package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ThemePreset represents a predefined theme configuration that can be applied to shops
type ThemePreset struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`                               // Preset name (e.g., "Modern Minimalist", "Bold & Colorful")
	Category    string             `bson:"category" json:"category"`                       // Category (modern, classic, minimal, bold, etc.)
	Description string             `bson:"description,omitempty" json:"description,omitempty"` // Preset description
	
	// Theme Configuration
	Colors      map[string]string      `bson:"colors" json:"colors"`                       // Color palette
	Fonts       map[string]string      `bson:"fonts" json:"fonts"`                         // Font configuration
	Layout      map[string]interface{} `bson:"layout" json:"layout"`                       // Layout settings
	Gradients   map[string]string      `bson:"gradients,omitempty" json:"gradients,omitempty"`   // Gradient definitions
	Shadows     map[string]string      `bson:"shadows,omitempty" json:"shadows,omitempty"`       // Shadow configurations
	Animations  map[string]string      `bson:"animations,omitempty" json:"animations,omitempty"` // Animation settings
	Spacing     map[string]string      `bson:"spacing,omitempty" json:"spacing,omitempty"`       // Spacing configurations
	Components  map[string]interface{} `bson:"components,omitempty" json:"components,omitempty"` // Component styling
	
	// Visual & Preview
	Preview     string   `bson:"preview,omitempty" json:"preview,omitempty"`         // Preview image URL
	DemoURL     string   `bson:"demoURL,omitempty" json:"demoURL,omitempty"`         // Demo storefront URL
	Tags        []string `bson:"tags,omitempty" json:"tags,omitempty"`               // Tags for filtering
	
	// Permissions & Visibility
	IsPublic    bool               `bson:"isPublic" json:"isPublic"`                           // Whether preset is publicly available
	IsPremium   bool               `bson:"isPremium,omitempty" json:"isPremium,omitempty"`     // Whether preset requires premium access
	CreatedBy   primitive.ObjectID `bson:"createdBy,omitempty" json:"createdBy,omitempty"`     // Creator ID
	
	// Metadata
	Version     string  `bson:"version,omitempty" json:"version,omitempty"`         // Preset version
	UsageCount  int     `bson:"usageCount,omitempty" json:"usageCount,omitempty"`   // How many times preset has been applied
	Rating      float64 `bson:"rating,omitempty" json:"rating,omitempty"`           // Average rating
	ReviewCount int     `bson:"reviewCount,omitempty" json:"reviewCount,omitempty"` // Number of reviews
	
	// Timestamps
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// ThemeVersion represents a saved version of a theme for rollback purposes
type ThemeVersion struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID      primitive.ObjectID `bson:"shopId" json:"shopId"`                           // Shop this version belongs to
	ThemeID     primitive.ObjectID `bson:"themeId" json:"themeId"`                         // Reference to the theme
	Version     string             `bson:"version" json:"version"`                         // Version identifier (e.g., "1.2.3")
	VersionType string             `bson:"versionType" json:"versionType"`                 // Type (auto, manual, backup)
	
	// Theme Data Snapshot
	Data        ShopTheme `bson:"data" json:"data"`                                   // Complete theme data at this version
	ChangeLog   string    `bson:"changeLog,omitempty" json:"changeLog,omitempty"`     // Description of changes
	ChangedBy   primitive.ObjectID `bson:"changedBy" json:"changedBy"`                 // Who made the changes
	
	// Metadata
	IsActive    bool      `bson:"isActive,omitempty" json:"isActive,omitempty"`       // Whether this is the active version
	Size        int64     `bson:"size,omitempty" json:"size,omitempty"`               // Data size in bytes
	Checksum    string    `bson:"checksum,omitempty" json:"checksum,omitempty"`       // Data checksum for integrity
	
	// Timestamps
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

// GetDefaultThemePresets returns a set of default theme presets
func GetDefaultThemePresets() []ThemePreset {
	now := time.Now()
	
	return []ThemePreset{
		{
			ID:          primitive.NewObjectID(),
			Name:        "Modern Minimalist",
			Category:    "modern",
			Description: "Clean, minimal design with plenty of white space",
			Colors: map[string]string{
				"primary":    "#2563EB",
				"secondary":  "#64748B",
				"background": "#FFFFFF",
				"heading":    "#1E293B",
				"bodyText":   "#475569",
				"border":     "#E2E8F0",
			},
			Fonts: map[string]string{
				"heading": "Inter",
				"body":    "Inter",
			},
			Layout: map[string]interface{}{
				"headerStyle":     "minimal",
				"containerWidth":  "boxed",
				"sidebarPosition": "none",
				"gridColumns":     "3",
				"borderStyle":     "rounded",
				"spacing":         "generous",
			},
			Gradients: map[string]string{
				"hero":   "linear-gradient(135deg, #2563EB 0%, #1D4ED8 100%)",
				"button": "linear-gradient(135deg, #2563EB 0%, #1E40AF 100%)",
			},
			Shadows: map[string]string{
				"card":   "0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)",
				"button": "0 1px 2px 0 rgba(0, 0, 0, 0.05)",
			},
			IsPublic:  true,
			Version:   "1.0.0",
			Tags:      []string{"modern", "minimal", "clean", "professional"},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Bold & Vibrant",
			Category:    "bold",
			Description: "Eye-catching design with vibrant colors and strong contrasts",
			Colors: map[string]string{
				"primary":    "#DC2626",
				"secondary":  "#F59E0B",
				"background": "#FFFFFF",
				"heading":    "#1F2937",
				"bodyText":   "#374151",
				"border":     "#E5E7EB",
			},
			Fonts: map[string]string{
				"heading": "Poppins",
				"body":    "Open Sans",
			},
			Layout: map[string]interface{}{
				"headerStyle":     "bold",
				"containerWidth":  "wide",
				"sidebarPosition": "none",
				"gridColumns":     "4",
				"borderStyle":     "sharp",
				"spacing":         "normal",
			},
			Gradients: map[string]string{
				"hero":   "linear-gradient(135deg, #DC2626 0%, #B91C1C 100%)",
				"button": "linear-gradient(135deg, #F59E0B 0%, #D97706 100%)",
			},
			Shadows: map[string]string{
				"card":   "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)",
				"button": "0 2px 4px -1px rgba(0, 0, 0, 0.1)",
			},
			IsPublic:  true,
			Version:   "1.0.0",
			Tags:      []string{"bold", "vibrant", "colorful", "energetic"},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Classic Elegant",
			Category:    "classic",
			Description: "Timeless design with elegant typography and refined colors",
			Colors: map[string]string{
				"primary":    "#1F2937",
				"secondary":  "#6B7280",
				"background": "#F9FAFB",
				"heading":    "#111827",
				"bodyText":   "#374151",
				"border":     "#D1D5DB",
			},
			Fonts: map[string]string{
				"heading": "Playfair Display",
				"body":    "Source Sans Pro",
			},
			Layout: map[string]interface{}{
				"headerStyle":     "classic",
				"containerWidth":  "boxed",
				"sidebarPosition": "none",
				"gridColumns":     "3",
				"borderStyle":     "rounded",
				"spacing":         "normal",
			},
			Gradients: map[string]string{
				"hero":   "linear-gradient(135deg, #1F2937 0%, #111827 100%)",
				"button": "linear-gradient(135deg, #6B7280 0%, #4B5563 100%)",
			},
			Shadows: map[string]string{
				"card":   "0 2px 4px -1px rgba(0, 0, 0, 0.1), 0 1px 2px -1px rgba(0, 0, 0, 0.06)",
				"button": "0 1px 3px 0 rgba(0, 0, 0, 0.1)",
			},
			IsPublic:  true,
			Version:   "1.0.0",
			Tags:      []string{"classic", "elegant", "timeless", "sophisticated"},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Nature Inspired",
			Category:    "nature",
			Description: "Earthy tones and organic shapes inspired by nature",
			Colors: map[string]string{
				"primary":    "#059669",
				"secondary":  "#D97706",
				"background": "#F7F8F7",
				"heading":    "#1F2937",
				"bodyText":   "#374151",
				"border":     "#D1D5DB",
			},
			Fonts: map[string]string{
				"heading": "Merriweather",
				"body":    "Lato",
			},
			Layout: map[string]interface{}{
				"headerStyle":     "organic",
				"containerWidth":  "boxed",
				"sidebarPosition": "none",
				"gridColumns":     "3",
				"borderStyle":     "rounded",
				"spacing":         "generous",
			},
			Gradients: map[string]string{
				"hero":   "linear-gradient(135deg, #059669 0%, #047857 100%)",
				"button": "linear-gradient(135deg, #D97706 0%, #B45309 100%)",
			},
			Shadows: map[string]string{
				"card":   "0 2px 8px -2px rgba(0, 0, 0, 0.1), 0 1px 4px -1px rgba(0, 0, 0, 0.06)",
				"button": "0 2px 4px -1px rgba(0, 0, 0, 0.1)",
			},
			IsPublic:  true,
			Version:   "1.0.0",
			Tags:      []string{"nature", "organic", "earthy", "sustainable"},
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}
