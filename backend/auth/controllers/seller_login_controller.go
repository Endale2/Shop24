package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/utils"
	"github.com/gin-gonic/gin"
)

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
