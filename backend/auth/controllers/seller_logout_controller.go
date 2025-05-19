package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SellerLogout clears the seller’s auth cookies.
func SellerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Seller logged out successfully"})
}
