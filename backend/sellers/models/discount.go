package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

// DiscountType enumerates the kinds of discounts you support.
type DiscountType string

const (
    DiscountTypeFixed      DiscountType = "fixed"      // fixed‐amount off
    DiscountTypePercentage DiscountType = "percentage" // percent‐off
)

// Discount represents a promotion you can apply to products or shops.
type Discount struct {
    ID                 primitive.ObjectID   `bson:"_id,omitempty"        json:"id"`
    Name               string               `bson:"name"                 json:"name"` 
	ShopID             primitive.ObjectID `bson:"shop_id,omitempty"       json:"shop_id"`
    SellerID             primitive.ObjectID `bson:"user_id,omitempty"       json:"user_id"`                         // e.g. "Summer Sale 2025"
    Description        string               `bson:"description,omitempty" json:"description,omitempty"`
    Type               DiscountType         `bson:"type"                 json:"type"`                          // fixed vs percentage
    Value              float64              `bson:"value"                json:"value"`                         // $X off or X% off
    CouponCode         string               `bson:"coupon_code,omitempty" json:"coupon_code,omitempty"`        // if empty → auto apply
    AppliesToProducts  []primitive.ObjectID `bson:"applies_to_products,omitempty" json:"applies_to_products,omitempty"` // limit to these product IDs
    AppliesToShops     []primitive.ObjectID `bson:"applies_to_shops,omitempty"   json:"applies_to_shops,omitempty"`     // or these shop IDs
    AllowedCustomerIDs []primitive.ObjectID `bson:"allowed_customers,omitempty"  json:"allowed_customers,omitempty"`  // or empty→everyone
    UsageLimit         *int                 `bson:"usage_limit,omitempty"        json:"usage_limit,omitempty"`        // total uses across customers
    PerCustomerLimit   *int                 `bson:"per_customer_limit,omitempty" json:"per_customer_limit,omitempty"` // uses per customer
    StartAt            time.Time            `bson:"start_at"                json:"start_at"`                 // when it goes live
    EndAt              time.Time            `bson:"end_at"                  json:"end_at"`                   // when it expires
    Active             bool                 `bson:"active"                  json:"active"`                   // quick on/off toggle
    CreatedAt          time.Time            `bson:"created_at,omitempty"    json:"created_at,omitempty"`
    UpdatedAt          time.Time            `bson:"updated_at,omitempty"    json:"updated_at,omitempty"`
}
