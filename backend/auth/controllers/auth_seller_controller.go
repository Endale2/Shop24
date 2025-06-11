package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/Endale2/DRPS/auth/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	"golang.org/x/oauth2"
)

// SellerRegister handles seller signup (email/password + auto-login).
func SellerRegister(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create auth & profile
	if err := services.SellerRegisterService(&models.AuthSeller{
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Immediately log them in
	sellerData, at, rt, err := services.SellerLoginService(&models.AuthSeller{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registration succeeded, login failed"})
		return
	}

	// Set cookies
	c.SetCookie("access_token", at, int((5*time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", rt, int((7*24*time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller registered & logged in",
		"seller":  sellerData,
	})
}

// SellerLogin handles email/password login.
func SellerLogin(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sellerData, at, rt, err := services.SellerLoginService(&models.AuthSeller{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", at, int((5*time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", rt, int((7*24*time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller logged in",
		"seller":  sellerData,
	})
}

// SellerOAuthRedirect redirects to Google’s OAuth consent page.
func SellerOAuthRedirect(c *gin.Context) {
	url := utils.GoogleOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

// SellerOAuthCallback handles Google’s OAuth callback.
func SellerOAuthCallback(c *gin.Context) {
	code := c.Query("code")
	tok, err := utils.GoogleOAuthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "oauth exchange failed"})
		return
	}
	idToken, _ := tok.Extra("id_token").(string)

	sellerData, at, rt, err := services.SellerLoginOAuth("google", idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", at, int((5*time.Minute).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", rt, int((7*24*time.Hour).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller logged in via Google",
		"seller":  sellerData,
	})
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

	c.SetCookie("access_token", at, int((5*time.Minute).Seconds()), "/", "", false, true)
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

// SellerLogout clears both auth cookies.
func SellerLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "seller logged out"})
}
