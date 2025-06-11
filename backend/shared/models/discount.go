package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

// DiscountCategory is *where* it applies:
type DiscountCategory string

const (
    DiscountCategoryProduct   DiscountCategory = "product"    // specific products or variants
    DiscountCategoryOrder     DiscountCategory = "order"      // entire cart
    DiscountCategoryShipping  DiscountCategory = "shipping"   // free shipping, etc.
    DiscountCategoryBuyXGetY  DiscountCategory = "buy_x_get_y"// buy X get Y
)

// DiscountType is *how much*:
type DiscountType string

const (
    DiscountTypeFixed      DiscountType = "fixed"       // $X off
    DiscountTypePercentage DiscountType = "percentage"  // X% off
)

type Discount struct {
    ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
    Name        string               `bson:"name" json:"name"`
    Description string               `bson:"description,omitempty" json:"description,omitempty"`

    // *where* it applies
    Category            DiscountCategory     `bson:"category" json:"category"`
    // *how much* off
    Type                DiscountType         `bson:"type,omitempty" json:"type,omitempty"`
    Value               float64              `bson:"value,omitempty" json:"value,omitempty"`

    ShopID              primitive.ObjectID   `bson:"shop_id" json:"shop_id"`
    SellerID            primitive.ObjectID   `bson:"seller_id" json:"seller_id"`

    // manual targets
    AppliesToCollections []primitive.ObjectID `bson:"applies_to_collections,omitempty" json:"applies_to_collections,omitempty"`
    AppliesToProducts    []primitive.ObjectID `bson:"applies_to_products,omitempty"    json:"applies_to_products,omitempty"`
    AppliesToVariants    []primitive.ObjectID `bson:"applies_to_variants,omitempty"    json:"applies_to_variants,omitempty"`

    CouponCode           string               `bson:"coupon_code,omitempty" json:"coupon_code,omitempty"`
    AllowedCustomerIDs   []primitive.ObjectID `bson:"allowed_customers,omitempty" json:"allowed_customers,omitempty"`
    UsageLimit           *int                 `bson:"usage_limit,omitempty" json:"usage_limit,omitempty"`
    PerCustomerLimit     *int                 `bson:"per_customer_limit,omitempty" json:"per_customer_limit,omitempty"`

    // Buy X Get Y
    BuyProductIDs       []primitive.ObjectID `bson:"buy_product_ids,omitempty"   json:"buy_product_ids,omitempty"`
    BuyQuantity         *int                 `bson:"buy_quantity,omitempty"      json:"buy_quantity,omitempty"`
    GetProductIDs       []primitive.ObjectID `bson:"get_product_ids,omitempty"   json:"get_product_ids,omitempty"`
    GetQuantity         *int                 `bson:"get_quantity,omitempty"      json:"get_quantity,omitempty"`

    // Order‐level
    MinimumOrderSubtotal          *float64 `bson:"minimum_order_subtotal,omitempty" json:"minimum_order_subtotal,omitempty"`
    // Shipping‐level
    FreeShipping                  bool     `bson:"free_shipping,omitempty" json:"free_shipping,omitempty"`
    MinimumOrderForFreeShipping   *float64 `bson:"minimum_free_shipping,omitempty" json:"minimum_free_shipping,omitempty"`

    StartAt        time.Time            `bson:"start_at"      json:"start_at"`
    EndAt          time.Time            `bson:"end_at"        json:"end_at"`
    Active         bool                 `bson:"active"        json:"active"`

    CreatedAt      time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
    UpdatedAt      time.Time            `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
