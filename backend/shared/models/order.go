package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order represents a customer's order
type Order struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID             primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`                     // ID of the customer
	ShopID             primitive.ObjectID `bson:"shop_id,omitempty" json:"shop_id"`                     // ID of the shop where the order was placed (important for multi-vendor/marketplace)
	OrderNumber        string             `bson:"order_number" json:"order_number"`                     // Human-readable, unique order number (e.g., #XYZ12345)
	Items              []OrderItem        `bson:"items" json:"items"`                                   // List of items in the order
	TotalAmount        float64            `bson:"total_amount" json:"total_amount"`                     // Total order amount (sum of item total prices + shipping + taxes - discounts)
	SubtotalAmount     float64            `bson:"subtotal_amount" json:"subtotal_amount"`               // Sum of item total prices before shipping, taxes, and discounts
	DiscountAmount     float64            `bson:"discount_amount,omitempty" json:"discount_amount,omitempty"` // Total discount applied to the order
	ShippingAmount     float64            `bson:"shipping_amount,omitempty" json:"shipping_amount,omitempty"` // Total shipping cost     // Total tax amount
	Currency           string             `bson:"currency" json:"currency"`                             // e.g., "USD", "EUR"
	PaymentStatus      PaymentStatus      `bson:"payment_status" json:"payment_status"`                 // Enum: "Pending", "Paid", "Failed", "Refunded", "PartiallyRefunded"
	PaymentDetails     PaymentDetails     `bson:"payment_details,omitempty" json:"payment_details,omitempty"` // Details about the payment transaction
	ShippingStatus     ShippingStatus     `bson:"shipping_status" json:"shipping_status"`               // Enum: "Pending", "Processing", "Shipped", "Delivered", "Canceled", "Returned"
	ShippingInfo       ShippingDetails    `bson:"shipping_info" json:"shipping_info"`                   // Shipping details for the order
	BillingAddress     Address            `bson:"billing_address" json:"billing_address"`               // Billing address for the order
	CustomerNotes      string             `bson:"customer_notes,omitempty" json:"customer_notes,omitempty"` // Any notes from the customer
	InternalNotes      string             `bson:"internal_notes,omitempty" json:"internal_notes,omitempty"` // Internal notes for order processing
	OrderSource        string             `bson:"order_source,omitempty" json:"order_source,omitempty"` // e.g., "Website", "MobileApp", "AdminPanel", "API"
	Channel            string             `bson:"channel,omitempty" json:"channel,omitempty"`           // e.g., "Web", "iOS", "Android", "MarketplaceA"
	AppliedCoupons     []AppliedCoupon    `bson:"applied_coupons,omitempty" json:"applied_coupons,omitempty"` // Details of applied coupons
	Events             []OrderEvent       `bson:"events,omitempty" json:"events,omitempty"`             // Log of significant order events
	CreatedAt          time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt          time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeliveredAt        *time.Time         `bson:"delivered_at,omitempty" json:"delivered_at,omitempty"` // Timestamp for delivery
	CanceledAt         *time.Time         `bson:"canceled_at,omitempty" json:"canceled_at,omitempty"`   // Timestamp for cancellation
}

// PaymentStatus defines the possible states of an order's payment.
type PaymentStatus string

const (
	PaymentStatusPending          PaymentStatus = "Pending"
	PaymentStatusPaid             PaymentStatus = "Paid"
	PaymentStatusFailed           PaymentStatus = "Failed"
	PaymentStatusRefunded         PaymentStatus = "Refunded"
	PaymentStatusPartiallyRefunded PaymentStatus = "PartiallyRefunded"
)

// ShippingStatus defines the possible states of an order's shipping.
type ShippingStatus string

const (
	ShippingStatusPending    ShippingStatus = "Pending"
	ShippingStatusProcessing ShippingStatus = "Processing"
	ShippingStatusShipped    ShippingStatus = "Shipped"
	ShippingStatusDelivered  ShippingStatus = "Delivered"
	ShippingStatusCanceled   ShippingStatus = "Canceled"
	ShippingStatusReturned   ShippingStatus = "Returned"
)

// OrderItem represents a single product item or package component in an order
type OrderItem struct {
	ProductID       primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"`       // ID of the ordered product or package
	SKU             string             `bson:"sku,omitempty" json:"sku,omitempty"`                   // Product SKU (Stock Keeping Unit)
	Name            string             `bson:"name" json:"name"`                                     // Name of the product or package
	Quantity        int                `bson:"quantity" json:"quantity"`                             // Number of units ordered
	UnitPrice       float64            `bson:"unit_price" json:"unit_price"`                         // Unit price at the time of order (before item-specific discounts)
	Discount        float64            `bson:"discount,omitempty" json:"discount,omitempty"`         // Item-specific discount amount
	TotalPrice      float64            `bson:"total_price" json:"total_price"`                       // (UnitPrice * Quantity) - Discount
	VariantOptions  map[string]string  `bson:"variant_options,omitempty" json:"variant_options,omitempty"` // Flexible variant options snapshot (e.g., {"color": "Red", "size": "M"} or {"capacity": "128GB"})
	VariantImage    string             `bson:"variant_image,omitempty" json:"variant_image,omitempty"` // Optional: snapshot image of the variant if available
	ItemType        string             `bson:"item_type,omitempty" json:"item_type,omitempty"`         // e.g., "physical", "digital", "subscription", "bundle"
	BundleComponents []BundleComponent `bson:"bundle_components,omitempty" json:"bundle_components,omitempty"` // For "bundle" ItemType, details of components
}

// BundleComponent represents a component within a package/bundle
type BundleComponent struct {
	ProductID primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"` // ID of the component product
	SKU       string             `bson:"sku,omitempty" json:"sku,omitempty"`           // Component SKU
	Name      string             `bson:"name" json:"name"`                             // Name of the component product
	Quantity  int                `bson:"quantity" json:"quantity"`                     // Quantity of this component within the bundle
	UnitPrice float64            `bson:"unit_price" json:"unit_price"`                 // Unit price of this component (for reference, might not sum up to bundle price)
	// Add other relevant fields if needed, e.g., variant options for the component
}

// ShippingDetails holds the shipping address and tracking information
type ShippingDetails struct {
	AddressLine1      string     `bson:"address_line1" json:"address_line1"`
	AddressLine2      string     `bson:"address_line2,omitempty" json:"address_line2,omitempty"`
	City              string     `bson:"city" json:"city"`
	StateProvince     string     `bson:"state_province,omitempty" json:"state_province,omitempty"`
	PostalCode        string     `bson:"postal_code" json:"postal_code"`
	Country           string     `bson:"country" json:"country"`
	RecipientName     string     `bson:"recipient_name" json:"recipient_name"`
	PhoneNumber       string     `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	ShippingMethod    string     `bson:"shipping_method,omitempty" json:"shipping_method,omitempty"`     // e.g., "Standard", "Express", "Pickup"
	ShippingProvider  string     `bson:"shipping_provider,omitempty" json:"shipping_provider,omitempty"` // e.g., "FedEx", "DHL", "Local Courier"
	TrackingID        string     `bson:"tracking_id,omitempty" json:"tracking_id,omitempty"`
	TrackingURL       string     `bson:"tracking_url,omitempty" json:"tracking_url,omitempty"`           // Direct URL to tracking info
	EstimatedDelivery *time.Time `bson:"estimated_delivery,omitempty" json:"estimated_delivery,omitempty"`
	ShippedAt         *time.Time `bson:"shipped_at,omitempty" json:"shipped_at,omitempty"`
}

// Address is a reusable struct for billing and potentially other addresses
type Address struct {
	AddressLine1  string `bson:"address_line1" json:"address_line1"`
	AddressLine2  string `bson:"address_line2,omitempty" json:"address_line2,omitempty"`
	City          string `bson:"city" json:"city"`
	StateProvince string `bson:"state_province,omitempty" json:"state_province,omitempty"`
	PostalCode    string `bson:"postal_code" json:"postal_code"`
	Country       string `bson:"country" json:"country"`
	FullName      string `bson:"full_name" json:"full_name"`
	PhoneNumber   string `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}

// PaymentDetails holds information about the payment transaction
type PaymentDetails struct {
	TransactionID   string    `bson:"transaction_id,omitempty" json:"transaction_id,omitempty"`     // ID from payment gateway
	PaymentMethod   string    `bson:"payment_method,omitempty" json:"payment_method,omitempty"`     // e.g., "Credit Card", "PayPal", "Bank Transfer"
	LastFour        string    `bson:"last_four,omitempty" json:"last_four,omitempty"`               // Last four digits of card if applicable
	Brand           string    `bson:"brand,omitempty" json:"brand,omitempty"`                       // Card brand (Visa, Mastercard)
	AmountPaid      float64   `bson:"amount_paid,omitempty" json:"amount_paid,omitempty"`           // Actual amount paid (can differ for partial refunds)
	PaymentCurrency string    `bson:"payment_currency,omitempty" json:"payment_currency,omitempty"` // Currency of the payment transaction
	PaidAt          *time.Time `bson:"paid_at,omitempty" json:"paid_at,omitempty"`                  // Timestamp when payment was confirmed
	RawResponse     string    `bson:"raw_response,omitempty" json:"raw_response,omitempty"`         // Optional: raw JSON response from payment gateway for debugging/auditing
}

// AppliedCoupon represents a coupon applied to the order
type AppliedCoupon struct {
	CouponCode string  `bson:"coupon_code" json:"coupon_code"`
	DiscountAmount float64 `bson:"discount_amount" json:"discount_amount"` // Amount discounted by this specific coupon
	CouponType string  `bson:"coupon_type,omitempty" json:"coupon_type,omitempty"` // e.g., "percentage", "fixed_amount", "free_shipping"
}

// OrderEvent logs significant changes or actions related to the order
type OrderEvent struct {
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	EventType string    `bson:"event_type" json:"event_type"` // e.g., "OrderPlaced", "PaymentReceived", "Shipped", "StatusUpdated", "ItemAdded", "RefundIssued"
	Description string    `bson:"description" json:"description"` // Human-readable description
	Actor       string    `bson:"actor,omitempty" json:"actor,omitempty"` // User or System that triggered the event
	Metadata    map[string]interface{} `bson:"metadata,omitempty" json:"metadata,omitempty"` // Additional event-specific data
}