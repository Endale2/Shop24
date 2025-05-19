package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
)

// SellerLogin handles seller login, issues tokens, and sets HTTP-only cookies.
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

	// Set access token cookie (5 min)
	c.SetCookie(
		"access_token",
		authData.AccessToken,
		int((5 * time.Minute).Seconds()),
		"/", "", false, true,
	)
	// Set refresh token cookie (7 days)
	c.SetCookie(
		"refresh_token",
		authData.RefreshToken,
		int((7 * 24 * time.Hour).Seconds()),
		"/", "", false, true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller logged in successfully",
		"seller":  sellerData,
	})
}
