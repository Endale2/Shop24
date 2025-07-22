package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Option represents a single variant option (e.g. Size: M).
type Option struct {
	Name  string `bson:"name"  json:"name"`
	Value string `bson:"value" json:"value"`
}

// Variant represents a specific version of a product.
type Variant struct {
	VariantID primitive.ObjectID `bson:"variant_id,omitempty" json:"id"`
	Options   []Option           `bson:"options"             json:"options"`

	// Price is the base price; DisplayPrice can override it in the UI
	Price float64 `bson:"price"                json:"price"`
	Stock int     `bson:"stock"                json:"stock"`
	Image string  `bson:"image,omitempty"      json:"image,omitempty"`

	// Discount display fields (not stored in DB)
	DisplayPrice      *float64 `bson:"-" json:"display_price,omitempty"`
	AppliedDiscountID *string  `bson:"-" json:"applied_discount_id,omitempty"`

	// Total could represent price * quantity, or any calculated total
	Total *float64 `bson:"total,omitempty"      json:"total,omitempty"`

	CreatedAt time.Time `bson:"createdAt,omitempty"  json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"  json:"updatedAt,omitempty"`
}

// Product represents the main product entity.
type Product struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"             json:"id"`
	ShopID primitive.ObjectID `bson:"shop_id,omitempty"         json:"shop_id"`
	UserID primitive.ObjectID `bson:"user_id,omitempty"         json:"user_id"`

	// Core product fields
	Name        string `bson:"name"                      json:"name"`
	Slug        string `bson:"slug"                      json:"slug"`
	Description string `bson:"description"               json:"description"`

	// SEO
	MetaTitle       string   `bson:"meta_title,omitempty"      json:"meta_title,omitempty"`
	MetaDescription string   `bson:"meta_description,omitempty" json:"meta_description,omitempty"`
	CanonicalURL    string   `bson:"canonical_url,omitempty"    json:"canonical_url,omitempty"`
	Tags            []string `bson:"tags,omitempty"            json:"tags,omitempty"`

	// Media
	MainImage string   `bson:"main_image,omitempty"      json:"main_image,omitempty"`
	Images    []string `bson:"images,omitempty"          json:"images,omitempty"`

	// Pricing & inventory
	CollectionID primitive.ObjectID `bson:"collection_id,omitempty" json:"collection_id,omitempty"`
	Price        float64            `bson:"price"                     json:"price"`
	Stock        int                `bson:"stock"                     json:"stock"`

	// Discount display fields (not stored in DB)
	DisplayPrice      *float64 `bson:"-" json:"display_price,omitempty"`
	AppliedDiscountID *string  `bson:"-" json:"applied_discount_id,omitempty"`

	// Variants & ratings
	Variants      []Variant `bson:"variants,omitempty"        json:"variants,omitempty"`
	AverageRating float64   `bson:"average_rating,omitempty"  json:"average_rating,omitempty"`
	ReviewCount   int       `bson:"review_count,omitempty"    json:"review_count,omitempty"`

	// Totals (e.g. sum of all variant totals, or other calculated field)
	Total *float64 `bson:"total,omitempty"           json:"total,omitempty"`

	// Recommendations
	RelatedProducts     []primitive.ObjectID `bson:"related_products,omitempty"     json:"related_products,omitempty"`
	RecommendedProducts []primitive.ObjectID `bson:"recommended_products,omitempty" json:"recommended_products,omitempty"`

	// Audit
	CreatedBy primitive.ObjectID `bson:"createdBy,omitempty"        json:"createdBy,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"        json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"        json:"updatedAt,omitempty"`
}
