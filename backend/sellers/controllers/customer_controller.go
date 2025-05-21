package controllers

import (
    "net/http"

    scService "github.com/Endale2/DRPS/shared/services"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

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


// Get one customer’s details by real customer ID,
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

    c.JSON(http.StatusOK, cust)
}