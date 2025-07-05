package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerSegment represents a customer segment/folder for organizing customers within a shop
type CustomerSegment struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	ShopID      primitive.ObjectID   `bson:"shop_id" json:"shop_id"`
	Name        string               `bson:"name" json:"name"`
	Description string               `bson:"description,omitempty" json:"description,omitempty"`
	Color       string               `bson:"color,omitempty" json:"color,omitempty"` // Hex color for UI
	CustomerIDs []primitive.ObjectID `bson:"customer_ids,omitempty" json:"customer_ids,omitempty"`
	CreatedAt   time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at" json:"updated_at"`
}
