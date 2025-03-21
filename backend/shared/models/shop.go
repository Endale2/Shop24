package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Shop Basic Info
type Shop struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Name           string               `bson:"name"`
	OwnerID        primitive.ObjectID   `bson:"ownerId"` // Reference to user ID
	Description    string               `bson:"description"`
	CreatedAt      time.Time            `bson:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
	
}


