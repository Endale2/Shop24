// File: auth/controllers/customer_me_controller.go
package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/auth/utils"
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerMe returns the currently-authenticated customer's profile.
func CustomerMe(c *gin.Context) {
	// 1) Get the access_token cookie
	tokenStr, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token missing"})
		return
	}

	// 2) Parse and validate the JWT
	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// 3) Convert the UserID (customer ID) into ObjectID
	custID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID in token"})
		return
	}

	// 4) Fetch the customer record
	customer, err := customerRepo.GetCustomerByID(custID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load customer data"})
		return
	}
	if customer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// 5) Return the customer profile
	c.JSON(http.StatusOK, customer)
}
