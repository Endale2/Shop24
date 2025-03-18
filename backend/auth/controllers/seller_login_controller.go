package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
)

// SellerLogin handles seller login requests.
func SellerLogin(c *gin.Context) {
	var req models.AuthSeller
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authData, sellerData, err := services.SellerLoginService(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller logged in successfully",
		"auth":    authData,
		"seller":  sellerData,
	})
}
