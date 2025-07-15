package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Endale2/DRPS/auth/services"
	"github.com/Endale2/DRPS/auth/utils"
	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	sellerSvc "github.com/Endale2/DRPS/sellers/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

// SellerOAuthRedirect redirects to Google’s OAuth consent page.
func SellerOAuthRedirect(c *gin.Context) {
	url := utils.GetGoogleOAuthConfigForSeller().AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

// SellerOAuthCallback handles Google’s OAuth callback.
func SellerOAuthCallback(c *gin.Context) {
	// 1. Exchange the code for a token
	code := c.Query("code")
	tok, err := utils.GetGoogleOAuthConfigForSeller().Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "oauth exchange failed"})
		return
	}
	idToken, _ := tok.Extra("id_token").(string)

	// 2. Create or fetch your seller, and issue app tokens
	_, at, rt, err := services.SellerLoginOAuth("google", idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 3. Set HTTP‐only cookies so the browser will send them automatically
	c.SetCookie("access_token", at, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", rt, int((7 * 24 * time.Hour).Seconds()), "/", "", false, true)

	// 4. Redirect into your SPA:
	//    replace `http://localhost:5174/shops` with your real front-end URL
	frontend := os.Getenv("FRONTEND_URL")
	if frontend == "" {
		frontend = "http://localhost:5174"
	}
	c.Redirect(http.StatusFound, frontend+"/shops")
}

// SellerTelegramOAuth handles Telegram OAuth login.
func SellerTelegramOAuth(c *gin.Context) {
	var telegramData map[string]string
	if err := c.ShouldBindJSON(&telegramData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid telegram data"})
		return
	}

	// Verify Telegram login signature
	botToken := utils.GetTelegramBotTokenForSeller()
	payload, err := utils.VerifyTelegramLogin(telegramData, botToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid telegram signature"})
		return
	}

	// Create or fetch seller profile
	seller, at, rt, err := services.SellerLoginTelegram(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set HTTP-only cookies
	c.SetCookie("access_token", at, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", rt, int((7 * 24 * time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, seller)
}

// SellerRefresh issues a new access token using the refresh_token cookie.
func SellerRefresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token missing"})
		return
	}

	claims, err := utils.ParseToken(rt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	at, err := utils.CreateToken(claims.UserID, 5*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create access token"})
		return
	}

	c.SetCookie("access_token", at, int((5 * time.Minute).Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "access token refreshed"})
}

// GetCurrentSeller returns the logged-in seller’s profile.
func GetCurrentSeller(c *gin.Context) {
	tokenStr, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access token missing"})
		return
	}
	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
		return
	}

	sellerID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed seller ID"})
		return
	}

	seller, err := sellerRepo.GetSellerByID(sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load seller"})
		return
	}
	if seller == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "seller not found"})
		return
	}

	c.JSON(http.StatusOK, seller)
}

// UpdateCurrentSeller PATCH /seller/me
func UpdateCurrentSeller(c *gin.Context) {
	// 1) get seller ID from JWT
	userHex, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(userHex.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid seller ID"})
		return
	}

	// 2) bind the fields we allow them to change
	var in struct {
		FirstName    *string `json:"firstName"`
		LastName     *string `json:"lastName"`
		ProfileImage *string `json:"profileImage"`
		Phone        *string `json:"phone"`
		Address      *string `json:"address"`
		BusinessName *string `json:"businessName"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3) build update document
	upd := bson.M{}
	if in.FirstName != nil {
		upd["first_name"] = *in.FirstName
	}
	if in.LastName != nil {
		upd["last_name"] = *in.LastName
	}
	if in.ProfileImage != nil {
		upd["profile_image"] = *in.ProfileImage
	}
	if in.Phone != nil {
		upd["phone"] = *in.Phone
	}
	if in.Address != nil {
		upd["address"] = *in.Address
	}
	if in.BusinessName != nil {
		upd["business_name"] = *in.BusinessName
	}

	// 4) call service
	updated, err := sellerSvc.UpdateSellerService(sellerID.Hex(), upd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// SellerLogout clears both auth cookies.
func SellerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "seller logged out"})
}
