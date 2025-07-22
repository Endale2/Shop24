package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Collection represents a group of products that are organized together.(for seller)
// This model can be used to represent both manual collections (where products are explicitly added)
// and smart collections (where products are automatically included based on defined filters).
type Collection struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID      primitive.ObjectID `bson:"shop_id,omitempty"         json:"shop_id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Handle      string             `bson:"handle" json:"handle"`
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
	// ProductIDs  []primitive.ObjectID `bson:"product_ids,omitempty" json:"product_ids,omitempty"` // Removed: products reference collection by collection_id
	Filters   map[string]interface{} `bson:"filters,omitempty" json:"filters,omitempty"`
	CreatedAt time.Time              `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time              `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
