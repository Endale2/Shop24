package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLogout(ctx *gin.Context) {
	// Clear the JWT cookie by setting it to an empty value and immediate expiration
	ctx.SetCookie("admin_token", "", -1, "/", "", true, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Admin logged out successfully",
	})
}
