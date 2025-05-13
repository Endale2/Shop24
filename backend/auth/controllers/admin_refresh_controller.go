package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/utils"
	"github.com/gin-gonic/gin"
)

// RefreshToken handles refresh token requests, verifies the refresh token, and issues a new access token.
func RefreshToken(c *gin.Context) {
	// 1) Read the refresh token cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token missing"})
		return
	}

	// 2) Parse and validate the refresh token
	claims, err := utils.ParseToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	// 3) Issue new access token (5-minute expiry)
	newAccessToken, err := utils.CreateToken(claims.UserID, 5*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
		return
	}

	// 4) Set the new access token cookie
	c.SetCookie(
		"access_token",
		newAccessToken,
		int((5 * time.Minute).Seconds()),
		"/", "", false, true,
	)

	// 5) Return success
	c.JSON(http.StatusOK, gin.H{"message": "Access token refreshed"})
}
