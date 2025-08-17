package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ComponentTemplate represents a reusable component template for the drag-and-drop builder
type ComponentTemplate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`                               // Component name (e.g., "Hero Banner", "Product Grid")
	Type        string             `bson:"type" json:"type"`                               // Component type (hero, product-grid, testimonials, etc.)
	Category    string             `bson:"category" json:"category"`                       // Category (layout, content, interactive, marketing)
	Description string             `bson:"description,omitempty" json:"description,omitempty"` // Component description
	
	// Component Configuration
	Settings    map[string]interface{} `bson:"settings,omitempty" json:"settings,omitempty"`       // Default settings for the component
	StyleConfig map[string]interface{} `bson:"styleConfig,omitempty" json:"styleConfig,omitempty"` // Style configuration options
	Content     map[string]interface{} `bson:"content,omitempty" json:"content,omitempty"`         // Default content structure
	
	// Visual & Preview
	Preview     string   `bson:"preview,omitempty" json:"preview,omitempty"`         // Preview image URL
	Icon        string   `bson:"icon,omitempty" json:"icon,omitempty"`               // Icon identifier
	Tags        []string `bson:"tags,omitempty" json:"tags,omitempty"`               // Tags for filtering and search
	
	// Responsive Configuration
	Responsive map[string]interface{} `bson:"responsive,omitempty" json:"responsive,omitempty"` // Responsive behavior settings
	
	// Permissions & Visibility
	IsPublic    bool               `bson:"isPublic" json:"isPublic"`                           // Whether component is publicly available
	IsPremium   bool               `bson:"isPremium,omitempty" json:"isPremium,omitempty"`     // Whether component requires premium access
	CreatedBy   primitive.ObjectID `bson:"createdBy,omitempty" json:"createdBy,omitempty"`     // Creator (seller) ID
	
	// Metadata
	Version     string    `bson:"version,omitempty" json:"version,omitempty"`         // Component version
	UsageCount  int       `bson:"usageCount,omitempty" json:"usageCount,omitempty"`   // How many times component has been used
	Rating      float64   `bson:"rating,omitempty" json:"rating,omitempty"`           // Average rating
	ReviewCount int       `bson:"reviewCount,omitempty" json:"reviewCount,omitempty"` // Number of reviews
	
	// Timestamps
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// ComponentInstance represents an instance of a component in a shop's layout
type ComponentInstance struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TemplateID   primitive.ObjectID `bson:"templateId" json:"templateId"`                   // Reference to ComponentTemplate
	ShopID       primitive.ObjectID `bson:"shopId" json:"shopId"`                           // Shop this instance belongs to
	PageType     string             `bson:"pageType" json:"pageType"`                       // Page where component is used (home, products, etc.)
	Position     int                `bson:"position" json:"position"`                       // Position in the layout
	
	// Instance Configuration
	Settings     map[string]interface{} `bson:"settings,omitempty" json:"settings,omitempty"`         // Instance-specific settings
	Content      map[string]interface{} `bson:"content,omitempty" json:"content,omitempty"`           // Instance-specific content
	StyleOverrides map[string]interface{} `bson:"styleOverrides,omitempty" json:"styleOverrides,omitempty"` // Style overrides
	
	// Responsive Overrides
	ResponsiveOverrides map[string]interface{} `bson:"responsiveOverrides,omitempty" json:"responsiveOverrides,omitempty"`
	
	// Visibility & Conditions
	IsVisible   bool                   `bson:"isVisible" json:"isVisible"`                           // Whether component is visible
	Conditions  map[string]interface{} `bson:"conditions,omitempty" json:"conditions,omitempty"`     // Display conditions
	
	// Metadata
	CreatedBy primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// PageLayout represents the layout configuration for a specific page type
type PageLayout struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID   primitive.ObjectID `bson:"shopId" json:"shopId"`
	PageType string             `bson:"pageType" json:"pageType"` // home, products, product, collections, etc.
	
	// Layout Configuration
	Components []primitive.ObjectID   `bson:"components" json:"components"`                     // Ordered list of component instance IDs
	Settings   map[string]interface{} `bson:"settings,omitempty" json:"settings,omitempty"`     // Page-specific settings
	
	// Responsive Layout
	ResponsiveSettings map[string]interface{} `bson:"responsiveSettings,omitempty" json:"responsiveSettings,omitempty"`
	
	// Metadata
	Name        string `bson:"name,omitempty" json:"name,omitempty"`               // Layout name
	Description string `bson:"description,omitempty" json:"description,omitempty"` // Layout description
	IsActive    bool   `bson:"isActive" json:"isActive"`                           // Whether this layout is active
	Version     string `bson:"version,omitempty" json:"version,omitempty"`         // Layout version
	
	// Audit
	CreatedBy primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// GetDefaultComponentTemplates returns a set of default component templates
func GetDefaultComponentTemplates() []ComponentTemplate {
	now := time.Now()
	
	return []ComponentTemplate{
		{
			ID:          primitive.NewObjectID(),
			Name:        "Hero Banner",
			Type:        "hero",
			Category:    "layout",
			Description: "Large banner with image, title, and call-to-action",
			Settings: map[string]interface{}{
				"height":          "400px",
				"backgroundType":  "image",
				"overlayOpacity":  0.3,
				"textAlignment":   "center",
				"showButton":      true,
			},
			Content: map[string]interface{}{
				"title":       "Welcome to Our Store",
				"subtitle":    "Discover amazing products",
				"buttonText":  "Shop Now",
				"buttonLink":  "/products",
				"backgroundImage": "",
			},
			Responsive: map[string]interface{}{
				"mobile": map[string]interface{}{
					"height": "300px",
					"textAlignment": "center",
				},
			},
			IsPublic:   true,
			Version:    "1.0.0",
			Tags:       []string{"banner", "hero", "landing"},
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Product Grid",
			Type:        "product-grid",
			Category:    "content",
			Description: "Grid layout for displaying products",
			Settings: map[string]interface{}{
				"columns":     3,
				"showRating":  true,
				"showBadges":  true,
				"showPrice":   true,
				"imageRatio":  "square",
				"hoverEffect": "lift",
			},
			Responsive: map[string]interface{}{
				"mobile": map[string]interface{}{
					"columns": 1,
				},
				"tablet": map[string]interface{}{
					"columns": 2,
				},
			},
			IsPublic:   true,
			Version:    "1.0.0",
			Tags:       []string{"products", "grid", "catalog"},
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Testimonials",
			Type:        "testimonials",
			Category:    "marketing",
			Description: "Customer testimonials and reviews",
			Settings: map[string]interface{}{
				"layout":      "carousel",
				"showRating":  true,
				"showAvatar":  true,
				"autoplay":    true,
				"interval":    5000,
			},
			Content: map[string]interface{}{
				"testimonials": []map[string]interface{}{
					{
						"name":    "John Doe",
						"rating":  5,
						"text":    "Amazing products and great service!",
						"avatar":  "",
					},
				},
			},
			IsPublic:   true,
			Version:    "1.0.0",
			Tags:       []string{"testimonials", "reviews", "social-proof"},
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Newsletter Signup",
			Type:        "newsletter",
			Category:    "interactive",
			Description: "Email newsletter subscription form",
			Settings: map[string]interface{}{
				"style":       "inline",
				"showIcon":    true,
				"placeholder": "Enter your email",
				"buttonText":  "Subscribe",
			},
			Content: map[string]interface{}{
				"title":       "Stay Updated",
				"description": "Get the latest news and exclusive offers",
			},
			IsPublic:   true,
			Version:    "1.0.0",
			Tags:       []string{"newsletter", "email", "subscription"},
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}
}
