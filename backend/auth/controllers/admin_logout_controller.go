// File: auth/controllers/admin_logout_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
