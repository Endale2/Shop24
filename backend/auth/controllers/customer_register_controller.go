package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
)

// CustomerRegister registers and logs in the customer in one call.
func CustomerRegister(c *gin.Context) {
	var req models.AuthCustomer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user records
	if err := services.CustomerRegisterService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Auto-login
	custData, accessToken, refreshToken, err := services.CustomerLoginService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registered but login failed"})
		return
	}

	// Set cookies
	c.SetCookie("access_token", accessToken, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, int((7*24*time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Registered & logged in", "customer": custData})
}
