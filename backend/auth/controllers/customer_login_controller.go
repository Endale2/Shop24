package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
)

// CustomerLogin logs in a customer and sets JWT cookies.
func CustomerLogin(c *gin.Context) {
	var req models.AuthCustomer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	custData, accessToken, refreshToken, err := services.CustomerLoginService(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, int((7*24*time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "customer": custData})
}
