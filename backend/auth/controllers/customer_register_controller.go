package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
)

// CustomerRegister handles new customer registration requests.
func CustomerRegister(c *gin.Context) {
	var req models.AuthCustomer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CustomerRegisterService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Customer registered successfully",
		"customer_id": req.CustomerID,
	})
}
