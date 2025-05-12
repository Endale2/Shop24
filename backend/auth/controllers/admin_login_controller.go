package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
)


// AdminLogin handles admin login requests, generates tokens, and sets them as HTTP-only cookies.
func AdminLogin(c *gin.Context) {
	var req models.AuthAdmin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authData, adminData, err := services.AdminLoginService(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set the access token as an HTTP-only cookie (expires in 15 minutes = 900 seconds)
	c.SetCookie("access_token", authData.AccessToken, 300, "/", "", true, true)
	// Set the refresh token as an HTTP-only cookie (expires in 7 days)
	c.SetCookie("refresh_token", authData.RefreshToken, 7*24*3600, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin logged in successfully",
		"admin":   adminData,
	})
}
