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
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
