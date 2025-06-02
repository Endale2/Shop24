package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

// DiscountCategory captures the broad category of this discount.
type DiscountCategory string

const (
    DiscountCategoryProduct     DiscountCategory = "product"      // applies to specific product(s) or variant(s)
    DiscountCategoryOrder       DiscountCategory = "order"        // applies to the entire cart/order
    DiscountCategoryShipping    DiscountCategory = "shipping"     // free shipping or shipping‐only promo
    DiscountCategoryBuyXGetY    DiscountCategory = "buy_x_get_y"   // buy X items, get Y items
)

// DiscountType enumerates how the discount value is interpreted.
type DiscountType string

const (
    DiscountTypeFixed      DiscountType = "fixed"      // $X off (product or order)
    DiscountTypePercentage DiscountType = "percentage" // X% off (product or order)
)

// Discount represents a promotion you can apply to products, orders, or shipping.
type Discount struct {
    ID                 primitive.ObjectID   `bson:"_id,omitempty"             json:"id"`
    Name               string               `bson:"name"                      json:"name"`
    Description        string               `bson:"description,omitempty"     json:"description,omitempty"`
    Category           DiscountCategory     `bson:"category"                  json:"category"`    // product, order, shipping, or buy_x_get_y
    Type               DiscountType         `bson:"type,omitempty"            json:"type,omitempty"`     // fixed or percentage
    Value              float64              `bson:"value,omitempty"           json:"value,omitempty"`    // $X off or X% off; use only if Category is product/order

    // ── Scope fields ──
    ShopID             primitive.ObjectID   `bson:"shop_id,omitempty"         json:"shop_id,omitempty"`   // if limiting to a specific shop
    SellerID           primitive.ObjectID   `bson:"seller_id,omitempty"       json:"seller_id,omitempty"`
    AppliesToProducts  []primitive.ObjectID `bson:"applies_to_products,omitempty" json:"applies_to_products,omitempty"` 
    AppliesToVariants  []primitive.ObjectID `bson:"applies_to_variants,omitempty" json:"applies_to_variants,omitempty"` 
    // If Category == "order", you might instead use MinimumOrderSubtotal (see below).

    // ── Customer & usage limits ──
    CouponCode         string               `bson:"coupon_code,omitempty"         json:"coupon_code,omitempty"`
    AllowedCustomerIDs []primitive.ObjectID `bson:"allowed_customers,omitempty"   json:"allowed_customers,omitempty"`
    UsageLimit         *int                 `bson:"usage_limit,omitempty"         json:"usage_limit,omitempty"`         // total redemptions
    PerCustomerLimit   *int                 `bson:"per_customer_limit,omitempty"  json:"per_customer_limit,omitempty"`  // uses per customer

    // ── Buy X Get Y fields ──
    BuyProductIDs      []primitive.ObjectID `bson:"buy_product_ids,omitempty"        json:"buy_product_ids,omitempty"`        // which product(s) to buy
    BuyQuantity        *int                 `bson:"buy_quantity,omitempty"           json:"buy_quantity,omitempty"`          // minimum total quantity
    GetProductIDs      []primitive.ObjectID `bson:"get_product_ids,omitempty"         json:"get_product_ids,omitempty"`        // which product(s) are free
    GetQuantity        *int                 `bson:"get_quantity,omitempty"            json:"get_quantity,omitempty"`         // how many free items
    // If GetProductIDs is empty, assume “the same product(s) you buy.”

    // ── Order-level discount fields ──
    MinimumOrderSubtotal *float64           `bson:"minimum_order_subtotal,omitempty" json:"minimum_order_subtotal,omitempty"` // e.g. $100
    // (Value + Type still apply, e.g. “10% off if order >= $100”)

    // ── Free-shipping fields ──
    FreeShipping       bool                `bson:"free_shipping,omitempty"          json:"free_shipping,omitempty"`
    MinimumOrderForFreeShipping *float64    `bson:"minimum_free_shipping,omitempty"  json:"minimum_free_shipping,omitempty"`  // e.g. $50

    // ── Timing & active flag ──
    StartAt            time.Time           `bson:"start_at"                         json:"start_at"`
    EndAt              time.Time           `bson:"end_at"                           json:"end_at"`
    Active             bool                `bson:"active"                           json:"active"` // manual on/off

    CreatedAt          time.Time           `bson:"created_at,omitempty"             json:"created_at,omitempty"`
    UpdatedAt          time.Time           `bson:"updated_at,omitempty"             json:"updated_at,omitempty"`
}
