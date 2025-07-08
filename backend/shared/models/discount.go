package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DiscountCategory is *where* it applies:
type DiscountCategory string

const (
	DiscountCategoryProduct  DiscountCategory = "product"     // specific products or variants
	DiscountCategoryOrder    DiscountCategory = "order"       // entire cart
	DiscountCategoryShipping DiscountCategory = "shipping"    // free shipping, etc.
	DiscountCategoryBuyXGetY DiscountCategory = "buy_x_get_y" // buy X get Y
)

// DiscountType is *how much*:
type DiscountType string

const (
	DiscountTypeFixed      DiscountType = "fixed"      // $X off
	DiscountTypePercentage DiscountType = "percentage" // X% off
)

// DiscountEligibilityType defines how customer eligibility is determined
type DiscountEligibilityType string

const (
	DiscountEligibilityAll      DiscountEligibilityType = "all"      // Everyone can use
	DiscountEligibilitySpecific DiscountEligibilityType = "specific" // Only specific customers
	DiscountEligibilitySegment  DiscountEligibilityType = "segment"  // Only customers in specific segments
)

// DiscountUsage tracks usage of a discount
type DiscountUsage struct {
	CustomerID primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	UsageCount int                `bson:"usage_count" json:"usage_count"`
	LastUsedAt time.Time          `bson:"last_used_at" json:"last_used_at"`
	TotalSpent float64            `bson:"total_spent" json:"total_spent"`
}

type Discount struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`

	// *where* it applies
	Category DiscountCategory `bson:"category" json:"category"`
	// *how much* off
	Type  DiscountType `bson:"type,omitempty" json:"type,omitempty"`
	Value float64      `bson:"value,omitempty" json:"value,omitempty"`

	ShopID   primitive.ObjectID `bson:"shop_id" json:"shop_id"`
	SellerID primitive.ObjectID `bson:"seller_id" json:"seller_id"`

	// manual targets
	AppliesToCollections []primitive.ObjectID `bson:"applies_to_collections,omitempty" json:"applies_to_collections,omitempty"`
	AppliesToProducts    []primitive.ObjectID `bson:"applies_to_products,omitempty"    json:"applies_to_products,omitempty"`
	AppliesToVariants    []primitive.ObjectID `bson:"applies_to_variants,omitempty"    json:"applies_to_variants,omitempty"`

	// Enhanced customer eligibility
	EligibilityType    DiscountEligibilityType `bson:"eligibility_type" json:"eligibility_type"`
	AllowedCustomerIDs []primitive.ObjectID    `bson:"allowed_customers,omitempty" json:"allowed_customers,omitempty"`
	AllowedSegmentIDs  []primitive.ObjectID    `bson:"allowed_segments,omitempty" json:"allowed_segments,omitempty"`

	// Usage limits
	UsageLimit       *int            `bson:"usage_limit,omitempty" json:"usage_limit,omitempty"`
	PerCustomerLimit *int            `bson:"per_customer_limit,omitempty" json:"per_customer_limit,omitempty"`
	CurrentUsage     int             `bson:"current_usage" json:"current_usage"`
	UsageTracking    []DiscountUsage `bson:"usage_tracking,omitempty" json:"usage_tracking,omitempty"`

	// Buy X Get Y
	BuyProductIDs []primitive.ObjectID `bson:"buy_product_ids,omitempty"   json:"buy_product_ids,omitempty"`
	BuyQuantity   *int                 `bson:"buy_quantity,omitempty"      json:"buy_quantity,omitempty"`
	GetProductIDs []primitive.ObjectID `bson:"get_product_ids,omitempty"   json:"get_product_ids,omitempty"`
	GetQuantity   *int                 `bson:"get_quantity,omitempty"      json:"get_quantity,omitempty"`

	// Order‐level
	MinimumOrderSubtotal *float64 `bson:"minimum_order_subtotal,omitempty" json:"minimum_order_subtotal,omitempty"`
	// Shipping‐level
	FreeShipping                bool     `bson:"free_shipping,omitempty" json:"free_shipping,omitempty"`
	MinimumOrderForFreeShipping *float64 `bson:"minimum_free_shipping,omitempty" json:"minimum_free_shipping,omitempty"`

	StartAt time.Time `bson:"start_at"      json:"start_at"`
	EndAt   time.Time `bson:"end_at"        json:"end_at"`
	Active  bool      `bson:"active"        json:"active"`

	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// IsEligible checks if a customer is eligible for this discount
func (d *Discount) IsEligible(customerID primitive.ObjectID, customerSegmentIDs []primitive.ObjectID) bool {
	switch d.EligibilityType {
	case DiscountEligibilityAll:
		return true
	case DiscountEligibilitySpecific:
		for _, allowedID := range d.AllowedCustomerIDs {
			if allowedID == customerID {
				return true
			}
		}
		return false
	case DiscountEligibilitySegment:
		for _, segmentID := range d.AllowedSegmentIDs {
			for _, customerSegmentID := range customerSegmentIDs {
				if segmentID == customerSegmentID {
					return true
				}
			}
		}
		return false
	default:
		return true
	}
}

// CanUse checks if a customer can use this discount (considering usage limits)
func (d *Discount) CanUse(customerID primitive.ObjectID) bool {
	// Check overall usage limit
	if d.UsageLimit != nil && d.CurrentUsage >= *d.UsageLimit {
		return false
	}

	// Check per-customer limit
	if d.PerCustomerLimit != nil {
		for _, usage := range d.UsageTracking {
			if usage.CustomerID == customerID && usage.UsageCount >= *d.PerCustomerLimit {
				return false
			}
		}
	}

	return true
}

// RecordUsage records a usage of this discount by a customer
func (d *Discount) RecordUsage(customerID primitive.ObjectID, amount float64) {
	d.CurrentUsage++
	now := time.Now()

	// Find existing usage record for this customer
	for i := range d.UsageTracking {
		if d.UsageTracking[i].CustomerID == customerID {
			d.UsageTracking[i].UsageCount++
			d.UsageTracking[i].LastUsedAt = now
			d.UsageTracking[i].TotalSpent += amount
			return
		}
	}

	// Create new usage record
	d.UsageTracking = append(d.UsageTracking, DiscountUsage{
		CustomerID: customerID,
		UsageCount: 1,
		LastUsedAt: now,
		TotalSpent: amount,
	})
}

// IsActive checks if the discount is currently active
func (d *Discount) IsActive() bool {
	now := time.Now()

	// Check if discount is marked as active
	if !d.Active {
		return false
	}

	// Check start time (if StartAt is zero, consider it as "started")
	if !d.StartAt.IsZero() && now.Before(d.StartAt) {
		return false
	}

	// Check end time (if EndAt is zero, consider it as "no end")
	if !d.EndAt.IsZero() && now.After(d.EndAt) {
		return false
	}

	return true
}

// AppliesToProduct checks if this discount applies to a specific product
func (d *Discount) AppliesToProduct(productID primitive.ObjectID, collectionIDs []primitive.ObjectID) bool {
	// Check direct product application
	for _, pid := range d.AppliesToProducts {
		if pid == productID {
			return true
		}
	}

	// Check collection application
	for _, discountCollectionID := range d.AppliesToCollections {
		for _, productCollectionID := range collectionIDs {
			if discountCollectionID == productCollectionID {
				return true
			}
		}
	}

	return false
}

// AppliesToVariant checks if this discount applies to a specific variant
func (d *Discount) AppliesToVariant(variantID primitive.ObjectID) bool {
	for _, vid := range d.AppliesToVariants {
		if vid == variantID {
			return true
		}
	}
	return false
}

// CalculateDiscount calculates the discount amount for a given price
func (d *Discount) CalculateDiscount(price float64) float64 {
	switch d.Type {
	case DiscountTypeFixed:
		if d.Value > price {
			return price
		}
		return d.Value
	case DiscountTypePercentage:
		return price * (d.Value / 100)
	default:
		return 0
	}
}
