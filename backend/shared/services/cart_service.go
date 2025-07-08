package services

import (
	"errors"
	
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrCartNotFound = errors.New("cart not found")

// CartService provides business logic for cart management.
type CartService struct{}

func NewCartService() *CartService {
	return &CartService{}
}

// AddItem adds a product/variant to the cart, or updates quantity if already present.
func (s *CartService) AddItem(cart *models.Cart, productID, variantID primitive.ObjectID, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	// Fetch product details to populate cart item
	product, err := GetProductByIDService(productID.Hex())
	if err != nil {
		return errors.New("product not found")
	}
	if product == nil {
		return errors.New("product not found")
	}

	found := false
	for i := range cart.Items {
		item := &cart.Items[i]
		if item.ProductID == productID && item.VariantID == variantID {
			item.Quantity += quantity
			found = true
			break
		}
	}

	if !found {
		// Create new cart item with product details
		cartItem := models.CartItem{
			ProductID:   productID,
			VariantID:   variantID,
			Quantity:    quantity,
			ProductName: product.Name,
			Image:       product.MainImage,
		}

		// Populate variant options if variant is selected
		if !variantID.IsZero() {
			for _, variant := range product.Variants {
				if variant.VariantID == variantID {
					cartItem.VariantOptions = make(map[string]string)
					for _, option := range variant.Options {
						cartItem.VariantOptions[option.Name] = option.Value
					}
					cartItem.Image = variant.Image // Use variant image if available
					break
				}
			}
		}

		cart.Items = append(cart.Items, cartItem)
	}

	cart.LastUpdated = time.Now()
	return s.CalculateTotals(cart)
}

// UpdateItem updates the quantity of a cart item.
func (s *CartService) UpdateItem(cart *models.Cart, productID, variantID primitive.ObjectID, quantity int) error {
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	for i := range cart.Items {
		item := &cart.Items[i]
		if item.ProductID == productID && item.VariantID == variantID {
			if quantity == 0 {
				// Remove item
				cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			} else {
				item.Quantity = quantity
			}
			cart.LastUpdated = time.Now()
			return s.CalculateTotals(cart)
		}
	}
	return errors.New("item not found in cart")
}

// RemoveItem removes a product/variant from the cart.
func (s *CartService) RemoveItem(cart *models.Cart, productID, variantID primitive.ObjectID) error {
	for i := range cart.Items {
		item := &cart.Items[i]
		if item.ProductID == productID && item.VariantID == variantID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			cart.LastUpdated = time.Now()
			return s.CalculateTotals(cart)
		}
	}
	return errors.New("item not found in cart")
}

// ApplyDiscountCode applies a discount code to the cart if valid.
func (s *CartService) ApplyDiscountCode(cart *models.Cart, code string) error {
	discount, err := GetDiscountByCouponCodeService(code)
	if err != nil {
		return err
	}

	// Check if discount is already applied
	for _, appliedID := range cart.AppliedDiscountIDs {
		if appliedID == discount.ID {
			return errors.New("discount code already applied")
		}
	}

	// TODO: Validate discount usage limits, customer eligibility, etc.
	cart.AppliedDiscountIDs = append(cart.AppliedDiscountIDs, discount.ID)
	cart.LastUpdated = time.Now()
	return s.CalculateTotals(cart)
}

// CalculateTotals recalculates subtotal, discounts, and grand total for the cart.
func (s *CartService) CalculateTotals(cart *models.Cart) error {
	subtotal := 0.0
	totalItemDiscounts := 0.0
	totalOrderDiscounts := 0.0

	// First pass: calculate subtotal and apply item-level discounts
	for i := range cart.Items {
		item := &cart.Items[i]
		// Fetch product/variant price
		product, err := GetProductByIDService(item.ProductID.Hex())
		if err != nil || product == nil {
			continue // skip missing products
		}

		var price float64
		if !item.VariantID.IsZero() {
			for _, v := range product.Variants {
				if v.VariantID == item.VariantID {
					price = v.Price
					break
				}
			}
		} else {
			price = product.Price
		}

		item.UnitPrice = price
		item.LineTotal = price * float64(item.Quantity)

		// Apply item-level discounts
		itemDiscountAmount := 0.0
		item.AppliedDiscountIDs = []primitive.ObjectID{} // Reset applied discounts

		// Get collection IDs for this product
		collectionIDs, err := GetCollectionIDsForProduct(item.ProductID)
		if err != nil {
			collectionIDs = []primitive.ObjectID{}
		}

		// Get active discounts for this product/variant
		discounts, err := GetActiveDiscountsForProductService(cart.ShopID, item.ProductID, item.VariantID, collectionIDs)
		if err == nil && len(discounts) > 0 {
			// Apply the best discount (highest value)
			bestDiscount := s.findBestDiscount(discounts, item.LineTotal)
			if bestDiscount != nil {
				itemDiscountAmount = bestDiscount.CalculateDiscount(item.LineTotal)
				item.AppliedDiscountIDs = append(item.AppliedDiscountIDs, bestDiscount.ID)
			}
		}

		item.DiscountAmount = itemDiscountAmount
		item.FinalLineTotal = item.LineTotal - itemDiscountAmount
		totalItemDiscounts += itemDiscountAmount
		subtotal += item.LineTotal
	}

	// Second pass: apply order-level discounts
	if len(cart.AppliedDiscountIDs) > 0 {
		for _, discountID := range cart.AppliedDiscountIDs {
			discount, err := GetDiscountByIDService(discountID.Hex())
			if err == nil && discount != nil && discount.Category == models.DiscountCategoryOrder {
				// Check minimum order subtotal if specified
				if discount.MinimumOrderSubtotal == nil || subtotal >= *discount.MinimumOrderSubtotal {
					orderDiscountAmount := discount.CalculateDiscount(subtotal - totalItemDiscounts)
					totalOrderDiscounts += orderDiscountAmount
				}
			}
		}
	}

	// Calculate final totals
	cart.Subtotal = subtotal
	cart.TotalDiscounts = totalItemDiscounts + totalOrderDiscounts
	cart.GrandTotal = subtotal - cart.TotalDiscounts + cart.ShippingCost + cart.TaxAmount

	return nil
}

// findBestDiscount finds the discount that provides the highest savings
func (s *CartService) findBestDiscount(discounts []models.Discount, price float64) *models.Discount {
	var bestDiscount *models.Discount
	bestSavings := 0.0

	for i := range discounts {
		discount := &discounts[i]
		if discount.IsActive() {
			savings := discount.CalculateDiscount(price)
			if savings > bestSavings {
				bestSavings = savings
				bestDiscount = discount
			}
		}
	}

	return bestDiscount
}

// ClearCart removes all items and discounts from the cart.
func (s *CartService) ClearCart(cart *models.Cart) error {
	cart.Items = nil
	cart.AppliedDiscountIDs = nil
	cart.Subtotal = 0
	cart.TotalDiscounts = 0
	cart.GrandTotal = 0
	cart.LastUpdated = time.Now()
	return nil
}

// GetOrCreateCartService fetches a cart for a shop/customer or creates one if not found.
func GetOrCreateCartService(shopID, customerID primitive.ObjectID) (*models.Cart, error) {
	cart, err := repositories.GetCartByCustomerID(shopID, customerID)
	if err != nil {
		return nil, err
	}
	if cart == nil {
		cart = &models.Cart{
			ShopID:      shopID,
			CustomerID:  &customerID,
			Items:       []models.CartItem{},
			Currency:    "USD",
			CreatedAt:   time.Now(),
			LastUpdated: time.Now(),
		}
		_, err := repositories.CreateCart(cart)
		if err != nil {
			return nil, err
		}
	}
	return cart, nil
}

// SaveCartService updates the cart in the database.
func SaveCartService(cart *models.Cart) error {
	_, err := repositories.UpdateCart(cart)
	return err
}
