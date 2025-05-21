// File: sellers/controllers/seller_customer_controller.go
package controllers

import (
    "net/http"

    "github.com/Endale2/DRPS/shared/models"
    customerService "github.com/Endale2/DRPS/shared/services"
    shopService     "github.com/Endale2/DRPS/shared/services"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateCustomer creates a new customer under the seller’s shop.
// POST /seller/shops/:shopId/customers
func CreateCustomer(c *gin.Context) {
    // 1) Authenticate seller
    sellerHex, _ := c.Get("user_id")
    sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

    // 2) Verify shop ownership
    shopID := c.Param("shopId")
    shop, err := shopService.GetShopByIDService(shopID)
    if err != nil || shop.OwnerID != sellerID {
        c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
        return
    }

    // 3) Bind JSON
    var cust models.Customer
    if err := c.ShouldBindJSON(&cust); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }

    // 4) Stamp the ShopID
    cust.ShopID = shop.ID

    // 5) Create
    res, err := customerService.CreateCustomerService(&cust)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "creation failed"})
        return
    }
    c.JSON(http.StatusOK, res)
}

// GetCustomers lists only this shop’s customers.
// GET /seller/shops/:shopId/customers
func GetCustomers(c *gin.Context) {
    sellerHex, _ := c.Get("user_id")
    sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

    shopID := c.Param("shopId")
    shop, err := shopService.GetShopByIDService(shopID)
    if err != nil || shop.OwnerID != sellerID {
        c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
        return
    }

    // fetch with filter
    filter := bson.M{ "shopId": shop.ID }
    list, err := customerService.GetAllCustomersServiceWithFilter(filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "fetch failed"})
        return
    }
    c.JSON(http.StatusOK, list)
}

// GetCustomer retrieves one customer under this shop.
// GET /seller/shops/:shopId/customers/:custId
func GetCustomer(c *gin.Context) {
    sellerHex, _ := c.Get("user_id")
    sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

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
    if cust.ShopID != shop.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "not your customer"})
        return
    }
    c.JSON(http.StatusOK, cust)
}

// UpdateCustomer updates a customer under this shop.
// PATCH /seller/shops/:shopId/customers/:custId
func UpdateCustomer(c *gin.Context) {
    sellerHex, _ := c.Get("user_id")
    sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

    shopID := c.Param("shopId")
    shop, err := shopService.GetShopByIDService(shopID)
    if err != nil || shop.OwnerID != sellerID {
        c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
        return
    }

    custID := c.Param("custId")
    cust, err := customerService.GetCustomerByIDService(custID)
    if err != nil || cust.ShopID != shop.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "not your customer"})
        return
    }

    var updates bson.M
    if err := c.ShouldBindJSON(&updates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }
    // Prevent ShopID reassignment:
    delete(updates, "shopId")

    res, err := customerService.UpdateCustomerService(custID, updates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
        return
    }
    c.JSON(http.StatusOK, res)
}

// DeleteCustomer removes a customer under this shop.
// DELETE /seller/shops/:shopId/customers/:custId
func DeleteCustomer(c *gin.Context) {
    sellerHex, _ := c.Get("user_id")
    sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

    shopID := c.Param("shopId")
    shop, err := shopService.GetShopByIDService(shopID)
    if err != nil || shop.OwnerID != sellerID {
        c.JSON(http.StatusForbidden, gin.H{"error": "not authorized for this shop"})
        return
    }

    custID := c.Param("custId")
    cust, err := customerService.GetCustomerByIDService(custID)
    if err != nil || cust.ShopID != shop.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "not your customer"})
        return
    }

    res, err := customerService.DeleteCustomerService(custID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
        return
    }
    c.JSON(http.StatusOK, res)
}
