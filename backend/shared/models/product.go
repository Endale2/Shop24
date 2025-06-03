package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// Variant represents a specific version of a product.
type Variant struct {
    VariantID primitive.ObjectID    `bson:"_id,omitempty"    json:"id"`
    Options   map[string]string     `bson:"options"          json:"options"`
    Price     float64               `bson:"price"            json:"price"`
    Stock     int                   `bson:"stock"            json:"stock"`
    Image     string                `bson:"image,omitempty"  json:"image,omitempty"`

    
    // Timestamps for each variant:
    CreatedAt time.Time             `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
    UpdatedAt time.Time             `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

// Product represents the main product entity.
type Product struct {
    ID            primitive.ObjectID   `bson:"_id,omitempty"       json:"id"`
    ShopID        primitive.ObjectID   `bson:"shop_id,omitempty"   json:"shop_id"`
    UserID        primitive.ObjectID   `bson:"user_id,omitempty"   json:"user_id"`
    DiscountID    []primitive.ObjectID `bson:"discount_id,omitempty" json:"discount_id"`
    Name          string               `bson:"name"               json:"name"`
    Slug          string               `bson:"slug"               json:"slug"`
    Description   string               `bson:"description"        json:"description"`
    MainImage     string               `bson:"main_image,omitempty" json:"main_image,omitempty"`
    Images        []string             `bson:"images"             json:"images"`
    Category      string               `bson:"category"           json:"category"`
    Price         float64              `bson:"price"              json:"price"`
    DisplayPrice  float64              `bson:"display_price"      json:"display_price"`
    CreatedBy     primitive.ObjectID   `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
    Variants      []Variant            `bson:"variants"           json:"variants"`
    CreatedAt     time.Time            `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
    UpdatedAt     time.Time            `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
