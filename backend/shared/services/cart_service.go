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
		cart.Items = append(cart.Items, models.CartItem{
			ProductID: productID,
			VariantID: variantID,
			Quantity:  quantity,
		})
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
	// TODO: Validate discount usage limits, customer eligibility, etc.
	cart.AppliedDiscountIDs = append(cart.AppliedDiscountIDs, discount.ID)
	cart.LastUpdated = time.Now()
	return s.CalculateTotals(cart)
}

// CalculateTotals recalculates subtotal, discounts, and grand total for the cart.
func (s *CartService) CalculateTotals(cart *models.Cart) error {
	subtotal := 0.0
	totalDiscounts := 0.0
	for i := range cart.Items {
		item := &cart.Items[i]
		// Fetch product/variant price (pseudo-code, replace with actual repo call)
		product, err := repositories.GetProductByID(item.ProductID.Hex())
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
		subtotal += item.LineTotal
		// TODO: Apply item-level discounts here if needed
	}
	// TODO: Apply order-level discounts from cart.AppliedDiscountIDs
	// For now, just set subtotal and grand total
	cart.Subtotal = subtotal
	cart.TotalDiscounts = totalDiscounts
	cart.GrandTotal = subtotal - totalDiscounts + cart.ShippingCost + cart.TaxAmount
	return nil
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
