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
	// Optionally enforce coupon_code uniqueness: attempt repo.GetDiscountByCouponCode
	// ...
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
	now := time.Now()
	if !d.Active || now.Before(d.StartAt) || now.After(d.EndAt) {
		return nil, ErrDiscountNotFound
	}
	return d, nil
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

func ValidateUsageLimits(discount *models.Discount, customerID primitive.ObjectID) (bool, error) {
	// stub or implement usage tracking
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
	if discount.Type == models.DiscountTypeFixed {
		discounted := product.Price - discount.Value
		if discounted < 0 {
			discounted = 0
		}
		// Only set DisplayPrice if the discount actually reduces the price
		if discounted < product.Price {
			product.DisplayPrice = &discounted
		}
	} else if discount.Type == models.DiscountTypePercentage {
		discounted := product.Price * (1 - discount.Value/100)
		if discounted < 0 {
			discounted = 0
		}
		// Only set DisplayPrice if the discount actually reduces the price
		if discounted < product.Price {
			product.DisplayPrice = &discounted
		}
	}
}

// Helper: apply discount to variant (sets DisplayPrice)
func applyDiscountToVariant(variant *models.Variant, discount *models.Discount) {
	if discount.Type == models.DiscountTypeFixed {
		discounted := variant.Price - discount.Value
		if discounted < 0 {
			discounted = 0
		}
		// Only set DisplayPrice if the discount actually reduces the price
		if discounted < variant.Price {
			variant.DisplayPrice = &discounted
		}
	} else if discount.Type == models.DiscountTypePercentage {
		discounted := variant.Price * (1 - discount.Value/100)
		if discounted < 0 {
			discounted = 0
		}
		// Only set DisplayPrice if the discount actually reduces the price
		if discounted < variant.Price {
			variant.DisplayPrice = &discounted
		}
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
			if d.Type == models.DiscountTypeFixed {
				totalDiscount += d.Value
			} else if d.Type == models.DiscountTypePercentage {
				totalDiscount += cart.Subtotal * (d.Value / 100)
			}
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
