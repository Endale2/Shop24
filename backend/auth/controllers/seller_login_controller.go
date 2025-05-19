package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
)

// SellerLogin handles seller login requests, issues tokens, and sets them as HTTP-only cookies.
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

	// Set HTTP-only cookies for access and refresh tokens
	c.SetCookie(
		"access_token",
		authData.AccessToken,
		int((5 * time.Minute).Seconds()), // 5 minutes
		"/", "", false, true,
	)
	c.SetCookie(
		"refresh_token",
		authData.RefreshToken,
		int((7 * 24 * time.Hour).Seconds()), // 7 days
		"/", "", false, true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller logged in successfully",
		"seller":  sellerData,
	})
}
