package controllers

import (
	"net/http"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/utils"
	authModels "github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	customerModels "github.com/Endale2/DRPS/customers/models"
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




// RegisterCustomer handles POST /customers/register
func RegisterCustomer(c *gin.Context) {
	// 1) Bind into a lightweight map or into AuthCustomer for fields username/email/password
	var payload map[string]string
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2) Validate required fields
	required := []string{"firstName", "lastName", "username", "email", "password", "phone", "address", "city", "state", "postalCode", "country"}
	for _, key := range required {
		if payload[key] == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": key + " is required"})
			return
		}
	}

	// 3) Check for existing auth record
	if existing, _ := authRepo.FindAuthCustomerByEmail(payload["email"]); existing != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "customer already exists"})
		return
	}

	// 4) Create the Customer record
	cust := &customerModels.Customer{
		FirstName:  payload["firstName"],
		LastName:   payload["lastName"],
		UserName:   payload["username"],
		Email:      payload["email"],
		Phone:      payload["phone"],
		Address:    payload["address"],
		City:       payload["city"],
		State:      payload["state"],
		PostalCode: payload["postalCode"],
		Country:    payload["country"],
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	res, err := customerRepo.CreateCustomer(cust)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create customer: " + err.Error()})
		return
	}

	// 5) Create the AuthCustomer record, linking to the new customer ID
	newID := res.InsertedID.(primitive.ObjectID)
	authCust := &authModels.AuthCustomer{
		Username:   payload["username"],
		Email:      payload["email"],
		Password:   payload["password"],
		CustomerID: newID,
	}
	if _, err := authRepo.CreateAuthCustomer(authCust); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create auth record: " + err.Error()})
		return
	}

	// 6) Return success
	c.JSON(http.StatusCreated, gin.H{
		"message":    "customer registered successfully",
		"customerId": newID.Hex(),
	})
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
