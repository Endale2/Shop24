package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// Variant represents a specific version of a product.
type Variant struct {
    VariantID primitive.ObjectID `bson:"_id,omitempty"       json:"id"`
    Options   map[string]string  `bson:"options"             json:"options"`
    Price     float64            `bson:"price"               json:"price"`
    Stock     int                `bson:"stock"               json:"stock"`
    Image     string             `bson:"image,omitempty"     json:"image,omitempty"`

    CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
    UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

// Product represents the main product entity, with SEO and recommendations.
type Product struct {
    ID                  primitive.ObjectID   `bson:"_id,omitempty"           json:"id"`
    ShopID              primitive.ObjectID   `bson:"shop_id,omitempty"       json:"shop_id"`
    UserID              primitive.ObjectID   `bson:"user_id,omitempty"       json:"user_id"`

    // Core product fields
    Name                string               `bson:"name"                    json:"name"`
    Slug                string               `bson:"slug"                    json:"slug"`
    Description         string               `bson:"description"             json:"description"`

    // SEO fields
    MetaTitle           string               `bson:"meta_title,omitempty"    json:"meta_title,omitempty"`
    MetaDescription     string               `bson:"meta_description,omitempty" json:"meta_description,omitempty"`
    CanonicalURL        string               `bson:"canonical_url,omitempty"  json:"canonical_url,omitempty"`
    Tags                []string             `bson:"tags,omitempty"          json:"tags,omitempty"`

    // Media
    MainImage           string               `bson:"main_image,omitempty"    json:"main_image,omitempty"`
    Images              []string             `bson:"images,omitempty"        json:"images,omitempty"`

    // Pricing & inventory
    Category            string               `bson:"category"                json:"category"`
    Price               float64              `bson:"price"                   json:"price"`
    DisplayPrice        float64              `bson:"display_price"           json:"display_price"`
    Stock               int                  `bson:"stock"                   json:"stock"`

    // Variants & ratings
    Variants            []Variant            `bson:"variants,omitempty"      json:"variants,omitempty"`
    AverageRating       float64              `bson:"average_rating,omitempty" json:"average_rating,omitempty"`
    ReviewCount         int                  `bson:"review_count,omitempty"  json:"review_count,omitempty"`

    // Recommendations
    RelatedProducts     []primitive.ObjectID `bson:"related_products,omitempty"     json:"related_products,omitempty"`
    RecommendedProducts []primitive.ObjectID `bson:"recommended_products,omitempty" json:"recommended_products,omitempty"`

    // Audit fields
    CreatedBy           primitive.ObjectID   `bson:"createdBy,omitempty"      json:"createdBy,omitempty"`
    CreatedAt           time.Time            `bson:"createdAt,omitempty"      json:"createdAt,omitempty"`
    UpdatedAt           time.Time            `bson:"updatedAt,omitempty"      json:"updatedAt,omitempty"`
}
