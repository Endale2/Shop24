// customers/controllers/storefront_cart_controller.go
package controllers

import (
	"errors"
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type addItemInput struct {
	ProductID string `json:"productId" binding:"required"`
	VariantID string `json:"variantId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
}

type updateItemInput struct {
	Quantity int `json:"quantity" binding:"required,gt=0"`
}

type applyCouponInput struct {
	CouponCode string `json:"couponCode" binding:"required"`
}

// getShopIDFromSlug is a small helper to convert slug → ObjectID by first looking up the shop.
func getShopIDFromSlug(c *gin.Context) (primitive.ObjectID, error) {
	slug := c.Param("shopSlug")
	shopObj, err := services.GetShopBySlugService(slug)
	if err != nil {
		return primitive.NilObjectID, err
	}
	if shopObj == nil {
		return primitive.NilObjectID, errors.New("shop not found")
	}
	return shopObj.ID, nil
}

// getCustomerIDFromContext returns the logged-in customer ID (if any), else nil
func getCustomerIDFromContext(c *gin.Context) *primitive.ObjectID {
	raw, exists := c.Get("customer_id")
	if !exists {
		return nil
	}
	hex, ok := raw.(string)
	if !ok {
		return nil
	}
	oid, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil
	}
	return &oid
}

// getSessionIDFromCookie returns a session identifier for guest users (or empty if none).
func getSessionIDFromCookie(c *gin.Context) string {
	// You might store a cookie named "cart_session_id". If absent, generate or return empty.
	sessionID, _ := c.Cookie("cart_session_id")
	return sessionID
}

// GetOrCreateCart returns an existing cart (for logged‐in customer or guest session).
// GET /shops/:shopSlug/cart
func GetOrCreateCart(c *gin.Context) {
	shopID, err := getShopIDFromSlug(c)
	if err != nil {
		if err.Error() == "shop not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If customer is logged in, use customer‐based cart
	if custID := getCustomerIDFromContext(c); custID != nil {
		cart, err := services.GetOrCreateCartForCustomer(shopID, *custID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart"})
			return
		}
		c.JSON(http.StatusOK, cart)
		return
	}

	// Else, treat as guest cart by sessionID
	sessionID := getSessionIDFromCookie(c)
	if sessionID == "" {
		// If no cookie, you may generate a new session ID and set cookie
		newSession := primitive.NewObjectID().Hex()
		c.SetCookie("cart_session_id", newSession, 3600*24*30, "/", "", false, true)
		sessionID = newSession
	}
	cart, err := services.GetOrCreateCartForSession(shopID, sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

// AddItemToCart handles POST /shops/:shopSlug/cart/items
func AddItemToCart(c *gin.Context) {
	var input addItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	shopID, err := getShopIDFromSlug(c)
	if err != nil {
		if err.Error() == "shop not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Resolve or create a cart
	var (
		cart *models.Cart
	)
	if custID := getCustomerIDFromContext(c); custID != nil {
		cart, err = services.GetOrCreateCartForCustomer(shopID, *custID)
	} else {
		sessionID := getSessionIDFromCookie(c)
		if sessionID == "" {
			// new guest session
			newSession := primitive.NewObjectID().Hex()
			c.SetCookie("cart_session_id", newSession, 3600*24*30, "/", "", false, true)
			sessionID = newSession
		}
		cart, err = services.GetOrCreateCartForSession(shopID, sessionID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get or create cart"})
		return
	}

	// Convert productID and variantID to ObjectIDs
	prodOID, err := primitive.ObjectIDFromHex(input.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	varOID, err := primitive.ObjectIDFromHex(input.VariantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid variant ID"})
		return
	}

	// Add item (service will recalc and persist)
	updatedCart, err := services.AddItem(cart, prodOID, varOID, input.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCart)
}

// UpdateCartItem handles PATCH /shops/:shopSlug/cart/items/:itemID
func UpdateCartItem(c *gin.Context) {
	itemIDhex := c.Param("itemID")
	itemOID, err := primitive.ObjectIDFromHex(itemIDhex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart item ID"})
		return
	}

	var input updateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	shopID, err := getShopIDFromSlug(c)
	if err != nil {
		if err.Error() == "shop not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Load existing cart
	var cart *models.Cart
	if custID := getCustomerIDFromContext(c); custID != nil {
		cart, err = services.GetOrCreateCartForCustomer(shopID, *custID)
	} else {
		sessionID := getSessionIDFromCookie(c)
		cart, err = services.GetOrCreateCartForSession(shopID, sessionID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load cart"})
		return
	}

	updatedCart, err := services.UpdateQuantity(cart, itemOID, input.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCart)
}

// RemoveCartItem handles DELETE /shops/:shopSlug/cart/items/:itemID
func RemoveCartItem(c *gin.Context) {
	itemIDhex := c.Param("itemID")
	itemOID, err := primitive.ObjectIDFromHex(itemIDhex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart item ID"})
		return
	}

	shopID, err := getShopIDFromSlug(c)
	if err != nil {
		if err.Error() == "shop not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var cart *models.Cart
	if custID := getCustomerIDFromContext(c); custID != nil {
		cart, err = services.GetOrCreateCartForCustomer(shopID, *custID)
	} else {
		sessionID := getSessionIDFromCookie(c)
		cart, err = services.GetOrCreateCartForSession(shopID, sessionID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load cart"})
		return
	}

	updatedCart, err := services.RemoveItem(cart, itemOID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCart)
}

// ClearCart handles POST /shops/:shopSlug/cart/clear
func ClearCart(c *gin.Context) {
	shopID, err := getShopIDFromSlug(c)
	if err != nil {
		if err.Error() == "shop not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var cart *models.Cart
	if custID := getCustomerIDFromContext(c); custID != nil {
		cart, err = services.GetOrCreateCartForCustomer(shopID, *custID)
	} else {
		sessionID := getSessionIDFromCookie(c)
		cart, err = services.GetOrCreateCartForSession(shopID, sessionID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load cart"})
		return
	}

	updatedCart, err := services.ClearCart(cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCart)
}

// ApplyCoupon handles POST /shops/:shopSlug/cart/apply‐coupon
func ApplyCoupon(c *gin.Context) {
	var input applyCouponInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	shopID, err := getShopIDFromSlug(c)
	if err != nil {
		if err.Error() == "shop not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// First, look up the discount by coupon code
	discount, err := services.GetDiscountByCouponCodeService(input.CouponCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or expired coupon"})
		return
	}

	// Verify discount applies to this shop
	if discount.ShopID != shopID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "coupon does not belong to this shop"})
		return
	}

	// Load cart
	var cart *models.Cart
	if custID := getCustomerIDFromContext(c); custID != nil {
		cart, err = services.GetOrCreateCartForCustomer(shopID, *custID)
	} else {
		sessionID := getSessionIDFromCookie(c)
		cart, err = services.GetOrCreateCartForSession(shopID, sessionID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load cart"})
		return
	}

	// TODO: implement ValidateUsageLimits inside discount_service or here
	ok, err := services.ValidateUsageLimits(discount, getCustomerIDOrNil(cart))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error validating coupon usage"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "coupon usage limit reached"})
		return
	}

	// Mark this discount as “applied” by adding its ID to cart.AppliedDiscountIDs
	cart.AppliedDiscountIDs = append(cart.AppliedDiscountIDs, discount.ID)

	// Recalc & persist
	updatedCart, err := services.PersistAndRecalcCart(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to apply coupon"})
		return
	}

	c.JSON(http.StatusOK, updatedCart)
}

// CheckoutCart handles POST /shops/:shopSlug/cart/checkout
// NOTE: This is a placeholder. In a real system, you’d:
// 1) Re‐verify inventory for each item
// 2) Create an Order record with all line snapshots
// 3) Decrement inventory
// 4) Clear or archive the cart
// 5) Return the new order.
func CheckoutCart(c *gin.Context) {
	shopID, err := getShopIDFromSlug(c)
	if err != nil {
		if err.Error() == "shop not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Load cart
	var cart *models.Cart
	if custID := getCustomerIDFromContext(c); custID != nil {
		cart, err = services.GetOrCreateCartForCustomer(shopID, *custID)
	} else {
		sessionID := getSessionIDFromCookie(c)
		cart, err = services.GetOrCreateCartForSession(shopID, sessionID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load cart"})
		return
	}

	// Placeholder: return current cart totals
	c.JSON(http.StatusOK, gin.H{
		"message":      "checkout is not yet implemented",
		"cartSubtotal": cart.Subtotal,
		"cartTotal":    cart.GrandTotal,
	})
}

// getCustomerIDOrNil unwraps *primitive.ObjectID or returns primitive.NilObjectID
func getCustomerIDOrNil(cart *models.Cart) primitive.ObjectID {
	if cart.CustomerID != nil {
		return *cart.CustomerID
	}
	return primitive.NilObjectID
}
