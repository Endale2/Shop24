// shared/models/order.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderItem represents a line in an order.
type OrderItem struct {
	ProductID  primitive.ObjectID `bson:"product_id"   json:"product_id"`
	VariantID  primitive.ObjectID `bson:"variant_id"   json:"variant_id,omitempty"` // omitempty will hide zero ObjectID
	Name       string             `bson:"name"         json:"name"`
	Quantity   int                `bson:"quantity"     json:"quantity"`
	UnitPrice  float64            `bson:"unit_price"   json:"unit_price"`
	TotalPrice float64            `bson:"total_price"  json:"total_price"`
	Image      string             `bson:"image"        json:"image"` // Product or variant image
}

// Order represents a shop order.
type Order struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"     json:"id"`
	ShopID        primitive.ObjectID `bson:"shop_id"           json:"shop_id"`
	CustomerID    primitive.ObjectID `bson:"customer_id"       json:"customer_id"`
	Items         []OrderItem        `bson:"items"             json:"items"`
	Subtotal      float64            `bson:"subtotal"          json:"subtotal"`
	DiscountTotal float64            `bson:"discount_total"    json:"discount_total"`
	Total         float64            `bson:"total"             json:"total"`
	Status        string             `bson:"status"            json:"status"` // e.g. "pending","paid","shipped"
	CreatedAt     time.Time          `bson:"created_at"        json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"        json:"updated_at"`
}
