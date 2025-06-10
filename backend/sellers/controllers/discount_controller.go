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
    c.JSON(http.StatusOK, d)
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
                    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid startAt format"}); return
                }
            }
        case "endAt":
            if s, ok := v.(string); ok {
                if t, err := parseTimeField(s); err == nil {
                    upd["end_at"] = t
                } else {
                    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid endAt format"}); return
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
