// project/auth/controllers/auth_admin_me_controller.go
package controllers

import (
	"net/http"

	//adminModels "github.com/Endale2/DRPS/admin/models"
	adminRepo "github.com/Endale2/DRPS/admin/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

// jwtSecret is the secret key used to sign tokens.
// In production, load this from a secure configuration.
var jwtSecret = []byte("your_secret_key")

// GetAuthAdminMe handles the request to show the currently logged-in admin's data
// by extracting the admin ID from the JWT in the access_token cookie.
func GetAuthAdminMe(c *gin.Context) {
	// Get the access token from the HTTP-only cookie
	accessTokenString, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Access token not found"})
		return
	}

	// Parse the JWT token
	token, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return jwtSecret, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid access token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the admin ID from the claims
		adminIDInterface := claims["admin_id"]
		adminIDHex, ok := adminIDInterface.(string) // Assuming admin_id is stored as string in JWT
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error - Invalid admin ID format in token"})
			return
		}

		// Convert the hex string ID to primitive.ObjectID
		objectAdminID, err := primitive.ObjectIDFromHex(adminIDHex)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Admin ID format in token"})
			return
		}

		// Retrieve Admin data by ID
		adminData, err := adminRepo.GetAdminByID(objectAdminID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve admin data"})
			return
		}
		if adminData == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Admin data not found"})
			return
		}

		c.JSON(http.StatusOK, adminData)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid access token"})
}

