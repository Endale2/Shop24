// File: auth/controllers/seller_me_controller.go
package controllers

import (
	"net/http"
	"os"

	"github.com/Endale2/DRPS/auth/utils"
	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
