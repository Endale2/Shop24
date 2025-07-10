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

// PlaceOrder handles POST /shops/:shopSlug/orders
func PlaceOrder(c *gin.Context) {
	// 1) Lookup the shop by its slug
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

	// 2) Get customer ID from authentication
	var customerID primitive.ObjectID
	if cidVal, exists := c.Get("user_id"); exists {
		if cidHex, ok := cidVal.(string); ok {
			if cid, err := primitive.ObjectIDFromHex(cidHex); err == nil {
				customerID = cid
				_, _, _ = services.LinkIfNotLinked(shop.ID, customerID)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid customer ID"})
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id format"})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	// 3) Bind the incoming order request
	var req PlaceOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order must contain at least one item"})
		return
	}

	var orderItems []models.OrderItem
	var subtotal float64
	var totalDiscount float64
	appliedDiscountIDsMap := map[primitive.ObjectID]struct{}{}
	itemDiscountDetails := make([]map[string]interface{}, 0)

	customerSegmentIDs, _ := services.GetCustomerSegmentIDs(shop.ID, customerID)

	for _, itemReq := range req.Items {
		productID, err := primitive.ObjectIDFromHex(itemReq.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID: " + itemReq.ProductID})
			return
		}
		var variantID primitive.ObjectID
		if itemReq.VariantID != "" {
			variantID, err = primitive.ObjectIDFromHex(itemReq.VariantID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid variant ID: " + itemReq.VariantID})
				return
			}
		}
		product, err := services.GetProductByIDService(productID.Hex())
		if err != nil || product == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product: " + itemReq.ProductID})
			return
		}
		if product.ShopID != shop.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product does not belong to this shop"})
			return
		}
		var unitPrice float64
		var productName string
		var productImage string
		var orderVariantID primitive.ObjectID
		var stock int
		hasRealVariants := false
		for _, v := range product.Variants {
			if len(v.Options) > 0 && (v.Options[0].Name != "" || v.Options[0].Value != "") {
				hasRealVariants = true
				break
			}
		}
		if hasRealVariants && !variantID.IsZero() {
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
			unitPrice = product.Price
			productName = product.Name
			productImage = product.MainImage
			stock = product.Stock
			orderVariantID = primitive.NilObjectID
		}
		if stock < itemReq.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient stock for product/variant: " + itemReq.ProductID})
			return
		}
		lineTotal := unitPrice * float64(itemReq.Quantity)
		subtotal += lineTotal
		collectionIDs, err := services.GetCollectionIDsForProduct(productID)
		if err != nil {
			collectionIDs = []primitive.ObjectID{}
		}
		discounts, err := services.GetActiveDiscountsForProductService(shop.ID, productID, variantID, collectionIDs)
		if err != nil {
			discounts = []models.Discount{}
		}
		var bestDiscount *models.Discount
		bestSavings := 0.0
		for i := range discounts {
			discount := &discounts[i]
			if discount.IsActive() {
				if canUse, err := services.CanCustomerUseDiscount(discount, customerID, customerSegmentIDs); err == nil && canUse {
					savings := discount.CalculateDiscount(lineTotal)
					if savings > bestSavings {
						bestSavings = savings
						bestDiscount = discount
					}
				}
			}
		}
		itemDiscountAmount := 0.0
		appliedDiscountIDs := []primitive.ObjectID{}
		if bestDiscount != nil && bestSavings > 0 {
			itemDiscountAmount = bestSavings
			appliedDiscountIDs = append(appliedDiscountIDs, bestDiscount.ID)
			appliedDiscountIDsMap[bestDiscount.ID] = struct{}{}
		}
		totalDiscount += itemDiscountAmount
		finalLineTotal := lineTotal - itemDiscountAmount
		if finalLineTotal < 0 {
			finalLineTotal = 0
		}
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
		orderItem := models.OrderItem{
			ProductID:  productID,
			VariantID:  orderVariantID,
			Name:       productName,
			Quantity:   itemReq.Quantity,
			UnitPrice:  unitPrice,
			TotalPrice: finalLineTotal, // Use discounted price for order item
			Image:      productImage,
		}
		orderItems = append(orderItems, orderItem)
	}
	finalTotal := subtotal - totalDiscount
	if finalTotal < 0 {
		finalTotal = 0
	}
	appliedDiscountIDs := make([]primitive.ObjectID, 0, len(appliedDiscountIDsMap))
	for id := range appliedDiscountIDsMap {
		appliedDiscountIDs = append(appliedDiscountIDs, id)
	}
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
	created, err := services.CreateOrderService(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for id := range appliedDiscountIDsMap {
		for _, item := range itemDiscountDetails {
			for _, appliedID := range item["applied_discount_ids"].([]primitive.ObjectID) {
				if appliedID == id {
					_ = services.RecordDiscountUsageAtomic(id.Hex(), customerID, item["line_total"].(float64))
				}
			}
		}
	}
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
	c.JSON(http.StatusCreated, gin.H{"id": created.ID.Hex(), "order": created, "item_discount_details": itemDiscountDetails})
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
