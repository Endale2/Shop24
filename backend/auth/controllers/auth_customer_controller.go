package controllers

import (
	"fmt"
	"net/http"
	"os" // Needed for os.Getenv
	"strings"
	"time"

	authModels "github.com/Endale2/DRPS/auth/models"
	authServices "github.com/Endale2/DRPS/auth/services"
	"github.com/Endale2/DRPS/auth/utils" // Ensure utils package has GoogleOAuthConfig and ParseToken
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	sharedServices "github.com/Endale2/DRPS/shared/services" // Ensure this path is correct
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2" // Needed for oauth2.AccessTypeOffline
)

// CustomerLogin handles POST /customers/login.
// Expects JSON: { "email", "password", "shopId" }.
func CustomerLogin(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		ShopID   string `json:"shopId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cust, at, rt, err := authServices.CustomerLoginService(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if shopOID, err := primitive.ObjectIDFromHex(req.ShopID); err == nil {
		// ensure link exists (we ignore the return values)
		_, _, _ = sharedServices.LinkIfNotLinked(shopOID, cust.ID)
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    at,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   int((5 * time.Minute).Seconds()),
		HttpOnly: true,
		Secure:   false, // true for HTTPS
		SameSite: http.SameSiteLaxMode,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    rt,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   int((7 * 24 * time.Hour).Seconds()),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "customer": cust})
}

// CustomerRegister handles POST /customers/register.
// Expects JSON: { "username", "email", "password", "shopId", plus any optional profile fields }.
func CustomerRegister(c *gin.Context) {
	var req struct {
		Username  string `json:"username" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required"`
		ShopID    string `json:"shopId" binding:"required"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Phone     string `json:"phone"`
		Address   string `json:"address"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postal    string `json:"postalCode"`
		Country   string `json:"country"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authCust := &authModels.AuthCustomer{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	profile := &authServices.OptionalProfile{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Address:    req.Address,
		City:       req.City,
		State:      req.State,
		PostalCode: req.Postal,
		Country:    req.Country,
	}

	cust, err := authServices.CustomerRegisterService(authCust, profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if shopOID, err := primitive.ObjectIDFromHex(req.ShopID); err == nil {
		_, _, _ = sharedServices.LinkIfNotLinked(shopOID, cust.ID)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "customer registered successfully",
		"customerId": cust.ID.Hex(),
	})
}

// CustomerOAuthRedirect redirects to Google's OAuth consent page for customers.
// It accepts an optional shopId query parameter to link the customer to a shop.
func CustomerOAuthRedirect(c *gin.Context) {
	shopID := c.Query("shopId")
	// Encode state as "customer_state:<shopId>" or just "customer_state:" if empty
	state := fmt.Sprintf("customer_state:%s", shopID)
	url := utils.GoogleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

// CustomerOAuthCallback handles Google's OAuth callback for customers.
// It links the customer to a shop if a shopId was passed in the state.
func CustomerOAuthCallback(c *gin.Context) {
	// 1. Parse state to extract shopId
	rawState := c.Query("state")
	parts := strings.SplitN(rawState, ":", 2)
	var shopID string
	if len(parts) == 2 {
		shopID = parts[1]
	}

	// 2. Exchange the code for a token
	code := c.Query("code")
	tok, err := utils.GoogleOAuthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OAuth exchange failed for customer"})
		return
	}
	// Extract ID token
	idToken, ok := tok.Extra("id_token").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID token not found in OAuth response"})
		return
	}

	// 3. Create or fetch your customer, and issue app tokens
	cust, at, rt, err := authServices.CustomerLoginOAuth("google", idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 4. Link customer to shop if shopID present
	if shopID != "" {
		if shopOID, err := primitive.ObjectIDFromHex(shopID); err == nil {
			_, _, _ = sharedServices.LinkIfNotLinked(shopOID, cust.ID)
		}
	}

	// 5. Set HTTP-only cookies
	c.SetCookie("access_token", at, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", rt, int((7 * 24 * time.Hour).Seconds()), "/", "", false, true)

	// 6. Redirect to your customer-specific SPA route
	frontend := os.Getenv("FRONTEND_URL")
	if frontend == "" {
		frontend = "http://localhost:5174"
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/customer/dashboard", frontend))
}

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

	c.SetCookie("access_token", at, int((5 * time.Minute).Seconds()), "/", "", false, true)
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

// CustomerLogout clears auth cookies.
func CustomerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
