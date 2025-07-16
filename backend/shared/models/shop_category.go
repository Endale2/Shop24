package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ShopCategory represents a category for shops (e.g., Electronics, Fashion, etc.)
type ShopCategory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Slug        string             `bson:"slug" json:"slug"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
