package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/DPRS/models"
	"github.com/DPRS/repositories/auth"
)

// AuthController handles authentication-related endpoints.
type AuthController struct {
	Repo auth.AuthRepository
}

// NewAuthController creates a new instance of AuthController.
func NewAuthController(repo auth.AuthRepository) *AuthController {
	return &AuthController{
		Repo: repo,
	}
}

// RegisterSeller handles registration of a new seller.
func (ac *AuthController) RegisterSeller(c *gin.Context) {
	var seller models.AuthSeller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// TODO: Hash the password before saving.
	// seller.Password = HashPassword(seller.Password)

	ctx := context.Background()
	if err := ac.Repo.CreateSeller(ctx, &seller); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating seller"})
		return
	}

	c.JSON(http.StatusCreated, seller)
}

// LoginSeller handles seller login.
func (ac *AuthController) LoginSeller(c *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ctx := context.Background()
	seller, err := ac.Repo.FindSellerByEmail(ctx, creds.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Seller not found"})
		return
	}

	// TODO: Validate the provided password with the stored hash.
	// if !ValidatePassword(creds.Password, seller.Password) {
	//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	//     return
	// }

	// Update the last login time.
	if err := ac.Repo.UpdateLastLogin(ctx, seller.ID, time.Now()); err != nil {
		// Optionally log the error and continue.
	}

	// TODO: Generate a JWT token or session if needed before responding.
	c.JSON(http.StatusOK, seller)
}
