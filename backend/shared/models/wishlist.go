package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wishlist struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	ShopID     primitive.ObjectID   `bson:"shop_id" json:"shop_id"`
	CustomerID primitive.ObjectID   `bson:"customer_id" json:"customer_id"`
	ProductIDs []primitive.ObjectID `bson:"product_ids" json:"product_ids"`
	CreatedAt  time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time            `bson:"updated_at" json:"updated_at"`
}
