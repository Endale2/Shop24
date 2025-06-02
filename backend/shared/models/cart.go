package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Cart struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	CustomerID     *primitive.ObjectID  `bson:"customer_id,omitempty" json:"customer_id,omitempty"`
	SessionID      *string              `bson:"session_id,omitempty" json:"session_id,omitempty"` // for guest carts
	ShopID         primitive.ObjectID   `bson:"shop_id" json:"shop_id"`
	Items          []CartItem           `bson:"items" json:"items"`

	Subtotal       float64              `bson:"subtotal" json:"subtotal"` // sum of item LineTotals (before discounts)
	TotalDiscounts float64              `bson:"total_discounts" json:"total_discounts"` // sum of all discounts (order + item)
	ShippingCost   float64              `bson:"shipping_cost,omitempty" json:"shipping_cost,omitempty"` // calculated separately
	TaxAmount      float64              `bson:"tax_amount,omitempty" json:"tax_amount,omitempty"`       // calculated separately
	GrandTotal     float64              `bson:"grand_total" json:"grand_total"` // Subtotal - Discounts + Shipping + Tax

	AppliedDiscountIDs []primitive.ObjectID `bson:"applied_discount_ids,omitempty" json:"applied_discount_ids,omitempty"` // for order-wide or shipping discounts

	Currency       string               `bson:"currency" json:"currency"` // e.g., "USD"
	LastUpdated    time.Time            `bson:"last_updated" json:"last_updated"`
	CreatedAt      time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
