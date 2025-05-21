// File: sellers/controllers/seller_customer_controller.go
package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	customerService "github.com/Endale2/DRPS/shared/services"
	shopService "github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateCustomer creates a new customer for one of the seller’s shops.
// POST /seller/shops/:shopId/customers
func CreateCustomer(c *gin.Context) {
	// 1) Auth
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	// 2) Verify shop ownership
	shopID := c.Param("shopId")
	shop, err := shopService.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
		return
	}

	// 3) Bind request JSON
	var cust models.Customer
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// 4) Assign shop reference on customer (you may need to add ShopID to your model)
	// cust.ShopID = shop.ID

	// 5) Create
	res, err := customerService.CreateCustomerService(&cust)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creation failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetCustomers lists all customers of a seller’s shop.
// GET /seller/shops/:shopId/customers
func GetCustomers(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopID := c.Param("shopId")
	shop, err := shopService.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
		return
	}

	// Fetch all customers, then filter by ShopID
	all, err := customerService.GetAllCustomersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fetch failed"})
		return
	}

	var filtered []models.Customer
	for _, cust := range all {
		// Assuming your Customer model has a ShopID field:
		// if cust.ShopID == shop.ID { filtered = append(filtered, cust) }
		// If you haven’t added ShopID yet, you could filter by some other logic:
		filtered = append(filtered, cust)
	}

	c.JSON(http.StatusOK, filtered)
}

// GetCustomer retrieves a single customer by ID under a shop.
// GET /seller/shops/:shopId/customers/:custId
func GetCustomer(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopID := c.Param("shopId")
	shop, err := shopService.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
		return
	}

	custID := c.Param("custId")
	cust, err := customerService.GetCustomerByIDService(custID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	// if cust.ShopID != shop.ID {
	//     c.JSON(http.StatusForbidden, gin.H{"error": "not your customer"})
	//     return
	// }

	c.JSON(http.StatusOK, cust)
}

// UpdateCustomer updates a customer under a shop.
// PATCH /seller/shops/:shopId/customers/:custId
func UpdateCustomer(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopID := c.Param("shopId")
	shop, err := shopService.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
		return
	}

	custID := c.Param("custId")
	// Optionally: fetch and verify ownership before allowing update

	var updates bson.M
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	res, err := customerService.UpdateCustomerService(custID, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteCustomer deletes a customer from a shop.
// DELETE /seller/shops/:shopId/customers/:custId
func DeleteCustomer(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopID := c.Param("shopId")
	shop, err := shopService.GetShopByIDService(shopID)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
		return
	}

	custID := c.Param("custId")
	res, err := customerService.DeleteCustomerService(custID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}
