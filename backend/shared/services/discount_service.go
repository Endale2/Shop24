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
var ErrDiscountExpired = errors.New("discount has expired")
var ErrDiscountNotStarted = errors.New("discount has not started yet")
var ErrDiscountInvalidValue = errors.New("invalid discount value")
var ErrDiscountInvalidType = errors.New("invalid discount type")

func CreateDiscountService(d *models.Discount) (*models.Discount, error) {
	// Enhanced validation
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

	// Enhanced discount value validation
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
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

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

// ListDiscountsByShopPaginatedService returns paginated discounts for a shop
func ListDiscountsByShopPaginatedService(shopIDStr string, page, limit int) ([]models.Discount, int64, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return nil, 0, errors.New("invalid shop ID")
	}
	return repositories.ListDiscountsByShopPaginated(shopID, page, limit)
}

func UpdateDiscountService(idStr string, upd bson.M) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid discount ID")
	}

	// Remove immutable fields
	delete(upd, "id")
	delete(upd, "_id")
	delete(upd, "shop_id")
	delete(upd, "seller_id")
	delete(upd, "created_at")

	// Validate time constraints if updating time fields
	if startRaw, ok := upd["start_at"]; ok {
		if startTime, ok2 := startRaw.(time.Time); ok2 {
			if endRaw, exists := upd["end_at"]; exists {
				if endTime, ok3 := endRaw.(time.Time); ok3 && endTime.Before(startTime) {
					return errors.New("endAt must be after startAt")
				}
			}
		}
	}

	// Validate discount value if updating
	if valueRaw, ok := upd["value"]; ok {
		if value, ok2 := valueRaw.(float64); ok2 {
			typeRaw, typeExists := upd["type"]
			if typeExists {
				if discountType, ok3 := typeRaw.(string); ok3 {
					if err := ValidateDiscountValue(models.DiscountType(discountType), value); err != nil {
						return err
					}
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

// GetActiveDiscountsForProductService gets all active discounts for a product/variant
func GetActiveDiscountsForProductService(shopID, productID, variantID primitive.ObjectID, collectionIDs []primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveDiscountsForProduct(shopID, productID, variantID, collectionIDs)
}

// GetEligibleDiscountsForCustomer gets all discounts a customer is eligible for
func GetEligibleDiscountsForCustomer(shopID, customerID primitive.ObjectID, customerSegmentIDs []primitive.ObjectID) ([]models.Discount, error) {
	// Get all active discounts for the shop
	discounts, err := repositories.ListDiscountsByShop(shopID)
	if err != nil {
		return nil, err
	}

	// Filter to only eligible discounts
	var eligibleDiscounts []models.Discount
	for _, discount := range discounts {
		if discount.IsActive() && discount.IsEligible(customerID, customerSegmentIDs) && discount.CanUse(customerID) {
			eligibleDiscounts = append(eligibleDiscounts, discount)
		}
	}

	return eligibleDiscounts, nil
}

// GetCustomerSegmentIDs gets the segment IDs for a customer in a shop
func GetCustomerSegmentIDs(shopID, customerID primitive.ObjectID) ([]primitive.ObjectID, error) {
	// Get customer segments from the customer segment repository
	segments, err := sellerRepo.GetSegmentsByCustomer(shopID, customerID)
	if err != nil {
		return []primitive.ObjectID{}, nil // Return empty array if no segments found
	}

	var segmentIDs []primitive.ObjectID
	for _, segment := range segments {
		segmentIDs = append(segmentIDs, segment.ID)
	}

	return segmentIDs, nil
}

// ValidateUsageLimits validates if a discount can be used by a customer
func ValidateUsageLimits(discount *models.Discount, customerID primitive.ObjectID) (bool, error) {
	return discount.CanUse(customerID), nil
}

// ApplyDiscountsToProduct applies the best discount to a product
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
			// Check variant-specific discounts first
			for _, vid := range d.AppliesToVariants {
				if vid == variantID && !variantID.IsZero() {
					applies = true
					break
				}
			}
			// Then check product-level discounts
			if !applies {
				for _, pid := range d.AppliesToProducts {
					if pid == productID {
						applies = true
						break
					}
				}
			}

			if applies {
				// Calculate actual savings for comparison
				var estimatedSavings float64
				if d.Type == models.DiscountTypePercentage {
					// For percentage discounts, estimate based on typical price
					estimatedSavings = 100 * (d.Value / 100)
				} else {
					// For fixed discounts, the value is the savings
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

	// Apply discounts to variants
	if len(product.Variants) > 0 {
		for i := range product.Variants {
			v := &product.Variants[i]
			best := getBestDiscount(product.ID, v.VariantID)
			if best != nil {
				discounted := v.Price - best.CalculateDiscount(v.Price)
				// Only set display price if it's actually discounted
				if discounted < v.Price {
					v.DisplayPrice = &discounted
					id := best.ID.Hex()
					v.AppliedDiscountID = &id
				} else {
					v.DisplayPrice = &v.Price
					v.AppliedDiscountID = nil
				}
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
			// Only set display price if it's actually discounted
			if discounted < product.Price {
				product.DisplayPrice = &discounted
				id := best.ID.Hex()
				product.AppliedDiscountID = &id
			} else {
				product.DisplayPrice = &product.Price
				product.AppliedDiscountID = nil
			}
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
		// Check variant-specific discounts first
		for _, vid := range d.AppliesToVariants {
			if vid == variantID && !variantID.IsZero() {
				applies = true
				break
			}
		}
		// Then check product-level discounts
		if !applies {
			for _, pid := range d.AppliesToProducts {
				if pid == productID {
					applies = true
					break
				}
			}
		}

		if applies && d.IsActive() {
			// Calculate actual savings for comparison
			var estimatedSavings float64
			if d.Type == models.DiscountTypePercentage {
				// Estimate savings based on a typical product price (e.g., $100)
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
			return ErrDiscountInvalidValue
		}
	case models.DiscountTypePercentage:
		if value <= 0 || value > 100 {
			return ErrDiscountInvalidValue
		}
	default:
		return ErrDiscountInvalidType
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

// ValidateDiscountForProduct validates if a discount applies to a specific product/variant
func ValidateDiscountForProduct(discount *models.Discount, productID, variantID primitive.ObjectID) bool {
	// Check if discount applies to this product/variant
	if !variantID.IsZero() {
		// Check variant-specific application
		for _, vid := range discount.AppliesToVariants {
			if vid == variantID {
				return true
			}
		}
	}

	// Check product-level application
	for _, pid := range discount.AppliesToProducts {
		if pid == productID {
			return true
		}
	}

	return false
}

// DiscountStatus provides detailed information about discount availability
type DiscountStatus struct {
	DiscountID           primitive.ObjectID `json:"discount_id"`
	Name                 string             `json:"name"`
	IsAvailable          bool               `json:"is_available"`
	IsDisabled           bool               `json:"is_disabled"`
	CustomerLimitHit     bool               `json:"customer_limit_hit"`
	TotalLimitHit        bool               `json:"total_limit_hit"`
	CustomerUsage        int                `json:"customer_usage"`
	CustomerLimit        *int               `json:"customer_limit"`
	TotalUsage           int                `json:"total_usage"`
	TotalLimit           *int               `json:"total_limit"`
	RemainingForCustomer *int               `json:"remaining_for_customer"`
	RemainingTotal       *int               `json:"remaining_total"`
	Reason               string             `json:"reason,omitempty"`
}

// GetDiscountStatusForCustomer provides detailed status of a discount for a customer
func GetDiscountStatusForCustomer(discount *models.Discount, customerID primitive.ObjectID) DiscountStatus {
	status := DiscountStatus{
		DiscountID:  discount.ID,
		Name:        discount.Name,
		IsAvailable: true,
		IsDisabled:  false,
		TotalUsage:  discount.CurrentUsage,
		TotalLimit:  discount.UsageLimit,
	}

	// Check if discount is active
	if !discount.IsActive() {
		status.IsAvailable = false
		status.IsDisabled = true
		status.Reason = "Discount is not active"
		return status
	}

	// Check total usage limit
	if discount.UsageLimit != nil {
		if discount.CurrentUsage >= *discount.UsageLimit {
			status.IsAvailable = false
			status.IsDisabled = true
			status.TotalLimitHit = true
			status.Reason = "Total usage limit reached"
			return status
		}
		remaining := *discount.UsageLimit - discount.CurrentUsage
		status.RemainingTotal = &remaining
	}

	// Check per-customer limit
	if discount.PerCustomerLimit != nil {
		status.CustomerLimit = discount.PerCustomerLimit

		// Find customer usage
		for _, usage := range discount.UsageTracking {
			if usage.CustomerID == customerID {
				status.CustomerUsage = usage.UsageCount
				if usage.UsageCount >= *discount.PerCustomerLimit {
					status.IsAvailable = false
					status.CustomerLimitHit = true
					status.Reason = "You have reached your usage limit for this discount"
					return status
				}
				remaining := *discount.PerCustomerLimit - usage.UsageCount
				status.RemainingForCustomer = &remaining
				break
			}
		}

		// If customer not found in tracking, they haven't used it yet
		if status.CustomerUsage == 0 {
			remaining := *discount.PerCustomerLimit
			status.RemainingForCustomer = &remaining
		}
	}

	return status
}

// GetBestEligibleDiscountForProduct gets the best eligible discount for a product/variant and customer
// Now returns detailed status information
func GetBestEligibleDiscountForProduct(productID, variantID primitive.ObjectID, customerID primitive.ObjectID, customerSegmentIDs []primitive.ObjectID, discounts []models.Discount) (*models.Discount, []DiscountStatus) {
	var best *models.Discount
	var maxSavings float64
	var allStatuses []DiscountStatus

	for i := range discounts {
		d := &discounts[i]

		// Check if discount applies to this product/variant
		if !ValidateDiscountForProduct(d, productID, variantID) {
			continue
		}

		// Get detailed status for this discount
		status := GetDiscountStatusForCustomer(d, customerID)
		allStatuses = append(allStatuses, status)

		// Only consider available discounts
		if !status.IsAvailable {
			continue
		}

		// Check if customer is eligible
		if !d.IsEligible(customerID, customerSegmentIDs) {
			continue
		}

		// Calculate estimated savings for comparison
		var estimatedSavings float64
		if d.Type == models.DiscountTypePercentage {
			// Estimate based on typical price
			estimatedSavings = 100 * (d.Value / 100)
		} else {
			estimatedSavings = d.Value
		}

		if best == nil || estimatedSavings > maxSavings {
			best = d
			maxSavings = estimatedSavings
		}
	}

	return best, allStatuses
}
