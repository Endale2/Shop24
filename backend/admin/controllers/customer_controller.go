package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/customers/models"
	"github.com/Endale2/DRPS/shared/services"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateCustomer handles the creation of a new customer.
func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := services.CreateCustomerService(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating customer"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetCustomer retrieves a single customer by its ID.
func GetCustomer(c *gin.Context) {
	id := c.Param("id")

	customer, err := services.GetCustomerByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// GetCustomers retrieves all customers.
func GetCustomers(c *gin.Context) {
	customers, err := services.GetAllCustomersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// UpdateCustomer updates an existing customer identified by its ID.
func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")

	var updatedData map[string]interface{}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := services.UpdateCustomerService(id, bson.M(updatedData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating customer"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteCustomer deletes a customer by its ID.
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	result, err := services.DeleteCustomerService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting customer"})
		return
	}
	c.JSON(http.StatusOK, result)
}
