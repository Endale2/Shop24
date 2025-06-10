package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ErrDiscountNotFound indicates no discount was found.
var ErrDiscountNotFound = errors.New("discount not found")

// GetDiscountByIDService retrieves a discount by its hex ID string.
// It converts hex to ObjectID and calls repositories.GetDiscountByID.
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

// GetDiscountByCouponCodeService looks up a discount by coupon code (exact match).
// It directly calls repositories.GetDiscountByCouponCode.
// The caller (e.g., controller) can then check active/time-range if desired.
// If you want to enforce active/time here, uncomment the time checks.
func GetDiscountByCouponCodeService(code string) (*models.Discount, error) {
	if code == "" {
		return nil, errors.New("coupon code is required")
	}
	d, err := repositories.GetDiscountByCouponCode(code)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, ErrDiscountNotFound
	}
	// Optionally validate active/time-range:
	now := time.Now()
	if !d.Active || now.Before(d.StartAt) || now.After(d.EndAt) {
		return nil, ErrDiscountNotFound
	}
	return d, nil
}

// GetActiveDiscountsForShopService wraps repositories.GetActiveDiscountsForShop.
// Accepts shop ID as hex string.
func GetActiveDiscountsForShopService(shopIDStr string) ([]models.Discount, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}
	return repositories.GetActiveDiscountsForShop(shopID)
}

// GetActiveDiscountsForProductService wraps repositories.GetActiveDiscountsForProduct.
// Accepts IDs as hex strings. Returns active “product” discounts applying to the given product/variant/customer.
func GetActiveDiscountsForProductService(
	shopIDStr, productIDStr, variantIDStr, customerIDStr string,
) ([]models.Discount, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}
	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}
	variantID, err := primitive.ObjectIDFromHex(variantIDStr)
	if err != nil {
		return nil, errors.New("invalid variant ID")
	}
	// customerID may be optional: if empty string, use NilObjectID
	var customerID primitive.ObjectID
	if customerIDStr != "" {
		cID, err := primitive.ObjectIDFromHex(customerIDStr)
		if err != nil {
			return nil, errors.New("invalid customer ID")
		}
		customerID = cID
	}
	return repositories.GetActiveDiscountsForProduct(shopID, productID, variantID, customerID)
}

// GetActiveOrderDiscountsForShopService wraps repositories.GetActiveOrderDiscountsForShop.
// Accepts shop ID hex.
func GetActiveOrderDiscountsForShopService(shopIDStr string) ([]models.Discount, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}
	return repositories.GetActiveOrderDiscountsForShop(shopID)
}

// GetActiveShippingDiscountsForShopService wraps repositories.GetActiveShippingDiscountsForShop.
// Accepts shop ID hex.
func GetActiveShippingDiscountsForShopService(shopIDStr string) ([]models.Discount, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}
	return repositories.GetActiveShippingDiscountsForShop(shopID)
}

// ValidateUsageLimits checks usage limits for a discount.
// Stub: in a real implementation you would:
// - Count total uses (global) against discount.UsageLimit
// - Count per-customer uses against PerCustomerLimit
// - If within limits, record/increment usage.
// Here we simply return true.
func ValidateUsageLimits(discount *models.Discount, customerID primitive.ObjectID) (bool, error) {
	// TODO: implement real usage-tracking
	return true, nil
}
