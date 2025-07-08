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
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShopID      primitive.ObjectID `bson:"shop_id" json:"shop_id"`
	CustomerID  primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	OrderNumber string             `bson:"order_number" json:"order_number"`
	Status      string             `bson:"status" json:"status"`

	// Items
	Items []OrderItem `bson:"items" json:"items"`

	// Pricing
	Subtotal      float64 `bson:"subtotal" json:"subtotal"`
	DiscountTotal float64 `bson:"discount_total" json:"discount_total"`
	ShippingCost  float64 `bson:"shipping_cost" json:"shipping_cost"`
	TaxAmount     float64 `bson:"tax_amount" json:"tax_amount"`
	Total         float64 `bson:"total" json:"total"`

	// Applied discounts
	AppliedDiscountIDs []primitive.ObjectID `bson:"applied_discount_ids,omitempty" json:"applied_discount_ids,omitempty"`

	// Shipping
	ShippingAddress map[string]interface{} `bson:"shipping_address" json:"shipping_address"`
	BillingAddress  map[string]interface{} `bson:"billing_address" json:"billing_address"`

	// Payment
	PaymentMethod string `bson:"payment_method" json:"payment_method"`
	PaymentStatus string `bson:"payment_status" json:"payment_status"`

	// Timestamps
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// OrderWithCustomer represents an order with customer information populated
type OrderWithCustomer struct {
	Order    Order     `json:"order"`
	Customer *Customer `json:"customer"`
}

// Customer represents customer information (imported from customers package)
// This is a duplicate of the customer model for JSON serialization purposes
type Customer struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string             `bson:"firstName" json:"firstName"`
	LastName     string             `bson:"lastName" json:"lastName"`
	ProfileImage string             `bson:"profile_image,omitempty" json:"profile_image,omitempty"`
	Email        string             `bson:"email" json:"email"`
	Phone        string             `bson:"phone" json:"phone"`
	Address      string             `bson:"address" json:"address"`
	City         string             `bson:"city" json:"city"`
	State        string             `bson:"state" json:"state"`
	PostalCode   string             `bson:"postalCode" json:"postalCode"`
	Country      string             `bson:"country" json:"country"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
