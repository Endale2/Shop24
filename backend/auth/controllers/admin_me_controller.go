package controllers

import (
	"net/http"

	adminRepo "github.com/Endale2/DRPS/admin/repositories"
	"github.com/Endale2/DRPS/auth/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAuthAdminMe returns the authenticated admin's information from the JWT in the access_token cookie.
func GetAuthAdminMe(c *gin.Context) {
	// Retrieve the access token from cookie
	tokenStr, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token missing"})
		return
	}

	// Parse the JWT and extract claims
	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Convert string ID to ObjectID
	adminID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID in token"})
		return
	}

	// Fetch admin data from DB
	admin, err := adminRepo.GetAdminByID(adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch admin data"})
		return
	}
	if admin == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	// Return admin data
	c.JSON(http.StatusOK, admin)
}
