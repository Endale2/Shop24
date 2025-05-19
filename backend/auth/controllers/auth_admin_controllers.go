package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Endale2/DRPS/auth/utils"
	adminRepo "github.com/Endale2/DRPS/admin/repositories"
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





// AdminRegister handles new admin registration requests.
func AdminRegister(c *gin.Context) {
	var req models.AuthAdmin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AdminRegisterService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin registered successfully", "admin_id": req.AdminID})
}






// RefreshToken handles refresh token requests, verifies the refresh token, and issues a new access token.
func AdminRefresh(c *gin.Context) {
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





// AdminLogout clears the authentication cookies.
func AdminLogout(c *gin.Context) {
	// Invalidate the access_token cookie
	c.SetCookie(
		"access_token",   // name
		"",               // value
		-1,               // maxAge < 0 means delete now
		"/",              // path
		"",               // domain (same as login)
		false,            // secure (same as login)
		true,             // httpOnly
	)

	// Invalidate the refresh_token cookie
	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}



