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
	Products       []primitive.ObjectID `bson:"products"` // Product IDs linked to this shop
	Orders         []primitive.ObjectID `bson:"orders"`   // Order IDs linked to this shop
	CreatedAt      time.Time            `bson:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
	
}


