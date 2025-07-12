// customers/controllers/order_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderItemRequest represents the minimal data needed from client for order placement
type OrderItemRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	VariantID string `json:"variant_id,omitempty"` // optional, empty string for no variant
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

// PlaceOrderRequest represents the complete request structure for placing an order
type PlaceOrderRequest struct {
	Items []OrderItemRequest `json:"items" binding:"required,min=1"`
}

// DebugProduct handles GET /shops/:shopSlug/debug/product/:productId
func DebugProduct(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	productID := c.Param("productId")

	// Lookup the shop by its slug
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Parse product ID
	pid, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	// Get customer ID from context if available
	var customerID *primitive.ObjectID
	if cidVal, exists := c.Get("user_id"); exists {
		if cidHex, ok := cidVal.(string); ok {
			if cid, err := primitive.ObjectIDFromHex(cidHex); err == nil {
				customerID = &cid
			}
		}
	}

	// Fetch product with discounts applied
	product, err := services.GetProductByIDWithDiscountsService(pid.Hex(), customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product: " + err.Error()})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	// Verify product belongs to this shop
	if product.ShopID != shop.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product does not belong to this shop"})
		return
	}

	// Build debug response
	variants := make([]map[string]interface{}, 0, len(product.Variants))
	for i, v := range product.Variants {
		variant := map[string]interface{}{
			"index":      i,
			"variant_id": v.VariantID.Hex(),
			"options":    v.Options,
			"price":      v.Price,
			"stock":      v.Stock,
			"image":      v.Image,
		}
		variants = append(variants, variant)
	}

	response := map[string]interface{}{
		"product_id":     product.ID.Hex(),
		"name":           product.Name,
		"shop_id":        product.ShopID.Hex(),
		"product_price":  product.Price,
		"product_stock":  product.Stock,
		"total_variants": len(product.Variants),
		"variants":       variants,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProductPrice handles POST /shops/:shopSlug/debug/product/:productId/price
// This is a temporary endpoint for testing - update product price
func UpdateProductPrice(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	productID := c.Param("productId")

	// Lookup the shop by its slug
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Parse product ID
	pid, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	// Parse request body
	var req struct {
		Price float64 `json:"price" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update product price
	_, err = services.UpdateProductService(pid.Hex(), bson.M{"price": req.Price})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product price updated successfully", "new_price": req.Price})
}

// FixVariantIDs handles POST /shops/:shopSlug/debug/fix-variant-ids
// This is a temporary endpoint for fixing variant IDs in existing products
func FixVariantIDs(c *gin.Context) {
	shopSlug := c.Param("shopSlug")

	// Lookup the shop by its slug
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Run the variant ID fix
	err = services.UpdateProductVariantIDs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fix variant IDs: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Variant IDs fixed successfully"})
}

// TestDiscounts handles GET /shops/:shopSlug/test-discounts
func TestDiscounts(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Get all discounts for this shop
	discounts, err := services.ListDiscountsByShopService(shop.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get customer ID if authenticated
	var customerID primitive.ObjectID
	var customerSegmentIDs []primitive.ObjectID
	if cidVal, exists := c.Get("user_id"); exists {
		if cidHex, ok := cidVal.(string); ok {
			if cid, err := primitive.ObjectIDFromHex(cidHex); err == nil {
				customerID = cid
				customerSegmentIDs, _ = services.GetCustomerSegmentIDs(shop.ID, customerID)
			}
		}
	}

	// Test each discount
	var results []map[string]interface{}
	for _, discount := range discounts {
		result := map[string]interface{}{
			"id":                  discount.ID.Hex(),
			"name":                discount.Name,
			"type":                discount.Type,
			"value":               discount.Value,
			"category":            discount.Category,
			"active":              discount.IsActive(),
			"shop_id":             discount.ShopID.Hex(),
			"applies_to_products": discount.AppliesToProducts,
			"applies_to_variants": discount.AppliesToVariants,
			"eligibility_type":    discount.EligibilityType,
			"allowed_customers":   discount.AllowedCustomerIDs,
			"allowed_segments":    discount.AllowedSegmentIDs,
			"usage_limit":         discount.UsageLimit,
			"per_customer_limit":  discount.PerCustomerLimit,
			"current_usage":       discount.CurrentUsage,
			"start_at":            discount.StartAt,
			"end_at":              discount.EndAt,
		}

		// Test customer eligibility
		if !customerID.IsZero() {
			canUse, err := services.CanCustomerUseDiscount(&discount, customerID, customerSegmentIDs)
			result["customer_can_use"] = canUse
			result["customer_error"] = err
		}

		// Test discount calculation
		testPrice := 100.0
		savings := discount.CalculateDiscount(testPrice)
		result["test_calculation"] = map[string]interface{}{
			"test_price":  testPrice,
			"savings":     savings,
			"final_price": testPrice - savings,
		}

		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"shop_id":           shop.ID.Hex(),
		"customer_id":       customerID.Hex(),
		"customer_segments": customerSegmentIDs,
		"total_discounts":   len(discounts),
		"discounts":         results,
	})
}

// PlaceOrder handles POST /shops/:shopSlug/orders
func PlaceOrder(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// Get customer ID from authentication
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
	customerID, err := primitive.ObjectIDFromHex(cidHex)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid customer ID"})
		return
	}

	// Link customer to shop if not already linked
	_, _, _ = services.LinkIfNotLinked(shop.ID, customerID)

	// Parse request - ONLY accept minimal data from frontend
	var req PlaceOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request - ensure we only have the minimal required data
	if len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no items in order"})
		return
	}

	// Get customer segments for discount validation
	customerSegmentIDs, err := services.GetCustomerSegmentIDs(shop.ID, customerID)
	if err != nil {
		customerSegmentIDs = []primitive.ObjectID{}
	}

	// Initialize order variables
	var orderItems []models.OrderItem
	var itemDiscountDetails []map[string]interface{}
	subtotal := 0.0
	totalDiscount := 0.0
	appliedDiscountIDsMap := make(map[primitive.ObjectID]struct{})

	// Process each item - ALL pricing calculated server-side
	for _, itemReq := range req.Items {
		// Validate product ID
		productID, err := primitive.ObjectIDFromHex(itemReq.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID: " + itemReq.ProductID})
			return
		}

		// Validate variant ID if provided
		variantID := primitive.NilObjectID
		if itemReq.VariantID != "" {
			variantID, err = primitive.ObjectIDFromHex(itemReq.VariantID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid variant ID: " + itemReq.VariantID})
				return
			}
		}

		// Validate quantity
		if itemReq.Quantity <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "quantity must be positive"})
			return
		}

		// Fetch product from database - server-side validation
		product, err := services.GetProductByIDService(itemReq.ProductID)
		if err != nil || product == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product not found: " + itemReq.ProductID})
			return
		}

		// Validate product belongs to this shop
		if product.ShopID != shop.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product not found in this shop"})
			return
		}

		// Determine unit price and product details - server-side calculation
		var unitPrice float64
		var productName string
		var productImage string
		var stock int
		var orderVariantID primitive.ObjectID

		// Check if product has variants
		hasRealVariants := len(product.Variants) > 0

		if hasRealVariants && !variantID.IsZero() {
			// Variant-specific pricing
			var foundVariant *models.Variant
			for i := range product.Variants {
				if product.Variants[i].VariantID == variantID {
					foundVariant = &product.Variants[i]
					break
				}
			}
			if foundVariant == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "variant not found: " + itemReq.VariantID + " for product: " + itemReq.ProductID})
				return
			}
			unitPrice = foundVariant.Price
			productName = product.Name
			if len(foundVariant.Options) > 0 {
				productName += " - "
				for i, opt := range foundVariant.Options {
					if i > 0 {
						productName += ", "
					}
					productName += opt.Name + ": " + opt.Value
				}
			}
			if foundVariant.Image != "" {
				productImage = foundVariant.Image
			} else {
				productImage = product.MainImage
			}
			stock = foundVariant.Stock
			orderVariantID = variantID
		} else {
			// Product-level pricing
			unitPrice = product.Price
			productName = product.Name
			productImage = product.MainImage
			stock = product.Stock
			orderVariantID = primitive.NilObjectID
		}

		// Validate stock availability
		if stock < itemReq.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient stock for product/variant: " + itemReq.ProductID})
			return
		}

		// Calculate line total server-side
		lineTotal := unitPrice * float64(itemReq.Quantity)
		subtotal += lineTotal

		// Get collection IDs for discount eligibility
		collectionIDs, err := services.GetCollectionIDsForProduct(productID)
		if err != nil {
			collectionIDs = []primitive.ObjectID{}
		}

		// Get active discounts for this product/variant - ALL discount logic server-side
		discounts, err := services.GetActiveDiscountsForProductService(shop.ID, productID, variantID, collectionIDs)
		if err != nil {
			fmt.Printf("Error getting discounts: %v\n", err)
			discounts = []models.Discount{}
		}

		// Find best eligible discount - server-side calculation only
		var bestDiscount *models.Discount
		bestSavings := 0.0

		fmt.Printf("Found %d discounts for product %s, variant %s\n", len(discounts), productID.Hex(), variantID.Hex())

		// Use the improved discount selection logic
		bestDiscount = services.GetBestEligibleDiscountForProduct(productID, variantID, customerID, customerSegmentIDs, discounts)

		if bestDiscount != nil {
			fmt.Printf("Selected best discount: %s (Type: %s, Value: %.2f)\n",
				bestDiscount.Name, bestDiscount.Type, bestDiscount.Value)

			// Calculate actual savings for this specific order
			bestSavings = bestDiscount.CalculateDiscountForQuantity(unitPrice, itemReq.Quantity)
			fmt.Printf("Discount savings: %.2f for unit price: %.2f, quantity: %d\n",
				bestSavings, unitPrice, itemReq.Quantity)
		} else {
			fmt.Printf("No eligible discount found for this product/variant\n")
		}

		// Apply discount if found - server-side calculation
		itemDiscountAmount := 0.0
		appliedDiscountIDs := []primitive.ObjectID{}
		if bestDiscount != nil && bestSavings > 0 {
			itemDiscountAmount = bestSavings
			appliedDiscountIDs = append(appliedDiscountIDs, bestDiscount.ID)
			appliedDiscountIDsMap[bestDiscount.ID] = struct{}{}
			fmt.Printf("Applied discount: %s, amount: %.2f\n", bestDiscount.Name, itemDiscountAmount)
		}

		totalDiscount += itemDiscountAmount
		finalLineTotal := lineTotal - itemDiscountAmount
		if finalLineTotal < 0 {
			finalLineTotal = 0
		}

		// Store discount details for response
		itemDiscountDetails = append(itemDiscountDetails, map[string]interface{}{
			"product_id":           productID.Hex(),
			"variant_id":           variantID.Hex(),
			"unit_price":           unitPrice,
			"quantity":             itemReq.Quantity,
			"line_total":           lineTotal,
			"discount":             itemDiscountAmount,
			"final_line_total":     finalLineTotal,
			"applied_discount_ids": appliedDiscountIDs,
		})

		// Create order item with server-calculated pricing
		orderItem := models.OrderItem{
			ProductID:  productID,
			VariantID:  orderVariantID,
			Name:       productName,
			Quantity:   itemReq.Quantity,
			UnitPrice:  unitPrice,
			TotalPrice: finalLineTotal, // Use server-calculated discounted price
			Image:      productImage,
		}
		orderItems = append(orderItems, orderItem)
	}

	// Calculate final totals server-side
	finalTotal := subtotal - totalDiscount
	if finalTotal < 0 {
		finalTotal = 0
	}

	// Prepare applied discount IDs
	appliedDiscountIDs := make([]primitive.ObjectID, 0, len(appliedDiscountIDsMap))
	for id := range appliedDiscountIDsMap {
		appliedDiscountIDs = append(appliedDiscountIDs, id)
	}

	// Create order with server-calculated totals
	order := &models.Order{
		ShopID:             shop.ID,
		CustomerID:         customerID,
		Items:              orderItems,
		Subtotal:           subtotal,
		DiscountTotal:      totalDiscount,
		Total:              finalTotal,
		Status:             "pending",
		AppliedDiscountIDs: appliedDiscountIDs,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	// Save order to database
	created, err := services.CreateOrderService(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Record discount usage
	for id := range appliedDiscountIDsMap {
		for _, item := range itemDiscountDetails {
			for _, appliedID := range item["applied_discount_ids"].([]primitive.ObjectID) {
				if appliedID == id {
					_ = services.RecordDiscountUsageAtomic(id.Hex(), customerID, item["line_total"].(float64))
				}
			}
		}
	}

	// Reduce stock for all items
	for _, itemReq := range req.Items {
		productID, _ := primitive.ObjectIDFromHex(itemReq.ProductID)
		if itemReq.VariantID != "" {
			variantID, _ := primitive.ObjectIDFromHex(itemReq.VariantID)
			if !variantID.IsZero() {
				err := services.ReduceVariantStock(productID, variantID, itemReq.Quantity)
				if err != nil {
					fmt.Printf("Warning: Failed to reduce stock for variant %s: %v\n", itemReq.VariantID, err)
				}
			} else {
				err := services.ReduceProductStock(productID, itemReq.Quantity)
				if err != nil {
					fmt.Printf("Warning: Failed to reduce stock for product %s: %v\n", itemReq.ProductID, err)
				}
			}
		} else {
			err := services.ReduceProductStock(productID, itemReq.Quantity)
			if err != nil {
				fmt.Printf("Warning: Failed to reduce stock for product %s: %v\n", itemReq.ProductID, err)
			}
		}
	}

	// Return order with server-calculated totals
	c.JSON(http.StatusCreated, gin.H{
		"id":                    created.ID.Hex(),
		"order":                 created,
		"item_discount_details": itemDiscountDetails,
		"server_calculated":     true, // Flag to indicate all pricing was calculated server-side
		"security_note":         "All pricing calculated securely on server",
	})
}

// ListShopOrders handles GET /shops/:shopSlug/orders
func ListShopOrders(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// List by shop ID
	list, err := services.ListOrdersByShopService(shop.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Filter to only the current customer, if logged in
	if cidVal, exists := c.Get("user_id"); exists {
		if cidHex, ok := cidVal.(string); ok {
			if cid, err := primitive.ObjectIDFromHex(cidHex); err == nil {
				filtered := make([]models.Order, 0, len(list))
				for _, o := range list {
					if o.CustomerID == cid {
						filtered = append(filtered, o)
					}
				}
				list = filtered
			}
		}
	}

	c.JSON(http.StatusOK, list)
}

// GetOrderDetail handles GET /shops/:shopSlug/orders/:orderId
func GetOrderDetail(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	orderID := c.Param("orderId")
	order, err := services.GetOrderByIDService(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if order == nil || order.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	// Only the placing customer may view it
	if cidVal, exists := c.Get("user_id"); exists {
		if cidHex, ok := cidVal.(string); ok {
			if cid, err := primitive.ObjectIDFromHex(cidHex); err == nil {
				if order.CustomerID != cid {
					c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to view this order"})
					return
				}
			}
		}
	}

	c.JSON(http.StatusOK, order)
}
