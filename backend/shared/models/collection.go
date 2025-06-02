package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Collection represents a group of products that are organized together.(for seller)
// This model can be used to represent both manual collections (where products are explicitly added)
// and smart collections (where products are automatically included based on defined filters).
type Collection struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title       string               `bson:"title" json:"title"`                 // The display title of the collection
	Description string               `bson:"description" json:"description"`     // A description of what the collection is about
	Handle      string               `bson:"handle" json:"handle"`               // URL-friendly unique identifier (e.g., "summer-sale")
	Image       string               `bson:"image,omitempty" json:"image,omitempty"` // An optional image representing the collection
	ProductIDs  []primitive.ObjectID `bson:"product_ids,omitempty" json:"product_ids,omitempty"` // List of product IDs included in the collection
	Filters     map[string]interface{} `bson:"filters,omitempty" json:"filters,omitempty"` // Optional filters for smart collections (e.g., {"tag": "summer"})
	CreatedAt   time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time            `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
