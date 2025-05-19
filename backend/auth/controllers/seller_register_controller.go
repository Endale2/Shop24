package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
)

// SellerRegister handles new seller registration and sets auth cookies.
func SellerRegister(c *gin.Context) {
	var req models.AuthSeller
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register the seller (creates auth + seller record)
	if err := services.SellerRegisterService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Automatically log in the seller after successful registration
	sellerData, accessToken, refreshToken, err := services.SellerLoginService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registration succeeded but login failed"})
		return
	}

	// Set secure HTTP-only cookies
	c.SetCookie("access_token", accessToken, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, int((7 * 24 * time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Seller registered and logged in successfully",
		"seller_id": req.SellerID,
		"seller":    sellerData,
	})
}
