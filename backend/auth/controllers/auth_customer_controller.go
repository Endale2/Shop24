package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/utils"
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerLogin logs in a customer and sets JWT cookies.
func CustomerLogin(c *gin.Context) {
	var req models.AuthCustomer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	custData, accessToken, refreshToken, err := services.CustomerLoginService(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, int((7*24*time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "customer": custData})
}




// CustomerRegister registers and logs in the customer in one call.
func CustomerRegister(c *gin.Context) {
	var req models.AuthCustomer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user records
	if err := services.CustomerRegisterService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Auto-login
	custData, accessToken, refreshToken, err := services.CustomerLoginService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registered but login failed"})
		return
	}

	// Set cookies
	c.SetCookie("access_token", accessToken, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, int((7*24*time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Registered & logged in", "customer": custData})
}





// CustomerRefresh issues a new access token if the refresh token is valid.
func CustomerRefresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token missing"})
		return
	}
	claims, err := utils.ParseToken(rt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	at, err := utils.CreateToken(claims.UserID, 5*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
		return
	}
	c.SetCookie("access_token", at, int((5*time.Minute).Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Access token refreshed"})
}




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




// CustomerLogout clears the auth cookies.
func CustomerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
