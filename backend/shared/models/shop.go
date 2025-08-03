package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Shop Basic Info
type Shop struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Slug        string             `bson:"slug" json:"slug"` // unique URL identifier
	Image       string             `bson:"image,omitempty"  json:"image,omitempty"`
	Banner      string             `bson:"banner,omitempty"  json:"banner,omitempty"` // shop banner/cover image
	OwnerID     primitive.ObjectID `bson:"ownerId" json:"ownerId"`
	Description string             `bson:"description" json:"description"`

	// Business Information
	Category string `bson:"category,omitempty" json:"category,omitempty"` // Shop category (fashion, electronics, etc.)
	Email    string `bson:"email,omitempty" json:"email,omitempty"`       // Contact email
	Phone    string `bson:"phone,omitempty" json:"phone,omitempty"`       // Contact phone
	Address  string `bson:"address,omitempty" json:"address,omitempty"`   // Business address
	Currency string `bson:"currency,omitempty" json:"currency,omitempty"` // Default currency (USD, EUR, etc.)

	// Business Status
	Status     string `bson:"status,omitempty" json:"status,omitempty"`         // Shop status (active, inactive, suspended)
	IsVerified bool   `bson:"isVerified,omitempty" json:"isVerified,omitempty"` // Shop verification status

	// Analytics & Metrics
	CustomerCount int     `bson:"customerCount,omitempty" json:"customerCount,omitempty"` // Number of customers
	ProductCount  int     `bson:"productCount,omitempty" json:"productCount,omitempty"`   // Number of products
	Revenue       float64 `bson:"revenue,omitempty" json:"revenue,omitempty"`             // Total revenue
	Rating        float64 `bson:"rating,omitempty" json:"rating,omitempty"`               // Average rating
	ReviewCount   int     `bson:"reviewCount,omitempty" json:"reviewCount,omitempty"`     // Number of reviews

	// Timestamps
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
