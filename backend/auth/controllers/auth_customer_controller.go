// File: auth/controllers/auth_customer_controller.go
package controllers

import (
    "net/http"
    "os"
    "time"

    "github.com/Endale2/DRPS/auth/models"
    authRepo   "github.com/Endale2/DRPS/auth/repositories"
    authSvc    "github.com/Endale2/DRPS/auth/services"
    "github.com/Endale2/DRPS/auth/utils"
    sharedCust "github.com/Endale2/DRPS/shared/models"
    custRepo   "github.com/Endale2/DRPS/shared/repositories"
    sharedSvc  "github.com/Endale2/DRPS/shared/services"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerRegister handles signup, JWT issuance, and shop-linking.
func CustomerRegister(c *gin.Context) {
    var payload struct {
        models.AuthCustomer
        // these fields come through JSON
        FirstName  string `json:"firstName"`
        LastName   string `json:"lastName"`
        Phone      string `json:"phone"`
        Address    string `json:"address"`
        City       string `json:"city"`
        State      string `json:"state"`
        PostalCode string `json:"postalCode"`
        Country    string `json:"country"`
        ShopID     string `json:"shopId" binding:"required"`
    }

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 1. Prevent double-registration
    if _, err := authRepo.FindAuthCustomerByEmail(payload.Email); err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email already registered; please login"})
        return
    }

    // 2. Create global Customer profile
    profile := &sharedCust.Customer{
        FirstName:  payload.FirstName,
        LastName:   payload.LastName,
        Email:      payload.Email,
        Phone:      payload.Phone,
        Address:    payload.Address,
        City:       payload.City,
        State:      payload.State,
        PostalCode: payload.PostalCode,
        Country:    payload.Country,
    }
    ins, err := custRepo.CreateCustomer(profile)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer profile"})
        return
    }
    profile.ID = ins.InsertedID.(primitive.ObjectID)

    // 3. Create the AuthCustomer record
    payload.CustomerID = profile.ID
    if _, err := authRepo.CreateAuthCustomer(&payload.AuthCustomer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create authentication record"})
        return
    }

    // 4. Issue JWTs
    accessToken, _  := utils.CreateToken(profile.ID.Hex(), 5*time.Minute)
    refreshToken, _ := utils.CreateToken(profile.ID.Hex(), 7*24*time.Hour)
    c.SetCookie("access_token",  accessToken,  int(5*time.Minute.Seconds()),   "/", "", false, true)
    c.SetCookie("refresh_token", refreshToken, int((7*24*time.Hour).Seconds()), "/", "", false, true)

    // 5. Link into shop (only if not already linked)
    if shopOID, err := primitive.ObjectIDFromHex(payload.ShopID); err == nil {
        if linked, _ := sharedSvc.IsCustomerLinked(shopOID, profile.ID); !linked {
            sharedSvc.LinkCustomerToShop(shopOID, profile.ID)
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registered & linked", "customer": profile})
}

// CustomerLogin authenticates, issues JWTs, and optionally links to a shop.
func CustomerLogin(c *gin.Context) {
    var payload struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
        ShopID   string `json:"shopId,omitempty"`
    }
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 1. Authenticate & get profile + tokens
    custData, accessToken, refreshToken, err := authSvc.CustomerLoginService(&models.AuthCustomer{
        Email:    payload.Email,
        Password: payload.Password,
    })
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // 2. Set cookies
    c.SetCookie("access_token",  accessToken,  int(5*time.Minute.Seconds()),   "/", "", false, true)
    c.SetCookie("refresh_token", refreshToken, int((7*24*time.Hour).Seconds()), "/", "", false, true)

    // 3. Optionally link to new shop
    if payload.ShopID != "" {
        if shopOID, err := primitive.ObjectIDFromHex(payload.ShopID); err == nil {
            if linked, _ := sharedSvc.IsCustomerLinked(shopOID, custData.ID); !linked {
                sharedSvc.LinkCustomerToShop(shopOID, custData.ID)
            }
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Logged in", "customer": custData})
}

// CustomerRefresh issues a fresh access token using the refresh_token cookie.
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
    at, _ := utils.CreateToken(claims.UserID, 5*time.Minute)
    c.SetCookie("access_token", at, int(5*time.Minute.Seconds()), "/", "", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "Access token refreshed"})
}

// CustomerMe returns the currently authenticated customer’s profile.
func CustomerMe(c *gin.Context) {
    if os.Getenv("JWT_SECRET") == "" {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Server misconfiguration"})
        return
    }
    tokenStr, err := c.Cookie("access_token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token missing"})
        return
    }
    claims, err := utils.ParseToken(tokenStr)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
    }
    custID, err := primitive.ObjectIDFromHex(claims.UserID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed customer ID"})
        return
    }
    cust, err := custRepo.GetCustomerByID(custID.Hex())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load customer"})
        return
    }
    if cust == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
        return
    }
    c.JSON(http.StatusOK, cust)
}

// CustomerLogout clears both auth cookies.
func CustomerLogout(c *gin.Context) {
    c.SetCookie("access_token",  "", -1, "/", "", false, true)
    c.SetCookie("refresh_token", "", -1, "/", "", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
