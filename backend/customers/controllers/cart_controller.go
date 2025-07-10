package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	sharedSvc "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// getCartIDOrCreate fetches or creates a cart for the authenticated customer and shop.
func getCartIDOrCreate(shopID, customerID primitive.ObjectID) (*models.Cart, error) {
	// For simplicity, assume one cart per customer per shop
	cart, err := sharedSvc.GetOrCreateCartService(shopID, customerID)
	return cart, err
}

// GetCart handles GET /shops/:shopSlug/cart
func GetCart(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := sharedSvc.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	cidVal, exists := c.Get("user_id")
	if !exists || cidVal == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}
	cidHex, ok := cidVal.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id"})
		return
	}
	customerID, _ := primitive.ObjectIDFromHex(cidHex)
	// Link customer to shop if not already linked
	_, _, _ = sharedSvc.LinkIfNotLinked(shop.ID, customerID)
	cart, err := sharedSvc.GetOrCreateCartService(shop.ID, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return cart with discount details
	cartService := sharedSvc.NewCartService()
	cartWithDetails, err := cartService.GetCartWithDiscountDetails(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartWithDetails)
}

// AddToCart handles POST /shops/:shopSlug/cart/items
func AddToCart(c *gin.Context) {
	shopSlug := c.Param("shopSlug")

	shop, err := sharedSvc.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	cidVal, exists := c.Get("user_id")
	if !exists || cidVal == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}
	cidHex, ok := cidVal.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id"})
		return
	}
	customerID, _ := primitive.ObjectIDFromHex(cidHex)
	// Link customer to shop if not already linked
	_, _, _ = sharedSvc.LinkIfNotLinked(shop.ID, customerID)
	cart, err := sharedSvc.GetOrCreateCartService(shop.ID, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		ProductID string `json:"product_id"`
		VariantID string `json:"variant_id"`
		Quantity  int    `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id format"})
		return
	}

	variantID := primitive.NilObjectID
	if req.VariantID != "" {
		variantID, err = primitive.ObjectIDFromHex(req.VariantID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid variant_id format"})
			return
		}
	}

	cartService := sharedSvc.NewCartService()
	if err := cartService.AddItem(cart, productID, variantID, req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_ = sharedSvc.SaveCartService(cart)

	// Return cart with discount details
	cartWithDetails, err := cartService.GetCartWithDiscountDetails(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartWithDetails)
}

// UpdateCartItem handles PUT /shops/:shopSlug/cart/items
func UpdateCartItem(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := sharedSvc.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	cidVal, exists := c.Get("user_id")
	if !exists || cidVal == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}
	cidHex, ok := cidVal.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id"})
		return
	}
	customerID, _ := primitive.ObjectIDFromHex(cidHex)
	// Link customer to shop if not already linked
	_, _, _ = sharedSvc.LinkIfNotLinked(shop.ID, customerID)
	cart, err := sharedSvc.GetOrCreateCartService(shop.ID, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req struct {
		ProductID string `json:"product_id"`
		VariantID string `json:"variant_id"`
		Quantity  int    `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productID, _ := primitive.ObjectIDFromHex(req.ProductID)
	variantID := primitive.NilObjectID
	if req.VariantID != "" {
		variantID, _ = primitive.ObjectIDFromHex(req.VariantID)
	}
	cartService := sharedSvc.NewCartService()
	if err := cartService.UpdateItem(cart, productID, variantID, req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_ = sharedSvc.SaveCartService(cart)

	// Return cart with discount details
	cartWithDetails, err := cartService.GetCartWithDiscountDetails(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartWithDetails)
}

// RemoveCartItem handles DELETE /shops/:shopSlug/cart/items
func RemoveCartItem(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := sharedSvc.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	cidVal, exists := c.Get("user_id")
	if !exists || cidVal == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}
	cidHex, ok := cidVal.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id"})
		return
	}
	customerID, _ := primitive.ObjectIDFromHex(cidHex)
	// Link customer to shop if not already linked
	_, _, _ = sharedSvc.LinkIfNotLinked(shop.ID, customerID)
	cart, err := sharedSvc.GetOrCreateCartService(shop.ID, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req struct {
		ProductID string `json:"product_id"`
		VariantID string `json:"variant_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productID, _ := primitive.ObjectIDFromHex(req.ProductID)
	variantID := primitive.NilObjectID
	if req.VariantID != "" {
		variantID, _ = primitive.ObjectIDFromHex(req.VariantID)
	}
	cartService := sharedSvc.NewCartService()
	if err := cartService.RemoveItem(cart, productID, variantID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_ = sharedSvc.SaveCartService(cart)

	// Return cart with discount details
	cartWithDetails, err := cartService.GetCartWithDiscountDetails(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartWithDetails)
}

// ClearCart handles POST /shops/:shopSlug/cart/clear
func ClearCart(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := sharedSvc.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	cidVal, exists := c.Get("user_id")
	if !exists || cidVal == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}
	cidHex, ok := cidVal.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id"})
		return
	}
	customerID, _ := primitive.ObjectIDFromHex(cidHex)
	// Link customer to shop if not already linked
	_, _, _ = sharedSvc.LinkIfNotLinked(shop.ID, customerID)
	cart, err := sharedSvc.GetOrCreateCartService(shop.ID, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cartService := sharedSvc.NewCartService()
	if err := cartService.ClearCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_ = sharedSvc.SaveCartService(cart)

	// Return cart with discount details
	cartWithDetails, err := cartService.GetCartWithDiscountDetails(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartWithDetails)
}
