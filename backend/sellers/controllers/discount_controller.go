package controllers

import (
	"net/http"
	"strconv"
	"time"

	"math"

	"github.com/Endale2/DRPS/sellers/repositories"
	"github.com/Endale2/DRPS/shared/models"
	sharedSvc "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Input struct for binding JSON
type DiscountInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description,omitempty"`
	Category    string  `json:"category" binding:"required"`
	Type        string  `json:"type,omitempty"`
	Value       float64 `json:"value,omitempty"`

	AppliesToProducts []string `json:"appliesToProducts,omitempty"`
	AppliesToVariants []string `json:"appliesToVariants,omitempty"`

	StartAt string `json:"startAt" binding:"required"` // RFC3339
	EndAt   string `json:"endAt" binding:"required"`   // RFC3339
	Active  bool   `json:"active"`

	// Enhanced discount functionality
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
		Name:        in.Name,
		Description: in.Description,
		Category:    models.DiscountCategory(in.Category),
		Type:        models.DiscountType(in.Type),
		Value:       in.Value,

		ShopID:           shop.ID,
		SellerID:         sellerID,
		StartAt:          startAt,
		EndAt:            endAt,
		Active:           in.Active,
		EligibilityType:  models.DiscountEligibilityType(in.EligibilityType),
		UsageLimit:       in.UsageLimit,
		PerCustomerLimit: in.PerCustomerLimit,
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

	// Deduplicate AppliesToProducts and AppliesToVariants
	uniqueProducts := make(map[primitive.ObjectID]struct{})
	for _, oid := range d.AppliesToProducts {
		uniqueProducts[oid] = struct{}{}
	}
	var dedupedProducts []primitive.ObjectID
	for oid := range uniqueProducts {
		dedupedProducts = append(dedupedProducts, oid)
	}
	d.AppliesToProducts = dedupedProducts

	uniqueVariants := make(map[primitive.ObjectID]struct{})
	for _, oid := range d.AppliesToVariants {
		uniqueVariants[oid] = struct{}{}
	}
	var dedupedVariants []primitive.ObjectID
	for oid := range uniqueVariants {
		dedupedVariants = append(dedupedVariants, oid)
	}
	d.AppliesToVariants = dedupedVariants
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

	// Convert each discount to the same format as GetDiscount
	var response []gin.H
	for _, d := range list {
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
				"price":       p.Price,
				"stock":       p.Stock,
			})
		}

		// Convert ObjectID arrays to string arrays
		var allowedCustomers []string
		for _, id := range d.AllowedCustomerIDs {
			allowedCustomers = append(allowedCustomers, id.Hex())
		}
		var allowedSegments []string
		for _, id := range d.AllowedSegmentIDs {
			allowedSegments = append(allowedSegments, id.Hex())
		}
		var appliesToVariants []string
		for _, id := range d.AppliesToVariants {
			appliesToVariants = append(appliesToVariants, id.Hex())
		}

		// Build response for each discount
		resp := gin.H{
			"id":          d.ID.Hex(),
			"name":        d.Name,
			"description": d.Description,
			"category":    d.Category,
			"type":        d.Type,
			"value":       d.Value,

			"start_at":            d.StartAt,
			"end_at":              d.EndAt,
			"active":              d.Active,
			"applies_to_products": products,
			"applies_to_variants": appliesToVariants,
			"usage_limit":         d.UsageLimit,
			"per_customer_limit":  d.PerCustomerLimit,
			"current_usage":       d.CurrentUsage,
			"usage_tracking":      d.UsageTracking,
			"eligibility_type":    d.EligibilityType,
			"allowed_customers":   allowedCustomers,
			"allowed_segments":    allowedSegments,
			"created_at":          d.CreatedAt,
			"updated_at":          d.UpdatedAt,
		}
		response = append(response, resp)
	}

	c.JSON(http.StatusOK, response)
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
			"price":       p.Price,
			"stock":       p.Stock,
		})
	}

	// Convert ObjectID arrays to string arrays
	var allowedCustomers []string
	for _, id := range d.AllowedCustomerIDs {
		allowedCustomers = append(allowedCustomers, id.Hex())
	}
	var allowedSegments []string
	for _, id := range d.AllowedSegmentIDs {
		allowedSegments = append(allowedSegments, id.Hex())
	}
	var appliesToVariants []string
	for _, id := range d.AppliesToVariants {
		appliesToVariants = append(appliesToVariants, id.Hex())
	}

	// Build response
	resp := gin.H{
		"id":          d.ID.Hex(),
		"name":        d.Name,
		"description": d.Description,
		"category":    d.Category,
		"type":        d.Type,
		"value":       d.Value,

		"start_at":            d.StartAt,
		"end_at":              d.EndAt,
		"active":              d.Active,
		"applies_to_products": products,
		"applies_to_variants": appliesToVariants,
		"usage_limit":         d.UsageLimit,
		"per_customer_limit":  d.PerCustomerLimit,
		"current_usage":       d.CurrentUsage,
		"usage_tracking":      d.UsageTracking,
		"eligibility_type":    d.EligibilityType,
		"allowed_customers":   allowedCustomers,
		"allowed_segments":    allowedSegments,
		"created_at":          d.CreatedAt,
		"updated_at":          d.UpdatedAt,
	}

	c.JSON(http.StatusOK, resp)
}

// GetDiscountProductsPaginated GET /seller/shops/:shopId/discounts/:id/products
func GetDiscountProductsPaginated(c *gin.Context) {
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

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Calculate pagination for products
	total := int64(len(d.AppliesToProducts))
	skip := (page - 1) * limit
	end := skip + limit

	var products []gin.H
	if skip < int(total) {
		// Get the paginated slice of product IDs
		var productIDs []primitive.ObjectID
		if end > int(total) {
			productIDs = d.AppliesToProducts[skip:]
		} else {
			productIDs = d.AppliesToProducts[skip:end]
		}

		// Fetch product details for the paginated slice
		for _, pid := range productIDs {
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
				"price":       p.Price,
				"stock":       p.Stock,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"total":    total,
		"page":     page,
		"limit":    limit,
		"pages":    int(math.Ceil(float64(total) / float64(limit))),
	})
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
		case "name", "description", "couponCode", "value", "active", "category", "type", "eligibilityType":
			// map json keys to bson keys if needed: e.g. "couponCode"->"coupon_code"
			field := k
			if k == "couponCode" {
				field = "coupon_code"
			} else if k == "eligibilityType" {
				field = "eligibility_type"
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
		case "usageLimit":
			if num, ok := v.(float64); ok {
				limit := int(num)
				upd["usage_limit"] = &limit
			}
		case "perCustomerLimit":
			if num, ok := v.(float64); ok {
				limit := int(num)
				upd["per_customer_limit"] = &limit
			}
		case "allowedCustomers":
			arr, _ := v.([]interface{})
			var oids []primitive.ObjectID
			for _, e := range arr {
				if str, ok := e.(string); ok {
					if oid, err := primitive.ObjectIDFromHex(str); err == nil {
						oids = append(oids, oid)
					}
				}
			}
			upd["allowed_customers"] = oids
		case "allowedSegments":
			arr, _ := v.([]interface{})
			var oids []primitive.ObjectID
			for _, e := range arr {
				if str, ok := e.(string); ok {
					if oid, err := primitive.ObjectIDFromHex(str); err == nil {
						oids = append(oids, oid)
					}
				}
			}
			upd["allowed_segments"] = oids
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

	var input struct {
		CustomerIDs []string `json:"customerIds" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert string IDs to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, idStr := range input.CustomerIDs {
		if oid, err := primitive.ObjectIDFromHex(idStr); err == nil {
			objectIDs = append(objectIDs, oid)
		}
	}

	// Add to existing allowed customers
	d.AllowedCustomerIDs = append(d.AllowedCustomerIDs, objectIDs...)
	d.EligibilityType = models.DiscountEligibilitySpecific

	// Update in database
	err = sharedSvc.UpdateDiscountService(idHex, bson.M{
		"allowed_customers": d.AllowedCustomerIDs,
		"eligibility_type":  d.EligibilityType,
	})

	if err != nil {
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

	var input struct {
		SegmentIDs []string `json:"segmentIds" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert string IDs to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, idStr := range input.SegmentIDs {
		if oid, err := primitive.ObjectIDFromHex(idStr); err == nil {
			objectIDs = append(objectIDs, oid)
		}
	}

	// Add to existing allowed segments
	d.AllowedSegmentIDs = append(d.AllowedSegmentIDs, objectIDs...)
	d.EligibilityType = models.DiscountEligibilitySegment

	// Update in database
	err = sharedSvc.UpdateDiscountService(idHex, bson.M{
		"allowed_segments": d.AllowedSegmentIDs,
		"eligibility_type": d.EligibilityType,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "segments added to discount"})
}

// ClearAllowedCustomersAndSegments POST /seller/shops/:shopId/discounts/:id/clear-eligibility
func ClearAllowedCustomersAndSegments(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))
	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}
	idHex := c.Param("id")
	upd := bson.M{
		"eligibility_type": models.DiscountEligibilityAll,
		"updated_at":       time.Now(),
	}
	if err := sharedSvc.UpdateDiscountService(idHex, upd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "eligibility set to allow everyone"})
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

	stats, err := sharedSvc.GetDiscountUsageStats(idHex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
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

	var input struct {
		CustomerID string `json:"customerId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerID, err := primitive.ObjectIDFromHex(input.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	// Get customer segments
	customerSegmentIDs, err := sharedSvc.GetCustomerSegmentIDs(shop.ID, customerID)
	if err != nil {
		customerSegmentIDs = []primitive.ObjectID{}
	}

	// Validate discount for customer
	canUse, err := sharedSvc.CanCustomerUseDiscount(d, customerID, customerSegmentIDs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"can_use": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"can_use": canUse,
		"error":   nil,
	})
}

// AddMixedEligibility POST /seller/shops/:shopId/discounts/:id/mixed-eligibility
func AddMixedEligibility(c *gin.Context) {
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

	var input struct {
		CustomerIDs []string `json:"customerIds,omitempty"`
		SegmentIDs  []string `json:"segmentIds,omitempty"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert string IDs to ObjectIDs
	var customerObjectIDs []primitive.ObjectID
	for _, idStr := range input.CustomerIDs {
		if oid, err := primitive.ObjectIDFromHex(idStr); err == nil {
			customerObjectIDs = append(customerObjectIDs, oid)
		}
	}

	var segmentObjectIDs []primitive.ObjectID
	for _, idStr := range input.SegmentIDs {
		if oid, err := primitive.ObjectIDFromHex(idStr); err == nil {
			segmentObjectIDs = append(segmentObjectIDs, oid)
		}
	}

	// Determine eligibility type based on what's provided
	var eligibilityType models.DiscountEligibilityType
	if len(customerObjectIDs) > 0 && len(segmentObjectIDs) > 0 {
		// Mixed eligibility - both customers and segments
		eligibilityType = models.DiscountEligibilitySpecific // We'll use specific for mixed
	} else if len(customerObjectIDs) > 0 {
		// Only customers
		eligibilityType = models.DiscountEligibilitySpecific
	} else if len(segmentObjectIDs) > 0 {
		// Only segments
		eligibilityType = models.DiscountEligibilitySegment
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "must provide at least one customer or segment ID"})
		return
	}

	// Add to existing allowed customers and segments
	d.AllowedCustomerIDs = append(d.AllowedCustomerIDs, customerObjectIDs...)
	d.AllowedSegmentIDs = append(d.AllowedSegmentIDs, segmentObjectIDs...)
	d.EligibilityType = eligibilityType

	// Update in database
	err = sharedSvc.UpdateDiscountService(idHex, bson.M{
		"allowed_customers": d.AllowedCustomerIDs,
		"allowed_segments":  d.AllowedSegmentIDs,
		"eligibility_type":  d.EligibilityType,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "eligibility updated successfully"})
}

// GetEligibleCustomers GET /seller/shops/:shopId/discounts/:id/eligible-customers
func GetEligibleCustomers(c *gin.Context) {
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

	// If eligibility is 'all', return empty array (everyone is eligible)
	if d.EligibilityType == models.DiscountEligibilityAll {
		c.JSON(http.StatusOK, gin.H{
			"eligible_customers": []gin.H{},
			"eligibility_type":   "all",
			"message":            "Everyone is eligible for this discount",
		})
		return
	}

	// Get direct customers
	var directCustomers []gin.H
	if len(d.AllowedCustomerIDs) > 0 {
		for _, customerID := range d.AllowedCustomerIDs {
			customer, err := sharedSvc.GetCustomerByIDService(customerID.Hex())
			if err == nil && customer != nil {
				directCustomers = append(directCustomers, gin.H{
					"id":        customer.ID.Hex(),
					"firstName": customer.FirstName,
					"lastName":  customer.LastName,
					"email":     customer.Email,
					"source":    "direct",
				})
			}
		}
	}

	// Get customers from segments
	var segmentCustomers []gin.H
	if len(d.AllowedSegmentIDs) > 0 {
		for _, segmentID := range d.AllowedSegmentIDs {
			segment, err := repositories.GetCustomerSegmentByID(segmentID)
			if err == nil && segment != nil {
				// Get customers in this segment
				customerIDs, err := repositories.GetCustomersInSegment(segmentID)
				if err == nil {
					for _, customerID := range customerIDs {
						// Get customer details
						customer, err := sharedSvc.GetCustomerByIDService(customerID.Hex())
						if err == nil && customer != nil {
							// Check if customer is already in direct customers
							isDirect := false
							for _, direct := range directCustomers {
								if direct["id"] == customer.ID.Hex() {
									isDirect = true
									break
								}
							}

							// Only add if not already in direct customers
							if !isDirect {
								segmentCustomers = append(segmentCustomers, gin.H{
									"id":          customer.ID.Hex(),
									"firstName":   customer.FirstName,
									"lastName":    customer.LastName,
									"email":       customer.Email,
									"source":      "segment",
									"segmentName": segment.Name,
								})
							}
						}
					}
				}
			}
		}
	}

	// Combine and deduplicate customers
	allCustomers := append(directCustomers, segmentCustomers...)

	c.JSON(http.StatusOK, gin.H{
		"eligible_customers": allCustomers,
		"eligibility_type":   d.EligibilityType,
		"direct_count":       len(directCustomers),
		"segment_count":      len(segmentCustomers),
		"total_count":        len(allCustomers),
	})
}
