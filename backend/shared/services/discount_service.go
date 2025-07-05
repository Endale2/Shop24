package services

import (
	"errors"
	"time"

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

	// Set default eligibility type if not specified
	if d.EligibilityType == "" {
		d.EligibilityType = models.DiscountEligibilityAll
	}

	// Initialize usage tracking
	d.CurrentUsage = 0
	d.UsageTracking = []models.DiscountUsage{}

	// Optionally enforce coupon_code uniqueness: attempt repo.GetDiscountByCouponCode
	if d.CouponCode != "" {
		existing, err := repositories.GetDiscountByCouponCode(d.CouponCode)
		if err == nil && existing != nil {
			return nil, errors.New("coupon code already exists")
		}
	}

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
	// remove immutable
	delete(upd, "id")
	delete(upd, "_id")
	delete(upd, "current_usage")
	delete(upd, "usage_tracking")

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

func GetDiscountByCouponCodeService(code string) (*models.Discount, error) {
	d, err := repositories.GetDiscountByCouponCode(code)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, ErrDiscountNotFound
	}

	// Check if discount is active
	if !d.IsActive() {
		return nil, ErrDiscountNotActive
	}

	return d, nil
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

func GetActiveOrderDiscountsForShopService(shopID primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveOrderDiscountsForShop(shopID)
}

func GetActiveShippingDiscountsForShopService(shopID primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveShippingDiscountsForShop(shopID)
}

// GetEligibleDiscountsForCustomer returns all discounts a customer is eligible for
func GetEligibleDiscountsForCustomer(shopID, customerID primitive.ObjectID, customerSegmentIDs []primitive.ObjectID) ([]models.Discount, error) {
	discounts, err := repositories.ListDiscountsByShop(shopID)
	if err != nil {
		return nil, err
	}

	var eligibleDiscounts []models.Discount
	for _, discount := range discounts {
		if err := ValidateDiscountForCustomer(&discount, customerID, customerSegmentIDs); err == nil {
			eligibleDiscounts = append(eligibleDiscounts, discount)
		}
	}

	return eligibleDiscounts, nil
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
	// Apply to product price if product-level discount exists
	for _, d := range discounts {
		if d.Category == models.DiscountCategoryProduct {
			// If product is in AppliesToProducts
			for _, pid := range d.AppliesToProducts {
				if pid == product.ID {
					applyDiscountToProduct(product, &d)
				}
			}
			// If any variant is in AppliesToVariants
			for i := range product.Variants {
				for _, vid := range d.AppliesToVariants {
					if vid == product.Variants[i].VariantID {
						applyDiscountToVariant(&product.Variants[i], &d)
					}
				}
			}
		}
	}
}

// Helper: apply discount to product (sets DisplayPrice)
func applyDiscountToProduct(product *models.Product, discount *models.Discount) {
	discountAmount := discount.CalculateDiscount(product.Price)
	discounted := product.Price - discountAmount
	if discounted < 0 {
		discounted = 0
	}
	// Only set DisplayPrice if the discount actually reduces the price
	if discounted < product.Price {
		product.DisplayPrice = &discounted
	}
}

// Helper: apply discount to variant (sets DisplayPrice)
func applyDiscountToVariant(variant *models.Variant, discount *models.Discount) {
	discountAmount := discount.CalculateDiscount(variant.Price)
	discounted := variant.Price - discountAmount
	if discounted < 0 {
		discounted = 0
	}
	// Only set DisplayPrice if the discount actually reduces the price
	if discounted < variant.Price {
		variant.DisplayPrice = &discounted
	}
}

// ApplyDiscountsToCart applies order-level discounts to the cart.
func ApplyDiscountsToCart(cart *models.Cart, discounts []models.Discount) {
	totalDiscount := 0.0
	for _, d := range discounts {
		if d.Category == models.DiscountCategoryOrder {
			// Check minimum order subtotal
			if d.MinimumOrderSubtotal != nil && cart.Subtotal < *d.MinimumOrderSubtotal {
				continue
			}
			totalDiscount += d.CalculateDiscount(cart.Subtotal)
		}
	}
	cart.TotalDiscounts = totalDiscount
	cart.GrandTotal = cart.Subtotal - totalDiscount + cart.ShippingCost + cart.TaxAmount
}

// GetBestProductDiscount returns the best discount for a product or variant.
func GetBestProductDiscount(productID, variantID primitive.ObjectID, discounts []models.Discount) *models.Discount {
	var best *models.Discount
	var maxValue float64
	for _, d := range discounts {
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
		if applies {
			// Estimate value (percentage or fixed)
			value := d.Value
			if d.Type == models.DiscountTypePercentage {
				value = d.Value // percent, but for comparison
			}
			if best == nil || value > maxValue {
				best = &d
				maxValue = value
			}
		}
	}
	return best
}

// GetDiscountUsageStats returns usage statistics for a discount
func GetDiscountUsageStats(discountID string) (*models.Discount, error) {
	discount, err := GetDiscountByIDService(discountID)
	if err != nil {
		return nil, err
	}
	return discount, nil
}

// AddCustomersToDiscount adds specific customers to a discount's allowed list
func AddCustomersToDiscount(discountID string, customerIDs []string) error {
	discount, err := GetDiscountByIDService(discountID)
	if err != nil {
		return err
	}

	// Convert string IDs to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, idStr := range customerIDs {
		if oid, err := primitive.ObjectIDFromHex(idStr); err == nil {
			objectIDs = append(objectIDs, oid)
		}
	}

	// Add to existing allowed customers
	discount.AllowedCustomerIDs = append(discount.AllowedCustomerIDs, objectIDs...)
	discount.EligibilityType = models.DiscountEligibilitySpecific

	// Update in database
	_, err = repositories.UpdateDiscount(discount.ID, bson.M{
		"allowed_customers": discount.AllowedCustomerIDs,
		"eligibility_type":  discount.EligibilityType,
		"updated_at":        time.Now(),
	})

	return err
}

// AddSegmentsToDiscount adds customer segments to a discount's allowed list
func AddSegmentsToDiscount(discountID string, segmentIDs []string) error {
	discount, err := GetDiscountByIDService(discountID)
	if err != nil {
		return err
	}

	// Convert string IDs to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, idStr := range segmentIDs {
		if oid, err := primitive.ObjectIDFromHex(idStr); err == nil {
			objectIDs = append(objectIDs, oid)
		}
	}

	// Add to existing allowed segments
	discount.AllowedSegmentIDs = append(discount.AllowedSegmentIDs, objectIDs...)
	discount.EligibilityType = models.DiscountEligibilitySegment

	// Update in database
	_, err = repositories.UpdateDiscount(discount.ID, bson.M{
		"allowed_segments": discount.AllowedSegmentIDs,
		"eligibility_type": discount.EligibilityType,
		"updated_at":       time.Now(),
	})

	return err
}
