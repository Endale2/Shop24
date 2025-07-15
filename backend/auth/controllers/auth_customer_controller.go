package controllers

import (
	"fmt"
	"net/http"
	"os" // Needed for os.Getenv
	"strings"
	"time"

	authServices "github.com/Endale2/DRPS/auth/services"
	"github.com/Endale2/DRPS/auth/utils" // Ensure utils package has GoogleOAuthConfig and ParseToken
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	sharedServices "github.com/Endale2/DRPS/shared/services" // Ensure this path is correct
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2" // Needed for oauth2.AccessTypeOffline
)

// CustomerOAuthRedirect redirects to Google's OAuth consent page for customers.
// It accepts an optional shopId query parameter to link the customer to a shop.
func CustomerOAuthRedirect(c *gin.Context) {
	shopID := c.Query("shopId")
	// Encode state as "customer_state:<shopId>" or just "customer_state:" if empty
	state := fmt.Sprintf("customer_state:%s", shopID)
	url := utils.GetGoogleOAuthConfigForCustomer().AuthCodeURL(state, oauth2.AccessTypeOffline)
	fmt.Println("CUSTOMER OAUTH REDIRECT URL:", url)
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
	tok, err := utils.GetGoogleOAuthConfigForCustomer().Exchange(c, code)
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
	frontend := os.Getenv("CUSTOMER_FRONTEND_URL")
	if frontend == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CUSTOMER_FRONTEND_URL not set in environment"})
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/customer/dashboard", frontend))
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
