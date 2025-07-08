package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetOrders retrieves all orders (optionally filter by shop_id)
func GetOrders(c *gin.Context) {
	shopID := c.Query("shop_id")
	var orders []models.Order
	var err error
	if shopID != "" {
		orders, err = services.ListOrdersByShopService(shopID)
	} else {
		orders, err = services.ListOrdersByShopService("") // List all
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrder retrieves a single order by its ID (with customer details)
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := services.GetOrderWithCustomerDetails(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// CreateOrder creates a new order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	result, err := services.CreateOrderService(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating order"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// UpdateOrder updates an existing order by its ID
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var updatedData map[string]interface{}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	result, err := services.UpdateOrderService(id, bson.M(updatedData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating order"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteOrder deletes an order by its ID
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteOrderService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}

// UpdateOrderStatus updates only the status of an order
func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var payload struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil || payload.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status payload"})
		return
	}
	result, err := services.UpdateOrderService(id, bson.M{"status": payload.Status})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating order status"})
		return
	}
	c.JSON(http.StatusOK, result)
}
