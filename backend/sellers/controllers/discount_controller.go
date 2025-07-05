package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	sharedSvc "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Input struct for binding JSON
type DiscountInput struct {
	Name                        string   `json:"name" binding:"required"`
	Description                 string   `json:"description,omitempty"`
	Category                    string   `json:"category" binding:"required"`
	Type                        string   `json:"type,omitempty"`
	Value                       float64  `json:"value,omitempty"`
	CouponCode                  string   `json:"couponCode,omitempty"`
	AppliesToProducts           []string `json:"appliesToProducts,omitempty"`
	AppliesToVariants           []string `json:"appliesToVariants,omitempty"`
	AppliesToCollections        []string `json:"appliesToCollections,omitempty"`
	FreeShipping                bool     `json:"freeShipping,omitempty"`
	MinimumOrderSubtotal        *float64 `json:"minimumOrderSubtotal,omitempty"`
	MinimumOrderForFreeShipping *float64 `json:"minimumOrderForFreeShipping,omitempty"`
	StartAt                     string   `json:"startAt" binding:"required"` // ISO8601
	EndAt                       string   `json:"endAt" binding:"required"`
	Active                      bool     `json:"active"`
	// New fields for enhanced discount functionality
	EligibilityType  string   `json:"eligibilityType,omitempty"`
	AllowedCustomers []string `json:"allowedCustomers,omitempty"`
	AllowedSegments  []string `json:"allowedSegments,omitempty"`
	UsageLimit       *int     `json:"usageLimit,omitempty"`
	PerCustomerLimit *int     `json:"perCustomerLimit,omitempty"`
}

// helper to parse time
func parseTimeField(value string) (time.Time, error) {
	// parse RFC3339
	return time.Parse(time.RFC3339, value)
}

// CreateDiscount POST /seller/shops/:shopId/discounts
func CreateDiscount(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	var in DiscountInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	startAt, err := parseTimeField(in.StartAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid startAt format, must be RFC3339"})
		return
	}
	endAt, err := parseTimeField(in.EndAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid endAt format, must be RFC3339"})
		return
	}
	// Build Discount model
	d := &models.Discount{
		Name:                        in.Name,
		Description:                 in.Description,
		Category:                    models.DiscountCategory(in.Category),
		Type:                        models.DiscountType(in.Type),
		Value:                       in.Value,
		CouponCode:                  in.CouponCode,
		FreeShipping:                in.FreeShipping,
		MinimumOrderSubtotal:        in.MinimumOrderSubtotal,
		MinimumOrderForFreeShipping: in.MinimumOrderForFreeShipping,
		ShopID:                      shop.ID,
		SellerID:                    sellerID,
		StartAt:                     startAt,
		EndAt:                       endAt,
		Active:                      in.Active,
		EligibilityType:             models.DiscountEligibilityType(in.EligibilityType),
		UsageLimit:                  in.UsageLimit,
		PerCustomerLimit:            in.PerCustomerLimit,
	}
	// parse arrays
	for _, pid := range in.AppliesToProducts {
		if oid, err := primitive.ObjectIDFromHex(pid); err == nil {
			d.AppliesToProducts = append(d.AppliesToProducts, oid)
		}
	}
	for _, vid := range in.AppliesToVariants {
		if oid, err := primitive.ObjectIDFromHex(vid); err == nil {
			d.AppliesToVariants = append(d.AppliesToVariants, oid)
		}
	}
	for _, cid := range in.AppliesToCollections {
		if oid, err := primitive.ObjectIDFromHex(cid); err == nil {
			d.AppliesToCollections = append(d.AppliesToCollections, oid)
		}
	}
	for _, cid := range in.AllowedCustomers {
		if oid, err := primitive.ObjectIDFromHex(cid); err == nil {
			d.AllowedCustomerIDs = append(d.AllowedCustomerIDs, oid)
		}
	}
	for _, sid := range in.AllowedSegments {
		if oid, err := primitive.ObjectIDFromHex(sid); err == nil {
			d.AllowedSegmentIDs = append(d.AllowedSegmentIDs, oid)
		}
	}
	// Call service
	created, err := sharedSvc.CreateDiscountService(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// ListDiscounts GET /seller/shops/:shopId/discounts
func ListDiscounts(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}
	list, err := sharedSvc.ListDiscountsByShopService(shopHex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// GetDiscount GET /seller/shops/:shopId/discounts/:id
func GetDiscount(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	idHex := c.Param("id")
	d, err := sharedSvc.GetDiscountByIDService(idHex)
	if err != nil {
		if err == sharedSvc.ErrDiscountNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	if d.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}

	// Fetch product summaries for each applied product
	var products []gin.H
	for _, pid := range d.AppliesToProducts {
		p, err := sharedSvc.GetProductByIDService(pid.Hex())
		if err != nil {
			// skip products we can't load
			continue
		}
		if p == nil {
			continue
		}
		products = append(products, gin.H{
			"id":          p.ID.Hex(),
			"name":        p.Name,
			"main_image":  p.MainImage,
			"description": p.Description,
		})
	}

	// Build response
	resp := gin.H{
		"id":                     d.ID.Hex(),
		"name":                   d.Name,
		"description":            d.Description,
		"category":               d.Category,
		"type":                   d.Type,
		"value":                  d.Value,
		"coupon_code":            d.CouponCode,
		"free_shipping":          d.FreeShipping,
		"minimum_free_shipping":  d.MinimumOrderForFreeShipping,
		"minimum_order_subtotal": d.MinimumOrderSubtotal,
		"start_at":               d.StartAt,
		"end_at":                 d.EndAt,
		"active":                 d.Active,
		"applies_to_products":    products,
		"applies_to_variants":    d.AppliesToVariants,
		"applies_to_collections": d.AppliesToCollections,
		"usage_limit":            d.UsageLimit,
		"per_customer_limit":     d.PerCustomerLimit,
		"current_usage":          d.CurrentUsage,
		"usage_tracking":         d.UsageTracking,
		"eligibility_type":       d.EligibilityType,
		"allowed_customers":      d.AllowedCustomerIDs,
		"allowed_segments":       d.AllowedSegmentIDs,
		"buy_product_ids":        d.BuyProductIDs,
		"buy_quantity":           d.BuyQuantity,
		"get_product_ids":        d.GetProductIDs,
		"get_quantity":           d.GetQuantity,
		"created_at":             d.CreatedAt,
		"updated_at":             d.UpdatedAt,
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateDiscount PATCH /seller/shops/:shopId/discounts/:id
func UpdateDiscount(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}
	idHex := c.Param("id")
	// Option A: bind into DiscountInput and build upd map; or Option B: bind raw map[string]interface{}
	var in map[string]interface{}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	upd := bson.M{}
	for k, v := range in {
		switch k {
		case "name", "description", "couponCode", "freeShipping", "value", "active":
			// map json keys to bson keys if needed: e.g. "couponCode"->"coupon_code"
			field := k
			if k == "couponCode" {
				field = "coupon_code"
			}
			upd[field] = v
		case "startAt":
			if s, ok := v.(string); ok {
				if t, err := parseTimeField(s); err == nil {
					upd["start_at"] = t
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "invalid startAt format"})
					return
				}
			}
		case "endAt":
			if s, ok := v.(string); ok {
				if t, err := parseTimeField(s); err == nil {
					upd["end_at"] = t
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "invalid endAt format"})
					return
				}
			}
		case "appliesToProducts":
			arr, _ := v.([]interface{})
			var oids []primitive.ObjectID
			for _, e := range arr {
				if str, ok := e.(string); ok {
					if oid, err := primitive.ObjectIDFromHex(str); err == nil {
						oids = append(oids, oid)
					}
				}
			}
			upd["applies_to_products"] = oids
		case "appliesToVariants":
			arr, _ := v.([]interface{})
			var oids []primitive.ObjectID
			for _, e := range arr {
				if str, ok := e.(string); ok {
					if oid, err := primitive.ObjectIDFromHex(str); err == nil {
						oids = append(oids, oid)
					}
				}
			}
			upd["applies_to_variants"] = oids
		case "appliesToCollections":
			arr, _ := v.([]interface{})
			var oids []primitive.ObjectID
			for _, e := range arr {
				if str, ok := e.(string); ok {
					if oid, err := primitive.ObjectIDFromHex(str); err == nil {
						oids = append(oids, oid)
					}
				}
			}
			upd["applies_to_collections"] = oids
		case "minimumOrderSubtotal":
			if num, ok := v.(float64); ok {
				upd["minimum_order_subtotal"] = num
			}
		case "minimumOrderForFreeShipping":
			if num, ok := v.(float64); ok {
				upd["minimum_free_shipping"] = num
			}
		// ignore other fields (shopID, sellerID, createdAt, etc.)
		default:
			// skip unknown/immutable keys
		}
	}
	if err := sharedSvc.UpdateDiscountService(idHex, upd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// DeleteDiscount DELETE /seller/shops/:shopId/discounts/:id
func DeleteDiscount(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}
	idHex := c.Param("id")
	if err := sharedSvc.DeleteDiscountService(idHex); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// AddCustomersToDiscount POST /seller/shops/:shopId/discounts/:id/customers
func AddCustomersToDiscount(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	var req struct {
		CustomerIDs []string `json:"customerIds" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idHex := c.Param("id")
	if err := sharedSvc.AddCustomersToDiscount(idHex, req.CustomerIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customers added to discount"})
}

// AddSegmentsToDiscount POST /seller/shops/:shopId/discounts/:id/segments
func AddSegmentsToDiscount(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	var req struct {
		SegmentIDs []string `json:"segmentIds" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idHex := c.Param("id")
	if err := sharedSvc.AddSegmentsToDiscount(idHex, req.SegmentIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "segments added to discount"})
}

// GetDiscountUsageStats GET /seller/shops/:shopId/discounts/:id/usage
func GetDiscountUsageStats(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	idHex := c.Param("id")
	discount, err := sharedSvc.GetDiscountUsageStats(idHex)
	if err != nil {
		if err == sharedSvc.ErrDiscountNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if discount.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}

	// Build usage statistics response
	resp := gin.H{
		"discount_id":             discount.ID.Hex(),
		"name":                    discount.Name,
		"current_usage":           discount.CurrentUsage,
		"usage_limit":             discount.UsageLimit,
		"per_customer_limit":      discount.PerCustomerLimit,
		"usage_tracking":          discount.UsageTracking,
		"eligibility_type":        discount.EligibilityType,
		"allowed_customers_count": len(discount.AllowedCustomerIDs),
		"allowed_segments_count":  len(discount.AllowedSegmentIDs),
	}

	c.JSON(http.StatusOK, resp)
}

// ValidateDiscountForCustomer POST /seller/shops/:shopId/discounts/:id/validate
func ValidateDiscountForCustomer(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	var req struct {
		CustomerID         string   `json:"customerId" binding:"required"`
		CustomerSegmentIDs []string `json:"customerSegmentIds,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idHex := c.Param("id")
	discount, err := sharedSvc.GetDiscountByIDService(idHex)
	if err != nil {
		if err == sharedSvc.ErrDiscountNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if discount.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}

	customerID, err := primitive.ObjectIDFromHex(req.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	var customerSegmentIDs []primitive.ObjectID
	for _, segmentID := range req.CustomerSegmentIDs {
		if oid, err := primitive.ObjectIDFromHex(segmentID); err == nil {
			customerSegmentIDs = append(customerSegmentIDs, oid)
		}
	}

	err = sharedSvc.ValidateDiscountForCustomer(discount, customerID, customerSegmentIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"valid": false,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":   true,
		"message": "discount is valid for this customer",
	})
}
