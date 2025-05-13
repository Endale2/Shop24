package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/utils"
	"github.com/gin-gonic/gin"
)

// RefreshToken issues a new access token using a valid refresh token.
func RefreshToken(c *gin.Context) {
	// Read refresh token from cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token missing"})
		return
	}

	// Verify the refresh token
	claims, err := utils.VerifyJWT(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	// Generate a new access token (15-minute expiry)
	accessToken, accessDur, err := utils.CreateJWT(claims.UserID, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
		return
	}

	// Set the new access token cookie
	c.SetCookie("access_token", accessToken, int(accessDur.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Access token refreshed"})
}
