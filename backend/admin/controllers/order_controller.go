package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateOrder handles the creation of a new order.
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

// GetOrder retrieves a single order by its ID.
func GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := services.GetOrderByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// GetOrders retrieves all orders.
func GetOrders(c *gin.Context) {
	orders, err := services.GetAllOrdersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// UpdateOrder updates fields of an order identified by its ID.
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

// DeleteOrder deletes an order by its ID.
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	result, err := services.DeleteOrderService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting order"})
		return
	}
	c.JSON(http.StatusOK, result)
}
