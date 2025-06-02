package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ErrCartNotFound is returned when no cart exists for a given key.
var ErrCartNotFound = errors.New("cart not found")

// GetOrCreateCartForCustomer returns an existing Cart for (shopID,customerID) or creates a new one.
func GetOrCreateCartForCustomer(shopID, customerID primitive.ObjectID) (*models.Cart, error) {
	cart, err := repositories.GetCartByCustomerID(shopID, customerID)
	if err != nil {
		return nil, err
	}
	if cart != nil {
		return cart, nil
	}

	newCart := &models.Cart{
		ShopID:      shopID,
		CustomerID:  &customerID,
		Items:       []models.CartItem{},
		Currency:    "USD",
		CreatedAt:   time.Now(),
		LastUpdated: time.Now(),
	}
	_, err = repositories.CreateCart(newCart)
	if err != nil {
		return nil, err
	}
	return newCart, nil
}

// GetOrCreateCartForSession returns or creates a guest Cart by sessionID.
func GetOrCreateCartForSession(shopID primitive.ObjectID, sessionID string) (*models.Cart, error) {
	cart, err := repositories.GetCartBySessionID(shopID, sessionID)
	if err != nil {
		return nil, err
	}
	if cart != nil {
		return cart, nil
	}

	newCart := &models.Cart{
		ShopID:      shopID,
		SessionID:   &sessionID,
		Items:       []models.CartItem{},
		Currency:    "USD",
		CreatedAt:   time.Now(),
		LastUpdated: time.Now(),
	}
	_, err = repositories.CreateCart(newCart)
	if err != nil {
		return nil, err
	}
	return newCart, nil
}

// AddItem adds (or updates) a CartItem in the Cart.
// - Snaps product/variant data
// - Ensures stock
// - Recalculates totals
func AddItem(
	cart *models.Cart,
	productID, variantID primitive.ObjectID,
	quantity int,
) (*models.Cart, error) {
	// 1) Load product & variant via service
	prod, err := GetProductByIDService(productID.Hex())
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, errors.New("product not found")
	}

	// 2) Find the specified variant
	var selVariant *models.Variant
	for idx := range prod.Variants {
		if prod.Variants[idx].VariantID == variantID {
			selVariant = &prod.Variants[idx]
			break
		}
	}
	if selVariant == nil {
		return nil, errors.New("variant not found")
	}
	if selVariant.Stock < quantity {
		return nil, errors.New("insufficient stock for variant")
	}

	// 3) Build CartItem snapshot
	lineTotal := selVariant.Price * float64(quantity)
	item := &models.CartItem{
		ProductID:          productID,
		VariantID:          variantID,
		ProductName:        prod.Name,
		VariantOptions:     selVariant.Options,
		Image:              selVariant.Image,
		UnitPrice:          selVariant.Price,
		Quantity:           quantity,
		LineTotal:          lineTotal,
		DiscountAmount:     0,
		FinalLineTotal:     lineTotal,
		AppliedDiscountIDs: []primitive.ObjectID{},
	}

	// 4) Add or update this item in Mongo
	if _, err := repositories.AddOrUpdateCartItem(cart.ID, item); err != nil {
		return nil, err
	}

	// 5) Reload cart, recalc, persist
	return persistAndRecalc(cart.ID, cart.CustomerID)
}

// RemoveItem removes a single CartItem from the Cart.
func RemoveItem(cart *models.Cart, itemID primitive.ObjectID) (*models.Cart, error) {
	if _, err := repositories.RemoveCartItem(cart.ID, itemID); err != nil {
		return nil, err
	}
	return persistAndRecalc(cart.ID, cart.CustomerID)
}

// UpdateQuantity updates the quantity of an existing CartItem.
func UpdateQuantity(cart *models.Cart, itemID primitive.ObjectID, newQty int) (*models.Cart, error) {
	// 1) Locate the existing item in cart.Items
	var found *models.CartItem
	for idx := range cart.Items {
		if cart.Items[idx].ID == itemID {
			found = &cart.Items[idx]
			break
		}
	}
	if found == nil {
		return nil, errors.New("cart item not found")
	}

	// 2) Re‐fetch product & variant to re‐check stock
	prod, err := GetProductByIDService(found.ProductID.Hex())
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, errors.New("product not found")
	}

	var selVariant *models.Variant
	for idx := range prod.Variants {
		if prod.Variants[idx].VariantID == found.VariantID {
			selVariant = &prod.Variants[idx]
			break
		}
	}
	if selVariant == nil {
		return nil, errors.New("variant not found")
	}
	if selVariant.Stock < newQty {
		return nil, errors.New("insufficient stock")
	}

	// 3) Recompute snapshot on that item
	newLineTotal := selVariant.Price * float64(newQty)
	found.Quantity = newQty
	found.LineTotal = newLineTotal
	found.FinalLineTotal = newLineTotal
	found.DiscountAmount = 0
	found.AppliedDiscountIDs = []primitive.ObjectID{}

	if _, err := repositories.AddOrUpdateCartItem(cart.ID, found); err != nil {
		return nil, err
	}

	return persistAndRecalc(cart.ID, cart.CustomerID)
}

// ClearCart removes all items from the cart, resets totals.
func ClearCart(cart *models.Cart) (*models.Cart, error) {
	if _, err := repositories.ClearCartItems(cart.ID); err != nil {
		return nil, err
	}

	updatedCart, err := repositories.GetCartByID(cart.ID)
	if err != nil {
		return nil, err
	}
	if updatedCart == nil {
		return nil, ErrCartNotFound
	}

	// Reset all totals and applied discounts
	updatedCart.Subtotal = 0
	updatedCart.TotalDiscounts = 0
	updatedCart.ShippingCost = 0
	updatedCart.TaxAmount = 0
	updatedCart.GrandTotal = 0
	updatedCart.AppliedDiscountIDs = []primitive.ObjectID{}

	if _, err := repositories.UpdateCart(updatedCart); err != nil {
		return nil, err
	}
	return updatedCart, nil
}

// persistAndRecalc is a helper: reloads the cart, runs recalcCart(), and persists.
func persistAndRecalc(cartID primitive.ObjectID, customerID *primitive.ObjectID) (*models.Cart, error) {
	updatedCart, err := repositories.GetCartByID(cartID)
	if err != nil {
		return nil, err
	}
	if updatedCart == nil {
		return nil, ErrCartNotFound
	}

	if err := recalcCart(updatedCart, customerID); err != nil {
		return nil, err
	}
	if _, err := repositories.UpdateCart(updatedCart); err != nil {
		return nil, err
	}
	return updatedCart, nil
}

// recalcCart recalculates Subtotal, Discounts, ShippingCost, TaxAmount, and GrandTotal.
// Pass in customerID (or nil) to filter allowed‐customer discounts.
func recalcCart(cart *models.Cart, customerID *primitive.ObjectID) error {
	subtotal := 0.0
	totalDiscount := 0.0

	// 1) Recalculate each CartItem
	for idx := range cart.Items {
		item := &cart.Items[idx]
		item.LineTotal = item.UnitPrice * float64(item.Quantity)
		item.DiscountAmount = 0
		item.FinalLineTotal = item.LineTotal
		item.AppliedDiscountIDs = []primitive.ObjectID{}

		// 1a) Product‐level discounts
		prodDiscounts, err := GetActiveDiscountsForProductService(
			cart.ShopID,
			item.ProductID,
			item.VariantID,
			getCustomerOrNil(customerID),
		)
		if err != nil {
			return err
		}
		for _, d := range prodDiscounts {
			if d.Category == models.DiscountCategoryProduct {
				// Check if this Discount applies to this product/variant
				matches := false
				for _, pid := range d.AppliesToProducts {
					if pid == item.ProductID {
						matches = true
						break
					}
				}
				for _, vid := range d.AppliesToVariants {
					if vid == item.VariantID {
						matches = true
						break
					}
				}
				if !matches {
					continue
				}
				// Compute discount amount
				var amt float64
				if d.Type == models.DiscountTypeFixed {
					amt = d.Value * float64(item.Quantity)
				} else if d.Type == models.DiscountTypePercentage {
					amt = item.LineTotal * (d.Value / 100.0)
				}
				item.DiscountAmount += amt
				item.FinalLineTotal = item.LineTotal - item.DiscountAmount
				item.AppliedDiscountIDs = append(item.AppliedDiscountIDs, d.ID)
			}
			// TODO: handle BuyXGetY (e.g. automatically adding a free CartItem)
		}

		subtotal += item.LineTotal
		totalDiscount += item.DiscountAmount
	}

	// 2) Order‐level discounts
	orderDiscounts, err := GetActiveOrderDiscountsForShopService(cart.ShopID)
	if err != nil {
		return err
	}
	for _, d := range orderDiscounts {
		if d.MinimumOrderSubtotal != nil && subtotal < *d.MinimumOrderSubtotal {
			continue
		}
		var orderAmt float64
		if d.Type == models.DiscountTypeFixed {
			orderAmt = d.Value
		} else if d.Type == models.DiscountTypePercentage {
			orderAmt = (subtotal - totalDiscount) * (d.Value / 100.0)
		}
		totalDiscount += orderAmt
		cart.AppliedDiscountIDs = append(cart.AppliedDiscountIDs, d.ID)
	}

	// 3) Shipping‐level discounts (free shipping)
	shippingCost := cart.ShippingCost
	shipDiscounts, err := GetActiveShippingDiscountsForShopService(cart.ShopID)
	if err != nil {
		return err
	}
	for _, d := range shipDiscounts {
		if d.FreeShipping {
			if d.MinimumOrderForFreeShipping != nil && subtotal < *d.MinimumOrderForFreeShipping {
				continue
			}
			shippingCost = 0
			cart.AppliedDiscountIDs = append(cart.AppliedDiscountIDs, d.ID)
		}
	}

	// 4) Tax calculation (stubbed)
	tax := 0.0
	// e.g.: tax = TaxService.Calculate((subtotal - totalDiscount), shippingAddress)

	// 5) Final totals
	cart.Subtotal = subtotal
	cart.TotalDiscounts = totalDiscount
	cart.ShippingCost = shippingCost
	cart.TaxAmount = tax
	cart.GrandTotal = subtotal - totalDiscount + shippingCost + tax

	return nil
}

// getCustomerOrNil unwraps *primitive.ObjectID or returns primitive.NilObjectID if nil.
func getCustomerOrNil(customerID *primitive.ObjectID) primitive.ObjectID {
	if customerID == nil {
		return primitive.NilObjectID
	}
	return *customerID
}

// PersistAndRecalcCart recalculates all totals on `cart` and then saves it.
// Returns the newly reloaded Cart document or an error.
func PersistAndRecalcCart(cart *models.Cart) (*models.Cart, error) {
	// 1) Recalculate every line + order + shipping discounts + taxes + grand total
	if err := recalcCart(cart, cart.CustomerID); err != nil {
		return nil, err
	}

	// 2) Persist the updated cart in MongoDB
	if _, err := repositories.UpdateCart(cart); err != nil {
		return nil, err
	}

	// 3) Reload and return the fresh cart document
	updatedCart, err := repositories.GetCartByID(cart.ID)
	if err != nil {
		return nil, err
	}
	return updatedCart, nil
}
