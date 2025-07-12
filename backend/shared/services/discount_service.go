package services

import (
	"errors"
	"time"

	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrDiscountNotFound = errors.New("discount not found")
var ErrDiscountNotEligible = errors.New("customer not eligible for this discount")
var ErrDiscountUsageLimitExceeded = errors.New("discount usage limit exceeded")
var ErrDiscountNotActive = errors.New("discount is not active")

func CreateDiscountService(d *models.Discount) (*models.Discount, error) {
	if d.Name == "" {
		return nil, errors.New("discount name is required")
	}
	if d.ShopID.IsZero() {
		return nil, errors.New("shop ID is required")
	}
	if d.Category == "" {
		return nil, errors.New("discount category is required")
	}
	if !d.StartAt.IsZero() && !d.EndAt.IsZero() && d.EndAt.Before(d.StartAt) {
		return nil, errors.New("endAt must be after startAt")
	}

	// Validate discount value
	if err := ValidateDiscountValue(d.Type, d.Value); err != nil {
		return nil, err
	}

	// Set default eligibility type if not specified
	if d.EligibilityType == "" {
		d.EligibilityType = models.DiscountEligibilityAll
	}

	// Initialize usage tracking
	d.CurrentUsage = 0
	d.UsageTracking = []models.DiscountUsage{}

	res, err := repositories.CreateDiscount(d)
	if err != nil {
		return nil, err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		d.ID = oid
	}
	return d, nil
}

func GetDiscountByIDService(idStr string) (*models.Discount, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid discount ID")
	}
	d, err := repositories.GetDiscountByID(id)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, ErrDiscountNotFound
	}
	return d, nil
}

func ListDiscountsByShopService(shopIDStr string) ([]models.Discount, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}
	return repositories.ListDiscountsByShop(shopID)
}

func UpdateDiscountService(idStr string, upd bson.M) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid discount ID")
	}
	// remove immutable fields
	delete(upd, "id")
	delete(upd, "_id")
	// Allow updating usage tracking fields for proper usage recording
	// delete(upd, "current_usage")
	// delete(upd, "usage_tracking")

	if startRaw, ok := upd["start_at"]; ok {
		if startTime, ok2 := startRaw.(time.Time); ok2 {
			if endRaw, exists := upd["end_at"]; exists {
				if endTime, ok3 := endRaw.(time.Time); ok3 && endTime.Before(startTime) {
					return errors.New("endAt must be after startAt")
				}
			}
		}
	}
	upd["updated_at"] = time.Now()
	_, err = repositories.UpdateDiscount(id, upd)
	return err
}

func DeleteDiscountService(idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid discount ID")
	}
	_, err = repositories.DeleteDiscount(id)
	return err
}

// ValidateDiscountForCustomer validates if a customer can use a discount
func ValidateDiscountForCustomer(discount *models.Discount, customerID primitive.ObjectID, customerSegmentIDs []primitive.ObjectID) error {
	// Check if discount is active
	if !discount.IsActive() {
		return ErrDiscountNotActive
	}

	// Check eligibility
	if !discount.IsEligible(customerID, customerSegmentIDs) {
		return ErrDiscountNotEligible
	}

	// Check usage limits
	if !discount.CanUse(customerID) {
		return ErrDiscountUsageLimitExceeded
	}

	return nil
}

// CanCustomerUseDiscount checks if a customer can use a specific discount
// This is a more comprehensive check that includes all validation rules
func CanCustomerUseDiscount(discount *models.Discount, customerID primitive.ObjectID, customerSegmentIDs []primitive.ObjectID) (bool, error) {
	// Check if discount is active
	if !discount.IsActive() {
		return false, ErrDiscountNotActive
	}

	// Check eligibility
	if !discount.IsEligible(customerID, customerSegmentIDs) {
		return false, ErrDiscountNotEligible
	}

	// Check overall usage limit
	if discount.UsageLimit != nil && discount.CurrentUsage >= *discount.UsageLimit {
		return false, ErrDiscountUsageLimitExceeded
	}

	// Check per-customer limit
	if discount.PerCustomerLimit != nil {
		for _, usage := range discount.UsageTracking {
			if usage.CustomerID == customerID && usage.UsageCount >= *discount.PerCustomerLimit {
				return false, ErrDiscountUsageLimitExceeded
			}
		}
	}

	return true, nil
}

// ApplyDiscountToOrder applies a discount to an order and records usage
func ApplyDiscountToOrder(discount *models.Discount, order *models.Order, customerID primitive.ObjectID) error {
	// Validate discount for customer
	if err := ValidateDiscountForCustomer(discount, customerID, []primitive.ObjectID{}); err != nil {
		return err
	}

	// Calculate discount amount
	discountAmount := discount.CalculateDiscount(order.Subtotal)

	// Apply discount
	order.DiscountTotal = discountAmount
	order.Total = order.Subtotal - discountAmount

	// Ensure total doesn't go negative
	if order.Total < 0 {
		order.Total = 0
		order.DiscountTotal = order.Subtotal
	}

	// Record usage
	discount.RecordUsage(customerID, order.Subtotal)

	// Update discount in database
	_, err := repositories.UpdateDiscount(discount.ID, bson.M{
		"current_usage":  discount.CurrentUsage,
		"usage_tracking": discount.UsageTracking,
		"updated_at":     time.Now(),
	})

	return err
}

// For product-level: need collectionIDs slice if applicable
func GetActiveDiscountsForProductService(shopID, productID, variantID primitive.ObjectID, collectionIDs []primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveDiscountsForProduct(shopID, productID, variantID, collectionIDs)
}

// GetEligibleDiscountsForCustomer returns all discounts a customer is eligible for
func GetEligibleDiscountsForCustomer(shopID, customerID primitive.ObjectID, customerSegmentIDs []primitive.ObjectID) ([]models.Discount, error) {
	discounts, err := repositories.ListDiscountsByShop(shopID)
	if err != nil {
		return nil, err
	}

	var eligibleDiscounts []models.Discount
	for _, discount := range discounts {
		// Only include active discounts
		if discount.IsActive() {
			if err := ValidateDiscountForCustomer(&discount, customerID, customerSegmentIDs); err == nil {
				eligibleDiscounts = append(eligibleDiscounts, discount)
			}
		}
	}

	return eligibleDiscounts, nil
}

// GetCustomerSegmentIDs retrieves the segment IDs for a customer in a shop
func GetCustomerSegmentIDs(shopID, customerID primitive.ObjectID) ([]primitive.ObjectID, error) {
	segments, err := sellerRepo.GetSegmentsByCustomer(shopID, customerID)
	if err != nil {
		return nil, err
	}
	ids := make([]primitive.ObjectID, 0, len(segments))
	for _, seg := range segments {
		ids = append(ids, seg.ID)
	}
	return ids, nil
}

// ValidateUsageLimits checks if a customer can use a discount (deprecated, use ValidateDiscountForCustomer)
func ValidateUsageLimits(discount *models.Discount, customerID primitive.ObjectID) (bool, error) {
	if !discount.CanUse(customerID) {
		return false, ErrDiscountUsageLimitExceeded
	}
	return true, nil
}

// ApplyDiscountsToProduct applies the best applicable discount to the product and its variants.
func ApplyDiscountsToProduct(product *models.Product, discounts []models.Discount) {
	// Helper: get best discount for a variant or product
	getBestDiscount := func(productID, variantID primitive.ObjectID) *models.Discount {
		var best *models.Discount
		var maxSavings float64
		for i := range discounts {
			d := &discounts[i]
			if d.Category != models.DiscountCategoryProduct || !d.IsActive() {
				continue
			}
			applies := false
			for _, vid := range d.AppliesToVariants {
				if vid == variantID && !variantID.IsZero() {
					applies = true
					break
				}
			}
			if !applies {
				for _, pid := range d.AppliesToProducts {
					if pid == productID {
						applies = true
						break
					}
				}
			}
			if applies {
				var estimatedSavings float64
				if d.Type == models.DiscountTypePercentage {
					estimatedSavings = 100 * (d.Value / 100)
				} else {
					estimatedSavings = d.Value
				}
				if best == nil || estimatedSavings > maxSavings {
					best = d
					maxSavings = estimatedSavings
				}
			}
		}
		return best
	}

	if len(product.Variants) > 0 {
		for i := range product.Variants {
			v := &product.Variants[i]
			best := getBestDiscount(product.ID, v.VariantID)
			if best != nil {
				discounted := v.Price - best.CalculateDiscount(v.Price)
				v.DisplayPrice = &discounted
				id := best.ID.Hex()
				v.AppliedDiscountID = &id
			} else {
				v.DisplayPrice = &v.Price
				v.AppliedDiscountID = nil
			}
		}
	} else {
		// Simple product (no variants)
		best := getBestDiscount(product.ID, primitive.NilObjectID)
		if best != nil {
			discounted := product.Price - best.CalculateDiscount(product.Price)
			product.DisplayPrice = &discounted
			id := best.ID.Hex()
			product.AppliedDiscountID = &id
		} else {
			product.DisplayPrice = &product.Price
			product.AppliedDiscountID = nil
		}
	}
}

// GetBestProductDiscount returns the best discount for a product or variant.
func GetBestProductDiscount(productID, variantID primitive.ObjectID, discounts []models.Discount) *models.Discount {
	var best *models.Discount
	var maxSavings float64

	for i := range discounts {
		d := &discounts[i]
		if d.Category != models.DiscountCategoryProduct {
			continue
		}

		applies := false
		for _, pid := range d.AppliesToProducts {
			if pid == productID {
				applies = true
				break
			}
		}
		for _, vid := range d.AppliesToVariants {
			if vid == variantID {
				applies = true
				break
			}
		}

		if applies && d.IsActive() {
			// Calculate actual savings for comparison (not just the value)
			// For percentage discounts, we need to estimate based on a typical price
			// For fixed discounts, the value is the savings
			var estimatedSavings float64
			if d.Type == models.DiscountTypePercentage {
				// Estimate savings based on a typical product price (e.g., $100)
				// This is a rough estimate for comparison purposes
				estimatedSavings = 100 * (d.Value / 100)
			} else {
				// Fixed discount - the value is the savings
				estimatedSavings = d.Value
			}

			if best == nil || estimatedSavings > maxSavings {
				best = d
				maxSavings = estimatedSavings
			}
		}
	}
	return best
}

// GetDiscountUsageStats returns usage statistics for a discount
func GetDiscountUsageStats(discountID string) (map[string]interface{}, error) {
	discount, err := GetDiscountByIDService(discountID)
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total_usage":      discount.CurrentUsage,
		"usage_limit":      discount.UsageLimit,
		"usage_tracking":   discount.UsageTracking,
		"is_active":        discount.IsActive(),
		"eligibility_type": discount.EligibilityType,
	}

	if discount.UsageLimit != nil {
		stats["usage_percentage"] = float64(discount.CurrentUsage) / float64(*discount.UsageLimit) * 100
	}

	return stats, nil
}

// ValidateDiscountValue validates the discount value based on type
func ValidateDiscountValue(discountType models.DiscountType, value float64) error {
	switch discountType {
	case models.DiscountTypeFixed:
		if value <= 0 {
			return errors.New("fixed discount value must be positive")
		}
	case models.DiscountTypePercentage:
		if value <= 0 || value > 100 {
			return errors.New("percentage discount value must be between 0 and 100")
		}
	default:
		return errors.New("invalid discount type")
	}
	return nil
}

// RecordDiscountUsageAtomic safely records discount usage with atomic operations
// This prevents race conditions when multiple orders use the same discount
func RecordDiscountUsageAtomic(discountID string, customerID primitive.ObjectID, amount float64) error {
	id, err := primitive.ObjectIDFromHex(discountID)
	if err != nil {
		return errors.New("invalid discount ID")
	}

	// First, check if we can still use this discount
	discount, err := GetDiscountByIDService(discountID)
	if err != nil {
		return err
	}

	// Validate usage limits before recording
	if !discount.CanUse(customerID) {
		return ErrDiscountUsageLimitExceeded
	}

	// Use atomic operations to safely update usage
	now := time.Now()

	// Find existing usage record for this customer
	found := false
	for i := range discount.UsageTracking {
		if discount.UsageTracking[i].CustomerID == customerID {
			// Update existing record
			discount.UsageTracking[i].UsageCount++
			discount.UsageTracking[i].LastUsedAt = now
			discount.UsageTracking[i].TotalSpent += amount
			found = true
			break
		}
	}

	if !found {
		// Create new usage record
		newUsage := models.DiscountUsage{
			CustomerID: customerID,
			UsageCount: 1,
			LastUsedAt: now,
			TotalSpent: amount,
		}
		discount.UsageTracking = append(discount.UsageTracking, newUsage)
	}

	// Single atomic operation that increments current_usage and updates usage_tracking
	atomicUpdate := bson.M{
		"$inc": bson.M{
			"current_usage": 1,
		},
		"$set": bson.M{
			"usage_tracking": discount.UsageTracking,
			"updated_at":     now,
		},
	}

	_, err = repositories.UpdateDiscountWithOperators(id, atomicUpdate)
	return err
}
