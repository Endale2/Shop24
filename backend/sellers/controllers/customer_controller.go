package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/sellers/models"
	"github.com/Endale2/DRPS/sellers/services"
	scService "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var customerSegmentSvc = services.NewCustomerSegmentService()

// Add an existing customer to the shop (by customer ID).
// POST /seller/shops/:shopId/customers/link
func LinkCustomer(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your shop"})
		return
	}

	// payload: { "customerId": "<hex>" }
	var body struct{ CustomerID string }
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	custID, err := primitive.ObjectIDFromHex(body.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	linkID, err := scService.LinkCustomerToShop(shop.ID, custID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not link"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"linkId": linkID.Hex()})
}

// List only the customers linked to this shop.
// GET /seller/shops/:shopId/customers
func GetLinkedCustomers(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your shop"})
		return
	}

	links, err := scService.GetCustomerLinksForShop(shop.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch links"})
		return
	}

	// collect actual Customer records
	var result []interface{}
	for _, link := range links {
		cust, err := scService.GetCustomerByIDService(link.CustomerID.Hex())
		if err == nil {
			// optionally add the link ID so you can unlink later:
			result = append(result, gin.H{
				"linkId":   link.ID.Hex(),
				"customer": cust,
			})
		}
	}
	c.JSON(http.StatusOK, result)
}

// Unlink (remove) a customer from this shop.
// DELETE /seller/shops/:shopId/customers/link/:linkId
func UnlinkCustomer(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your shop"})
		return
	}

	linkIDhex := c.Param("linkId")
	linkID, err := primitive.ObjectIDFromHex(linkIDhex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid link ID"})
		return
	}

	if err := scService.UnlinkCustomerFromShop(linkID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not unlink"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer removed from shop"})
}

// Get one customer's details by real customer ID,
// but only if linked to this shop.
// GET /seller/shops/:shopId/customers/:customerId
func GetCustomerDetail(c *gin.Context) {
	// 1) Auth & ownership
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopIDhex := c.Param("shopId")
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your shop"})
		return
	}

	// 2) parse customerId
	custHex := c.Param("customerId")
	custID, err := primitive.ObjectIDFromHex(custHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	// 3) check linkage
	linked, err := scService.IsCustomerLinked(shop.ID, custID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not verify linkage"})
		return
	}
	if !linked {
		c.JSON(http.StatusForbidden, gin.H{"error": "customer not in this shop"})
		return
	}

	// 4) fetch real customer record
	cust, err := scService.GetCustomerByIDService(custHex)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	// 5) fetch customer segments
	segments, err := customerSegmentSvc.GetSegmentsByCustomerService(shop.ID, custID)
	if err != nil {
		// Don't fail the request if segments can't be fetched
		segments = []models.CustomerSegment{}
	}

	c.JSON(http.StatusOK, gin.H{
		"customer": cust,
		"segments": segments,
	})
}

// ===== CUSTOMER SEGMENTS =====

// CreateCustomerSegmentInput is the expected JSON body for creating a customer segment
type CreateCustomerSegmentInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Color       string `json:"color"` // Hex color code
}

// CreateCustomerSegment handles POST /seller/shops/:shopId/customer-segments
func CreateCustomerSegment(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopIDhex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// Verify shop ownership
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	var in CreateCustomerSegmentInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	segment, err := customerSegmentSvc.CreateCustomerSegmentService(shopID, sellerID, in.Name, in.Description, in.Color)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create segment: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, segment)
}

// GetCustomerSegments handles GET /seller/shops/:shopId/customer-segments
func GetCustomerSegments(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopIDhex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// Verify shop ownership
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	segments, err := customerSegmentSvc.GetCustomerSegmentsByShopService(shopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch segments: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, segments)
}

// UpdateCustomerSegmentInput is the JSON body for updating a customer segment
type UpdateCustomerSegmentInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
}

// UpdateCustomerSegment handles PATCH /seller/shops/:shopId/customer-segments/:segmentId
func UpdateCustomerSegment(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	segmentHex := c.Param("segmentId")
	segmentID, err := primitive.ObjectIDFromHex(segmentHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid segment ID"})
		return
	}

	var in UpdateCustomerSegmentInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	updates := bson.M{}
	if in.Name != nil {
		updates["name"] = *in.Name
	}
	if in.Description != nil {
		updates["description"] = *in.Description
	}
	if in.Color != nil {
		updates["color"] = *in.Color
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	err = customerSegmentSvc.UpdateCustomerSegmentService(segmentID, sellerID, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update segment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "segment updated successfully"})
}

// DeleteCustomerSegment handles DELETE /seller/shops/:shopId/customer-segments/:segmentId
func DeleteCustomerSegment(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	segmentHex := c.Param("segmentId")
	segmentID, err := primitive.ObjectIDFromHex(segmentHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid segment ID"})
		return
	}

	err = customerSegmentSvc.DeleteCustomerSegmentService(segmentID, sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete segment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "segment deleted successfully"})
}

// AddCustomerToSegmentInput is the JSON body for adding a customer to a segment
type AddCustomerToSegmentInput struct {
	CustomerID string `json:"customerId" binding:"required"`
}

// AddCustomerToSegment handles POST /seller/shops/:shopId/customer-segments/:segmentId/customers
func AddCustomerToSegment(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	segmentHex := c.Param("segmentId")
	segmentID, err := primitive.ObjectIDFromHex(segmentHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid segment ID"})
		return
	}

	var in AddCustomerToSegmentInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	customerID, err := primitive.ObjectIDFromHex(in.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	err = customerSegmentSvc.AddCustomerToSegmentService(segmentID, customerID, sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add customer to segment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer added to segment successfully"})
}

// RemoveCustomerFromSegment handles DELETE /seller/shops/:shopId/customer-segments/:segmentId/customers/:customerId
func RemoveCustomerFromSegment(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	shopIDhex := c.Param("shopId")
	shop, err := scService.GetShopByIDService(shopIDhex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	segmentHex := c.Param("segmentId")
	segmentID, err := primitive.ObjectIDFromHex(segmentHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid segment ID"})
		return
	}

	customerHex := c.Param("customerId")
	customerID, err := primitive.ObjectIDFromHex(customerHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	err = customerSegmentSvc.RemoveCustomerFromSegmentService(segmentID, customerID, sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove customer from segment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer removed from segment successfully"})
}
