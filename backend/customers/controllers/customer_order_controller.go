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
				// Link customer to shop if not already linked
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

	// 4) Validate request
	if len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order must contain at least one item"})
		return
	}

	// 5) Build order items with product data from database
	var orderItems []models.OrderItem
	var subtotal float64
	var totalDiscount float64

	for _, itemReq := range req.Items {
		// Parse product ID
		productID, err := primitive.ObjectIDFromHex(itemReq.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID: " + itemReq.ProductID})
			return
		}

		// Parse variant ID if provided
		var variantID primitive.ObjectID
		if itemReq.VariantID != "" {
			variantID, err = primitive.ObjectIDFromHex(itemReq.VariantID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid variant ID: " + itemReq.VariantID})
				return
			}
			fmt.Printf("Requested variant ID: %s\n", variantID.Hex())
		}

		// Fetch product with discounts applied
		product, err := services.GetProductByIDWithDiscountsService(productID.Hex(), &customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product: " + err.Error()})
			return
		}
		if product == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product not found: " + itemReq.ProductID})
			return
		}

		// Debug: Log product and variant information
		fmt.Printf("Product %s: Name=%s, Price=%.2f, Stock=%d\n",
			productID.Hex(), product.Name, product.Price, product.Stock)
		fmt.Printf("Product %s has %d variants:\n", productID.Hex(), len(product.Variants))
		for i, v := range product.Variants {
			fmt.Printf("  Variant %d: ID=%s, Options=%+v, Price=%.2f, Stock=%d\n", i, v.VariantID.Hex(), v.Options, v.Price, v.Stock)
		}

		// Verify product belongs to this shop
		if product.ShopID != shop.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product does not belong to this shop"})
			return
		}

		// Determine price and product details
		var unitPrice float64
		var productName string
		var productImage string
		var orderVariantID primitive.ObjectID

		// Check if product has real variants
		hasRealVariants := false
		for _, v := range product.Variants {
			if len(v.Options) > 0 && (v.Options[0].Name != "" || v.Options[0].Value != "") {
				hasRealVariants = true
				break
			}
		}

		if hasRealVariants && !variantID.IsZero() {
			// Product has real variants and variant was specified
			var foundVariant *models.Variant
			availableVariantIDs := make([]string, 0, len(product.Variants))

			for i := range product.Variants {
				availableVariantIDs = append(availableVariantIDs, product.Variants[i].VariantID.Hex())
				// Check VariantID (this should match what's in the database)
				if product.Variants[i].VariantID == variantID {
					foundVariant = &product.Variants[i]
					break
				}
			}
			if foundVariant == nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":              "variant not found: " + itemReq.VariantID + " for product: " + itemReq.ProductID,
					"available_variants": availableVariantIDs,
					"requested_variant":  variantID.Hex(),
				})
				return
			}

			// Use regular price if available, otherwise use regular price
			unitPrice = foundVariant.Price

			productName = product.Name
			if len(foundVariant.Options) > 0 {
				// Add variant options to product name
				productName += " - "
				for i, opt := range foundVariant.Options {
					if i > 0 {
						productName += ", "
					}
					productName += opt.Name + ": " + opt.Value
				}
			}

			// Use variant image if available, otherwise product main image
			if foundVariant.Image != "" {
				productImage = foundVariant.Image
			} else {
				productImage = product.MainImage
			}

			// Check stock
			if foundVariant.Stock < itemReq.Quantity {
				c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient stock for variant: " + itemReq.VariantID})
				return
			}

			orderVariantID = variantID
		} else {
			// Product has no real variants or no variant specified - use product base price
			unitPrice = product.Price

			fmt.Printf("Using product base price: %.2f (Product.Price: %.2f)\n",
				unitPrice, product.Price)

			productName = product.Name
			productImage = product.MainImage

			// Check stock
			if product.Stock < itemReq.Quantity {
				c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient stock for product: " + itemReq.ProductID})
				return
			}

			// For products without variants, don't set a variant ID
			orderVariantID = primitive.NilObjectID
		}

		// Calculate line total
		lineTotal := unitPrice * float64(itemReq.Quantity)
		subtotal += lineTotal

		fmt.Printf("Creating order item: Product=%s, VariantID=%s, UnitPrice=%.2f, TotalPrice=%.2f\n",
			productID.Hex(), orderVariantID.Hex(), unitPrice, lineTotal)

		// Create order item
		orderItem := models.OrderItem{
			ProductID:  productID,
			VariantID:  orderVariantID, // Will be zero ObjectID for products without variants
			Name:       productName,
			Quantity:   itemReq.Quantity,
			UnitPrice:  unitPrice,
			TotalPrice: lineTotal,
			Image:      productImage,
		}

		orderItems = append(orderItems, orderItem)
	}

	// 7) Calculate final total
	finalTotal := subtotal - totalDiscount
	if finalTotal < 0 {
		finalTotal = 0
	}

	// 8) Create the order
	order := &models.Order{
		ShopID:        shop.ID,
		CustomerID:    customerID,
		Items:         orderItems,
		Subtotal:      subtotal,
		DiscountTotal: totalDiscount,
		Total:         finalTotal,
		Status:        "pending", // Default status
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// 9) Save the order
	created, err := services.CreateOrderService(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Record usage for item-level discounts (from products with automatic discounts)
	for _, itemReq := range req.Items {
		productID, _ := primitive.ObjectIDFromHex(itemReq.ProductID)
		var variantID primitive.ObjectID
		if itemReq.VariantID != "" {
			variantID, _ = primitive.ObjectIDFromHex(itemReq.VariantID)
		}

		// Get collection IDs for this product
		collectionIDs, err := services.GetCollectionIDsForProduct(productID)
		if err != nil {
			collectionIDs = []primitive.ObjectID{}
		}

		// Get active discounts for this product/variant
		discounts, err := services.GetActiveDiscountsForProductService(shop.ID, productID, variantID, collectionIDs)
		if err == nil && len(discounts) > 0 {
			// Find the best discount that was applied
			var bestDiscount *models.Discount
			bestSavings := 0.0

			for i := range discounts {
				discount := &discounts[i]
				if discount.IsActive() {
					// Use the improved validation function
					if canUse, err := services.CanCustomerUseDiscount(discount, customerID, []primitive.ObjectID{}); err == nil && canUse {
						// Calculate what the savings would be
						product, _ := services.GetProductByIDService(productID.Hex())
						if product != nil {
							var price float64
							if !variantID.IsZero() {
								for _, v := range product.Variants {
									if v.VariantID == variantID {
										price = v.Price
										break
									}
								}
							} else {
								price = product.Price
							}
							lineTotal := price * float64(itemReq.Quantity)
							savings := discount.CalculateDiscount(lineTotal)
							if savings > bestSavings {
								bestSavings = savings
								bestDiscount = discount
							}
						}
					}
				}
			}

			// Record usage for the best discount if one was found
			if bestDiscount != nil {
				// Calculate line total for usage recording
				product, _ := services.GetProductByIDService(productID.Hex())
				if product != nil {
					var price float64
					if !variantID.IsZero() {
						for _, v := range product.Variants {
							if v.VariantID == variantID {
								price = v.Price
								break
							}
						}
					} else {
						price = product.Price
					}
					lineTotal := price * float64(itemReq.Quantity)

					// Use atomic function to safely record usage
					err = services.RecordDiscountUsageAtomic(bestDiscount.ID.Hex(), customerID, lineTotal)
					if err != nil {
						fmt.Printf("Warning: Failed to record item discount usage: %v\n", err)
					}
				}
			}
		}
	}

	// 10) Reduce stock for all items
	for _, itemReq := range req.Items {
		productID, _ := primitive.ObjectIDFromHex(itemReq.ProductID)

		if itemReq.VariantID != "" {
			variantID, _ := primitive.ObjectIDFromHex(itemReq.VariantID)
			// Only reduce variant stock if it's a real variant (not zero ObjectID)
			if !variantID.IsZero() {
				err := services.ReduceVariantStock(productID, variantID, itemReq.Quantity)
				if err != nil {
					fmt.Printf("Warning: Failed to reduce stock for variant %s: %v\n", itemReq.VariantID, err)
				}
			} else {
				// Zero ObjectID means no variant - reduce product stock
				err := services.ReduceProductStock(productID, itemReq.Quantity)
				if err != nil {
					fmt.Printf("Warning: Failed to reduce stock for product %s: %v\n", itemReq.ProductID, err)
				}
			}
		} else {
			// No variant specified - reduce product stock
			err := services.ReduceProductStock(productID, itemReq.Quantity)
			if err != nil {
				fmt.Printf("Warning: Failed to reduce stock for product %s: %v\n", itemReq.ProductID, err)
			}
		}
	}

	// 11) Return the created order
	c.JSON(http.StatusCreated, created)
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
