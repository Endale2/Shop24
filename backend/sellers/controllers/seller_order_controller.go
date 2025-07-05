// sellers/controllers/order_controller.go
package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST   /seller/shops/:shopId/orders
func CreateOrder(c *gin.Context) {
	// we still need the user to be authenticated, but we don't use sellerID below
	_, _ = c.Get("user_id")

	shopHex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// TODO: fetch the shop by shopID and verify it belongs to the authenticated seller

	var in models.Order
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	in.ShopID = shopID
	in.Status = "pending"     // default status
	in.CreatedAt = time.Now() // service will override, but good practice
	in.UpdatedAt = time.Now()

	created, err := services.CreateOrderService(&in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// GET    /seller/shops/:shopId/orders
func ListOrders(c *gin.Context) {
	_, _ = c.Get("user_id")

	shopHex := c.Param("shopId")
	if _, err := primitive.ObjectIDFromHex(shopHex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// TODO: verify that shop belongs to seller

	orders, err := services.ListOrdersByShopService(shopHex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GET    /seller/shops/:shopId/orders/:orderId
func GetOrder(c *gin.Context) {
	_, _ = c.Get("user_id")

	shopHex := c.Param("shopId")
	if _, err := primitive.ObjectIDFromHex(shopHex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// TODO: verify that shop belongs to seller

	orderID := c.Param("orderId")
	o, err := services.GetOrderByIDService(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if o == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, o)
}

// GET    /seller/shops/:shopId/orders/:orderId/details
func GetOrderWithCustomerDetails(c *gin.Context) {
	_, _ = c.Get("user_id")

	shopHex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// TODO: verify that shop belongs to seller

	orderID := c.Param("orderId")
	orderWithCustomer, err := services.GetOrderWithCustomerDetails(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if orderWithCustomer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	// Verify the order belongs to this shop
	if orderWithCustomer.Order.ShopID != shopID {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, orderWithCustomer)
}

// PATCH  /seller/shops/:shopId/orders/:orderId
func UpdateOrder(c *gin.Context) {
	_, _ = c.Get("user_id")

	shopHex := c.Param("shopId")
	if _, err := primitive.ObjectIDFromHex(shopHex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// TODO: verify that shop belongs to seller

	var upd bson.M
	if err := c.ShouldBindJSON(&upd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID := c.Param("orderId")
	updated, err := services.UpdateOrderService(orderID, upd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DELETE /seller/shops/:shopId/orders/:orderId
func DeleteOrder(c *gin.Context) {
	_, _ = c.Get("user_id")

	shopHex := c.Param("shopId")
	if _, err := primitive.ObjectIDFromHex(shopHex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// TODO: verify that shop belongs to seller

	orderID := c.Param("orderId")
	if err := services.DeleteOrderService(orderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
