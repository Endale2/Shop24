// seller/controllers/seller_discount_controller.go
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

// DiscountInput for binding JSON in create/update
type DiscountInput struct {
	Name                        string                  `json:"name" binding:"required"`
	Description                 string                  `json:"description"`
	Category                    models.DiscountCategory `json:"category" binding:"required"`
	Type                        models.DiscountType     `json:"type"`
	Value                       float64                 `json:"value"`
	CouponCode                  string                  `json:"couponCode"`
	FreeShipping                bool                    `json:"freeShipping"`
	MinimumOrderSubtotal        *float64                `json:"minimumOrderSubtotal"`
	MinimumOrderForFreeShipping *float64                `json:"minimumOrderForFreeShipping"`
	AppliesToProducts           []string                `json:"appliesToProducts"`
	AppliesToVariants           []string                `json:"appliesToVariants"`
	StartAt                     time.Time               `json:"startAt" binding:"required"`
	EndAt                       time.Time               `json:"endAt" binding:"required"`
	Active                      bool                    `json:"active"`
}

// CreateDiscount POST /seller/shops/:shopId/discounts
func CreateDiscount(c *gin.Context) {
	// Auth & shop ownership
	userHex, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(userHex.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user context"})
		return
	}
	shopIDhex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopIDhex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch shop"})
		return
	}
	if shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	var in DiscountInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate times
	if in.EndAt.Before(in.StartAt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endAt must be after startAt"})
		return
	}
	// Build Discount model
	d := &models.Discount{
		Name:                        in.Name,
		Description:                 in.Description,
		Category:                    in.Category,
		Type:                        in.Type,
		Value:                       in.Value,
		CouponCode:                  in.CouponCode,
		FreeShipping:                in.FreeShipping,
		MinimumOrderSubtotal:        in.MinimumOrderSubtotal,
		MinimumOrderForFreeShipping: in.MinimumOrderForFreeShipping,
		ShopID:                      shop.ID,
		SellerID:                    sellerID,
		StartAt:                     in.StartAt,
		EndAt:                       in.EndAt,
		Active:                      in.Active,
	}
	// Parse product/variant IDs
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
	created, err := sharedSvc.CreateDiscountService(d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// ListDiscounts GET /seller/shops/:shopId/discounts
func ListDiscounts(c *gin.Context) {
	userHex, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(userHex.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user context"})
		return
	}
	shopIDhex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopIDhex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch shop"})
		return
	}
	if shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}
	list, err := sharedSvc.ListDiscountsByShopService(shopIDhex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// GetDiscount GET /seller/shops/:shopId/discounts/:id
func GetDiscount(c *gin.Context) {
	userHex, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(userHex.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user context"})
		return
	}
	shopIDhex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopIDhex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch shop"})
		return
	}
	if shop == nil || shop.OwnerID != sellerID {
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
	c.JSON(http.StatusOK, d)
}

// UpdateDiscount PATCH /seller/shops/:shopId/discounts/:id
func UpdateDiscount(c *gin.Context) {
	userHex, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(userHex.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user context"})
		return
	}
	shopIDhex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopIDhex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch shop"})
		return
	}
	if shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}
	idHex := c.Param("id")
	existing, err := sharedSvc.GetDiscountByIDService(idHex)
	if err != nil {
		if err == sharedSvc.ErrDiscountNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	if existing.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}
	// Bind update fields into a map
	var upd bson.M
	if err := c.ShouldBindJSON(&upd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Convert camelCase JSON keys to snake_case if needed:
	// e.g., client might send "startAt", but repository expects "start_at". 
	// For simplicity, assume client uses snake_case matching BSON tags.
	// Remove immutable:
	delete(upd, "id")
	delete(upd, "_id")
	delete(upd, "shop_id")
	delete(upd, "seller_id")
	// Validate time fields if present, service layer also checks.
	if startRaw, ok := upd["start_at"]; ok {
		if startTime, ok2 := startRaw.(time.Time); ok2 {
			if endRaw, exists := upd["end_at"]; exists {
				if endTime, ok3 := endRaw.(time.Time); ok3 && endTime.Before(startTime) {
					c.JSON(http.StatusBadRequest, gin.H{"error": "endAt must be after startAt"})
					return
				}
			}
		}
	}
	if err := sharedSvc.UpdateDiscountService(idHex, upd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// DeleteDiscount DELETE /seller/shops/:shopId/discounts/:id
func DeleteDiscount(c *gin.Context) {
	userHex, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(userHex.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user context"})
		return
	}
	shopIDhex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopIDhex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch shop"})
		return
	}
	if shop == nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}
	idHex := c.Param("id")
	existing, err := sharedSvc.GetDiscountByIDService(idHex)
	if err != nil {
		if err == sharedSvc.ErrDiscountNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	if existing.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}
	if err := sharedSvc.DeleteDiscountService(idHex); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
