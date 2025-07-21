package controllers

import (
	"net/http" // Needed for os.Getenv
	"time"

	authServices "github.com/Endale2/DRPS/auth/services"
	"github.com/Endale2/DRPS/auth/utils"                          // Ensure utils package has GoogleOAuthConfig and ParseToken
	customerRepo "github.com/Endale2/DRPS/customers/repositories" // Ensure this path is correct
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	
	"go.mongodb.org/mongo-driver/bson"
)

// Only keep OTP and profile logic for customer authentication

// CustomerRequestOTP handles POST /auth/customer/request-otp
func CustomerRequestOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := authServices.SendCustomerOTP(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent if email is valid"})
}

// CustomerVerifyOTP handles POST /auth/customer/verify-otp
func CustomerVerifyOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		OTP   string `json:"otp" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cust, at, rt, profileComplete, err := authServices.VerifyCustomerOTP(req.Email, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Set cookies for cross-path local dev: Domain=.localhost, Path=/, HttpOnly, Secure=false, SameSite=Lax
	c.SetCookie("access_token", at, int((5 * time.Minute).Seconds()), "/", "localhost", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("refresh_token", rt, int((7 * 24 * time.Hour).Seconds()), "/", "localhost", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{"profile": cust, "profileComplete": profileComplete})
}

// TODO: Add CustomerTelegramOAuth handler here

// CustomerRefresh issues a new access token.
func CustomerRefresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token missing"})
		return
	}
	claims, err := utils.ParseToken(rt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}
	at, err := utils.CreateToken(claims.UserID, 5*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create access token"})
		return
	}
	c.SetCookie("access_token", at, int((5 * time.Minute).Seconds()), "/", ".localhost", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{"message": "access token refreshed"})
}

// CustomerMe returns the authenticated customer's profile.
func CustomerMe(c *gin.Context) {
	tokenStr, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access token missing"})
		return
	}

	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	custID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID in token"})
		return
	}

	customer, err := customerRepo.GetCustomerByID(custID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load customer data"})
		return
	}
	if customer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// UpdateCustomerMe allows the authenticated customer to update their profile
func UpdateCustomerMe(c *gin.Context) {
	tokenStr, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access token missing"})
		return
	}
	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	custID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID in token"})
		return
	}
	var in struct {
		FirstName  *string `json:"firstName"`
		LastName   *string `json:"lastName"`
		Phone      *string `json:"phone"`
		Address    *string `json:"address"`
		City       *string `json:"city"`
		State      *string `json:"state"`
		PostalCode *string `json:"postalCode"`
		Country    *string `json:"country"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := bson.M{}
	if in.FirstName != nil {
		update["firstName"] = *in.FirstName
	}
	if in.LastName != nil {
		update["lastName"] = *in.LastName
	}
	if in.Phone != nil {
		update["phone"] = *in.Phone
	}
	if in.Address != nil {
		update["address"] = *in.Address
	}
	if in.City != nil {
		update["city"] = *in.City
	}
	if in.State != nil {
		update["state"] = *in.State
	}
	if in.PostalCode != nil {
		update["postalCode"] = *in.PostalCode
	}
	if in.Country != nil {
		update["country"] = *in.Country
	}
	if len(update) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}
	update["updatedAt"] = time.Now()
	err = customerRepo.UpdateCustomerByID(custID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile"})
		return
	}
	customer, err := customerRepo.GetCustomerByID(custID)
	if err != nil || customer == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load updated profile"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// CustomerLogout clears auth cookies.
func CustomerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", ".localhost", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("refresh_token", "", -1, "/", ".localhost", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
