package controllers

import (
	"net/http"
	"time"
    "os"
	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/utils"
	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// SellerLogin handles seller login and sets HTTP-only cookies for tokens.
func SellerLogin(c *gin.Context) {
	var req models.AuthSeller
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sellerData, accessToken, refreshToken, err := services.SellerLoginService(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set access token cookie (5 min)
	c.SetCookie("access_token", accessToken, int((5 * time.Minute).Seconds()), "/", "", false, true)

	// Set refresh token cookie (7 days)
	c.SetCookie("refresh_token", refreshToken, int((7 * 24 * time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller logged in successfully",
		"seller":  sellerData,
	})
}


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




// SellerRefresh issues a new access token using the refresh_token cookie.
func SellerRefresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token missing"})
		return
	}

	claims, err := utils.ParseToken(rt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	// Issue new 5-minute access token
	at, _ := utils.CreateToken(claims.UserID, 5*time.Minute)
	c.SetCookie("access_token", at, int((5*time.Minute).Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Access token refreshed"})
}



// GetCurrentSeller returns the authenticated seller's profile.
func GetCurrentSeller(c *gin.Context) {
	// 1) Load JWT_SECRET (ensure main.go has loaded .env)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server misconfiguration"})
		return
	}

	// 2) Extract access_token cookie
	tokenStr, err := c.Cookie("access_token")
	if err != nil || tokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token missing"})
		return
	}

	// 3) Parse token and get claims
	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// 4) Convert the UserID claim (here SellerID) to ObjectID
	sellerID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed seller ID"})
		return
	}

	// 5) Fetch seller record
	seller, err := sellerRepo.GetSellerByID(sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load seller data"})
		return
	}
	if seller == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}

	// 6) Return seller profile (no sensitive fields in this model)
	c.JSON(http.StatusOK, seller)
}



// SellerLogout clears the seller’s auth cookies.
func SellerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Seller logged out successfully"})
}



