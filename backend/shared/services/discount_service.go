// shared/services/discount_service.go
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

// CreateDiscountService creates a new Discount in MongoDB, assigning ID/timestamps.
func CreateDiscountService(d *models.Discount) (*models.Discount, error) {
	// Basic validation
	if d.Name == "" {
		return nil, errors.New("discount name is required")
	}
	if d.ShopID.IsZero() {
		return nil, errors.New("shop ID is required")
	}
	if d.Category == "" {
		return nil, errors.New("discount category is required")
	}
	// Validate time range
	if !d.StartAt.IsZero() && !d.EndAt.IsZero() && d.EndAt.Before(d.StartAt) {
		return nil, errors.New("endAt must be after startAt")
	}
	// Call repository
	res, err := repositories.CreateDiscount(d)
	if err != nil {
		return nil, err
	}
	// Set generated ID
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		d.ID = oid
	}
	return d, nil
}

// GetDiscountByIDService retrieves a Discount by its hex ID string.
func GetDiscountByIDService(idHex string) (*models.Discount, error) {
	oid, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, ErrDiscountNotFound
	}
	d, err := repositories.GetDiscountByID(oid)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, ErrDiscountNotFound
	}
	return d, nil
}

// ListDiscountsByShopService lists all discounts for a given shop ID (hex string).
func ListDiscountsByShopService(shopIDHex string) ([]models.Discount, error) {
	shopOID, err := primitive.ObjectIDFromHex(shopIDHex)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}
	return repositories.ListDiscountsByShop(shopOID)
}

// UpdateDiscountService updates fields of a discount by its hex ID string.
// upd should use BSON field names (e.g. "name", "start_at", "end_at", etc.).
func UpdateDiscountService(idHex string, upd bson.M) error {
	oid, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return errors.New("invalid discount ID")
	}
	// Prevent immutable fields
	delete(upd, "id")
	delete(upd, "_id")
	delete(upd, "shop_id")
	delete(upd, "seller_id")
	// Validate time fields if present: expects "start_at" and/or "end_at" in upd as time.Time
	if startRaw, ok := upd["start_at"]; ok {
		if startTime, ok2 := startRaw.(time.Time); ok2 {
			if endRaw, exists := upd["end_at"]; exists {
				if endTime, ok3 := endRaw.(time.Time); ok3 && endTime.Before(startTime) {
					return errors.New("endAt must be after startAt")
				}
			}
		}
	}
	// Update timestamp
	upd["updated_at"] = time.Now()
	_, err = repositories.UpdateDiscount(oid, upd)
	return err
}

// DeleteDiscountService deletes a discount by its hex ID string.
func DeleteDiscountService(idHex string) error {
	oid, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return errors.New("invalid discount ID")
	}
	_, err = repositories.DeleteDiscount(oid)
	return err
}

// GetDiscountByCouponCodeService looks up a discount by coupon code,
// verifies active and within date range.
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

// (Optionally) other application-level methods, e.g., GetActiveDiscountsForShopService, etc.
// For example:
func GetActiveDiscountsForShopService(shopOID primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveDiscountsForShop(shopOID)
}
func GetActiveDiscountsForProductService(
	shopID, productID, variantID, customerID primitive.ObjectID,
) ([]models.Discount, error) {
	return repositories.GetActiveDiscountsForProduct(shopID, productID, variantID, customerID)
}
func GetActiveOrderDiscountsForShopService(shopID primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveOrderDiscountsForShop(shopID)
}
func GetActiveShippingDiscountsForShopService(shopID primitive.ObjectID) ([]models.Discount, error) {
	return repositories.GetActiveShippingDiscountsForShop(shopID)
}

// ValidateUsageLimits stub (implement usage tracking separately).
func ValidateUsageLimits(discount *models.Discount, customerID primitive.ObjectID) (bool, error) {
	// TODO: implement tracking
	return true, nil
}
