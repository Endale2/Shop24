package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
)

// AdminLogin handles admin login requests.
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin logged in successfully",
		"auth":    authData,
		"admin":   adminData,
	})
}
