package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrDiscountNotFound = errors.New("discount not found")


// GetDiscountByCouponCodeService looks up a discount by coupon code.
func GetDiscountByCouponCodeService(code string) (*models.Discount, error) {
	d, err := repositories.GetDiscountByCouponCode(code)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, ErrDiscountNotFound
	}
	now := time.Now()
	if !d.Active || now.Before(d.StartAt) || now.After(d.EndAt) {
		return nil, ErrDiscountNotFound
	}
	return d, nil
}

// GetActiveDiscountsForProductService returns all product‐category discounts that apply.
func GetActiveDiscountsForProductService(
	shopID, productID, variantID, customerID primitive.ObjectID,
) ([]models.Discount, error) {
	return repositories.GetActiveDiscountsForProduct(shopID, productID, variantID, customerID)
}

// GetActiveOrderDiscountsForShopService returns all valid order‐category discounts for this shop.
func GetActiveOrderDiscountsForShopService(shopID primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveOrderDiscountsForShop(shopID)
}

// GetActiveShippingDiscountsForShopService returns all valid shipping‐category discounts for this shop.
func GetActiveShippingDiscountsForShopService(shopID primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveShippingDiscountsForShop(shopID)
}

// ValidateUsageLimits checks and updates usage counts (stub).
func ValidateUsageLimits(discount *models.Discount, customerID primitive.ObjectID) (bool, error) {
	// TODO: implement a “usage tracking” collection:
	//  - If discount.UsageLimit != nil, verify global usage < limit
	//  - If discount.PerCustomerLimit != nil, verify this cpn was not already used too many times by this customerID
	//  - If within limits, increment counters then return (true, nil)
	return true, nil
}
