package controllers

import (
    "net/http"
    "time"

    
    "github.com/Endale2/DRPS/auth/services"
    "github.com/Endale2/DRPS/auth/utils"
    "github.com/Endale2/DRPS/admin/repositories"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// RegisterLocalAdmin handles email/password signup.
func AdminRegister(c *gin.Context) {
    var req struct {
        Email     string `json:"email" binding:"required,email"`
        Password  string `json:"password" binding:"required,min=8"`
        FirstName string `json:"firstName" binding:"required"`
        LastName  string `json:"lastName"  binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.RegisterLocalAdmin(req.Email, req.Password, req.FirstName, req.LastName); err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

// LoginLocalAdmin handles email/password login.
func AdminLogin(c *gin.Context) {
    var req struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    cred, prof, err := services.LoginLocalAdmin(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    // set HTTP-only cookies
    c.SetCookie("access_token",  cred.AccessToken,  int((5 * time.Minute).Seconds()), "/", "", false, true)
    c.SetCookie("refresh_token", cred.RefreshToken, int((7 * 24 * time.Hour).Seconds()), "/", "", false, true)
    c.JSON(http.StatusOK, prof)
}

// AdminOAuth handles social login (e.g. Google).
func AdminOAuth(c *gin.Context) {
    var req struct {
        Provider   string `json:"provider" binding:"required"`   // e.g. "google"
        ProviderID string `json:"providerId" binding:"required"` // OAuth subject
        Email      string `json:"email" binding:"required,email"`
        FirstName  string `json:"firstName"`
        LastName   string `json:"lastName"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    cred, prof, err := services.LoginOAuthAdmin(
        req.Provider, req.ProviderID,
        req.Email, req.FirstName, req.LastName,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.SetCookie("access_token",  cred.AccessToken,  int((5 * time.Minute).Seconds()), "/", "", false, true)
    c.SetCookie("refresh_token", cred.RefreshToken, int((7 * 24 * time.Hour).Seconds()), "/", "", false, true)
    c.JSON(http.StatusOK, prof)
}

// AdminRefresh issues a new access-token given a valid refresh-token.
func AdminRefresh(c *gin.Context) {
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
    newAT, err := utils.CreateToken(claims.UserID, 5*time.Minute)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
        return
    }
    c.SetCookie("access_token", newAT, int((5*time.Minute).Seconds()), "/", "", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "access token refreshed"})
}

// AdminMe returns the current admin’s profile, based on access-token.
func GetAuthAdminMe(c *gin.Context) {
    at, err := c.Cookie("access_token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "access token missing"})
        return
    }
    claims, err := utils.ParseToken(at)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
        return
    }
    aid, err := primitive.ObjectIDFromHex(claims.UserID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }
    prof, err := repositories.GetAdminByID(aid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch profile"})
        return
    }
    if prof == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
        return
    }
    c.JSON(http.StatusOK, prof)
}

// AdminLogout clears both cookies.
func AdminLogout(c *gin.Context) {
    // clear access_token
    c.SetCookie("access_token",  "", -1, "/", "", false, true)
    // clear refresh_token
    c.SetCookie("refresh_token", "", -1, "/", "", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
