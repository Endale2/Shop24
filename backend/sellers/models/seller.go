package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Seller represents a vendor on the platform
type Seller struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID   `bson:"user_id,omitempty" json:"user_id"` // Reference to the user account for authentication
	ShopIDs      []primitive.ObjectID `bson:"shop_ids,omitempty" json:"shop_ids,omitempty"` // List of shops owned by the seller
	Name         string               `bson:"name" json:"name"`         // Seller's display name
	Email        string               `bson:"email" json:"email"`       // Contact email
	Phone        string               `bson:"phone,omitempty" json:"phone,omitempty"` // Optional phone number
	Address      string               `bson:"address,omitempty" json:"address,omitempty"` // Business or mailing address
	BusinessName string               `bson:"business_name,omitempty" json:"business_name,omitempty"` // Legal or trading name of the business
	Ratings      float64              `bson:"ratings,omitempty" json:"ratings,omitempty"` // Average rating from customer reviews
	TotalSales   float64              `bson:"total_sales,omitempty" json:"total_sales,omitempty"` // Total revenue or sales figure
	ReviewsCount int                  `bson:"reviews_count,omitempty" json:"reviews_count,omitempty"` // Total number of reviews received
	CreatedAt    time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    time.Time            `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
