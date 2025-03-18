package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
)

// SellerRegister handles new seller registration requests.
func SellerRegister(c *gin.Context) {
	var req models.AuthSeller
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.SellerRegisterService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seller registered successfully", "seller_id": req.SellerID})
}
