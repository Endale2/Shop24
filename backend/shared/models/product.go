package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)



// Variant represents a specific version of a product.
// The Options map can hold any variant attribute such as "color", "size", "capacity", etc.

type Variant struct {
	Options map[string]string `bson:"options" json:"options"` // e.g., {"color": "Red", "size": "M"} or {"capacity": "128GB"}
	Price   float64           `bson:"price" json:"price"`     // Price for this specific variant
	Stock   int               `bson:"stock" json:"stock"`     // Available stock
	Image   string            `bson:"image,omitempty" json:"image,omitempty"` // Optional image for the variant
}



// Product represents the main product entity.
// It can have multiple variants with different options, prices, and stock levels.
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID      primitive.ObjectID `bson:"shop_id,omitempty" json:"shop_id"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Images      []string           `bson:"images" json:"images"` // General images for the product
	Category    string             `bson:"category" json:"category"`
	Price       float64            `bson:"price" json:"price"` // Default price for products without variants
	CreatedBy   primitive.ObjectID `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	Variants    []Variant          `bson:"variants" json:"variants"` // List of variants with flexible options
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
