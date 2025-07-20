package services

import (
	"errors"

	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrCartNotFound = errors.New("cart not found")

// CartWithDiscountDetails represents cart data with detailed discount information
type CartWithDiscountDetails struct {
	*models.Cart
	ItemDiscountDetails []ItemDiscountDetail `json:"item_discount_details,omitempty"`
	DiscountStatuses    []DiscountStatus     `json:"discount_statuses,omitempty"`
}

// ItemDiscountDetail contains information about discounts applied to a specific item
type ItemDiscountDetail struct {
	ProductID  primitive.ObjectID      `json:"product_id"`
	VariantID  primitive.ObjectID      `json:"variant_id,omitempty"`
	DiscountID primitive.ObjectID      `json:"discount_id"`
	Name       string                  `json:"name"`
	Type       models.DiscountType     `json:"type"`
	Value      float64                 `json:"value"`
	Category   models.DiscountCategory `json:"category"`
	Amount     float64                 `json:"amount"`
	Status     DiscountStatus          `json:"status,omitempty"`
}

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
	return s.CalculateTotals(cart, *cart.CustomerID)
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
			return s.CalculateTotals(cart, *cart.CustomerID)
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
			return s.CalculateTotals(cart, *cart.CustomerID)
		}
	}
	return errors.New("item not found in cart")
}

// CalculateTotals recalculates subtotal, discounts, and grand total for the cart.
func (s *CartService) CalculateTotals(cart *models.Cart, customerID primitive.ObjectID) error {
	subtotal := 0.0
	totalItemDiscounts := 0.0

	// Get customer segments for proper discount validation
	customerSegmentIDs, err := s.getCustomerSegmentIDs(cart.ShopID, customerID)
	if err != nil {
		// Log error but continue with empty segments
		customerSegmentIDs = []primitive.ObjectID{}
	}

	// Calculate subtotal and apply item-level discounts
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

		// Get collection IDs for this product (for future collection support)
		collectionIDs, err := GetCollectionIDsForProduct(item.ProductID)
		if err != nil {
			collectionIDs = []primitive.ObjectID{}
		}

		// Get active discounts for this product/variant
		discounts, err := GetActiveDiscountsForProductService(cart.ShopID, item.ProductID, item.VariantID, collectionIDs)
		if err == nil && len(discounts) > 0 {
			// Use the improved discount selection logic with status information
			bestDiscount, _ := GetBestEligibleDiscountForProduct(item.ProductID, item.VariantID, customerID, customerSegmentIDs, discounts)

			if bestDiscount != nil {
				itemDiscountAmount = bestDiscount.CalculateDiscountForQuantity(item.UnitPrice, item.Quantity)
				item.AppliedDiscountIDs = append(item.AppliedDiscountIDs, bestDiscount.ID)
			}
		}

		item.DiscountAmount = itemDiscountAmount
		item.FinalLineTotal = item.LineTotal - itemDiscountAmount
		totalItemDiscounts += itemDiscountAmount
		subtotal += item.LineTotal
	}

	// Calculate final totals (product-level discounts only)
	cart.Subtotal = subtotal
	cart.TotalDiscounts = totalItemDiscounts
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

// findBestEligibleDiscount finds the discount that provides the highest savings for a customer
func (s *CartService) findBestEligibleDiscount(discounts []models.Discount, unitPrice float64, quantity int, customerID primitive.ObjectID, customerSegmentIDs []primitive.ObjectID) *models.Discount {
	var bestDiscount *models.Discount
	bestSavings := 0.0

	for i := range discounts {
		discount := &discounts[i]
		if discount.IsActive() {
			// Use the improved validation function
			canUse, err := CanCustomerUseDiscount(discount, customerID, customerSegmentIDs)

			if err == nil && canUse {
				// Calculate actual savings for this discount using the correct method
				savings := discount.CalculateDiscountForQuantity(unitPrice, quantity)

				if savings > bestSavings {
					bestSavings = savings
					bestDiscount = discount
				}
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

// GetCartWithDiscountDetails returns cart with detailed discount information
func (s *CartService) GetCartWithDiscountDetails(cart *models.Cart) (*CartWithDiscountDetails, error) {
	if cart == nil {
		return nil, ErrCartNotFound
	}

	result := &CartWithDiscountDetails{
		Cart: cart,
	}

	// Get customer segments for proper validation
	customerSegmentIDs, err := s.getCustomerSegmentIDs(cart.ShopID, *cart.CustomerID)
	if err != nil {
		customerSegmentIDs = []primitive.ObjectID{}
	}

	// Get item-level discount details and collect all discount statuses
	for _, item := range cart.Items {
		for _, discountID := range item.AppliedDiscountIDs {
			discount, err := GetDiscountByIDService(discountID.Hex())
			if err == nil && discount != nil && discount.IsActive() {
				// Use the improved validation function
				if canUse, err := CanCustomerUseDiscount(discount, *cart.CustomerID, customerSegmentIDs); err == nil && canUse {
					// Calculate the actual discount amount for this item
					discountAmount := discount.CalculateDiscountForQuantity(item.UnitPrice, item.Quantity)

					// Get detailed status for this discount
					status := GetDiscountStatusForCustomer(discount, *cart.CustomerID)

					result.ItemDiscountDetails = append(result.ItemDiscountDetails, ItemDiscountDetail{
						ProductID:  item.ProductID,
						VariantID:  item.VariantID,
						DiscountID: discount.ID,
						Name:       discount.Name,
						Type:       discount.Type,
						Value:      discount.Value,
						Category:   discount.Category,
						Amount:     discountAmount,
						Status:     status,
					})

					// Add to overall discount statuses
					result.DiscountStatuses = append(result.DiscountStatuses, status)
				}
			}
		}

		// Also collect statuses for all available discounts for this item (even if not applied)
		collectionIDs, err := GetCollectionIDsForProduct(item.ProductID)
		if err != nil {
			collectionIDs = []primitive.ObjectID{}
		}

		discounts, err := GetActiveDiscountsForProductService(cart.ShopID, item.ProductID, item.VariantID, collectionIDs)
		if err == nil && len(discounts) > 0 {
			_, allStatuses := GetBestEligibleDiscountForProduct(item.ProductID, item.VariantID, *cart.CustomerID, customerSegmentIDs, discounts)

			// Add unique statuses to the result
			for _, status := range allStatuses {
				// Check if this status is already in the result
				found := false
				for _, existingStatus := range result.DiscountStatuses {
					if existingStatus.DiscountID == status.DiscountID {
						found = true
						break
					}
				}
				if !found {
					result.DiscountStatuses = append(result.DiscountStatuses, status)
				}
			}
		}
	}

	return result, nil
}

// getCustomerSegmentIDs retrieves the segment IDs for a customer
func (s *CartService) getCustomerSegmentIDs(shopID, customerID primitive.ObjectID) ([]primitive.ObjectID, error) {
	// Use the shared service function
	return GetCustomerSegmentIDs(shopID, customerID)
}
