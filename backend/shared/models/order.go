package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order represents a customer's order
type Order struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`   // ID of the customer
	ShopID         primitive.ObjectID `bson:"shop_id,omitempty" json:"shop_id"`   // ID of the shop where the order was placed
	Items          []OrderItem        `bson:"items" json:"items"`                 // List of items in the order
	TotalAmount    float64            `bson:"total_amount" json:"total_amount"`   // Total order amount
	Currency       string             `bson:"currency" json:"currency"`           // e.g., "USD", "EUR"
	PaymentStatus  string             `bson:"payment_status" json:"payment_status"` // e.g., "Paid", "Pending"
	ShippingStatus string             `bson:"shipping_status" json:"shipping_status"` // e.g., "Pending", "Shipped", "Delivered"
	ShippingInfo   ShippingDetails    `bson:"shipping_info" json:"shipping_info"`   // Shipping details for the order
	CreatedAt      time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt      time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// OrderItem represents a single product item in an order
type OrderItem struct {
	ProductID      primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"` // ID of the ordered product
	Name           string             `bson:"name" json:"name"`                                 // Name of the product
	Quantity       int                `bson:"quantity" json:"quantity"`                         // Number of units ordered
	Price          float64            `bson:"price" json:"price"`                               // Unit price at the time of order
	TotalPrice     float64            `bson:"total_price" json:"total_price"`                   // Price * Quantity
	VariantOptions map[string]string  `bson:"variant_options" json:"variant_options"`           // Flexible variant options snapshot (e.g., {"color": "Red", "size": "M"} or {"capacity": "128GB"})
	VariantImage   string             `bson:"variant_image,omitempty" json:"variant_image,omitempty"` // Optional: snapshot image of the variant if available
}

// ShippingDetails holds the shipping address and tracking information
type ShippingDetails struct {
	Address           string    `bson:"address" json:"address"`
	City              string    `bson:"city" json:"city"`
	PostalCode        string    `bson:"postal_code" json:"postal_code"`
	Country           string    `bson:"country" json:"country"`
	TrackingID        string    `bson:"tracking_id,omitempty" json:"tracking_id,omitempty"`
	EstimatedDelivery time.Time `bson:"estimated_delivery,omitempty" json:"estimated_delivery,omitempty"`
}
