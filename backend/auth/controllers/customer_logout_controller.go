package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomerLogout clears the auth cookies.
func CustomerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
