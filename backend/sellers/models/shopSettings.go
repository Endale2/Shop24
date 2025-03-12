package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Shop Settings (Shipping, Taxes, etc.)
type ShopSettings struct {
	ShippingPolicy      string  `bson:"shippingPolicy"`  // E.g., "worldwide", "domestic only"
	TaxRate             float64 `bson:"taxRate"`         // Tax rate percentage for the shop (e.g., 0.10 for 10%)
	RefundPolicy        string  `bson:"refundPolicy"`    // Refund policy (e.g., "14-day return")
	SupportEmail        string  `bson:"supportEmail"`    // Contact email for customer support
	Language            string  `bson:"language"`        // Language of the shop (e.g., "en", "es")
	Currency            string  `bson:"currency"`        // Currency of the shop (e.g., "USD", "EUR")
	PaymentMethods      []string `bson:"paymentMethods"` // List of supported payment methods (e.g., ["PayPal", "Stripe"])
	ShippingMethods     []string `bson:"shippingMethods"` // List of available shipping methods (e.g., ["Free Shipping", "Express"])
}