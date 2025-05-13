package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
)

// AdminLogin handles admin login requests, issues tokens, and sets them as HTTP-only cookies.
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

	// Set HTTP-only cookies for access and refresh tokens
	c.SetCookie("access_token", authData.AccessToken, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", authData.RefreshToken, int((7 * 24 * time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin logged in successfully",
		"admin":   adminData,
	})
}
