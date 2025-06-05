package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	sellerModels "github.com/Endale2/DRPS/sellers/models"
	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	"github.com/Endale2/DRPS/auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SellerRegisterService handles new seller sign-up.
// It creates a Seller record, links its ID into authSeller.SellerID,
// then creates the AuthSeller record.
func SellerRegisterService(authSeller *models.AuthSeller) error {
	// Validate required fields
	if authSeller.Username == "" || authSeller.Email == "" || authSeller.Password == "" {
		return errors.New("username, email and password are required")
	}

	// Check for existing authentication record
	if existing, _ := authRepo.FindAuthSellerByEmail(authSeller.Email); existing != nil {
		return errors.New("seller with that email already exists")
	}

	// Create the detailed Seller record
	newSeller := &sellerModels.Seller{
		// UserID can stay zero unless you track a separate user account
		
		Email:     authSeller.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	insertRes, err := sellerRepo.CreateSeller(newSeller)
	if err != nil {
		return err
	}
	newSeller.ID = insertRes.InsertedID.(primitive.ObjectID)

	// Link AuthSeller → Seller
	authSeller.SellerID = newSeller.ID

	// Insert the AuthSeller record
	_, err = authRepo.CreateAuthSeller(authSeller)
	return err
}

// SellerLoginService authenticates a seller and issues JWTs.
// Returns the full Seller record, an accessToken, a refreshToken, and error.
func SellerLoginService(authSeller *models.AuthSeller) (*sellerModels.Seller, string, string, error) {
	// 1) Lookup credentials
	foundAuth, err := authRepo.FindAuthSellerByEmail(authSeller.Email)
	if err != nil {
		return nil, "", "", errors.New("seller not found")
	}

	// 2) Verify password (in production use hashed compare)
	if authSeller.Password != foundAuth.Password {
		return nil, "", "", errors.New("invalid credentials")
	}

	// 3) Load full Seller profile
	sellerData, err := sellerRepo.GetSellerByID(foundAuth.SellerID)
	if err != nil {
		return nil, "", "", errors.New("failed to load seller data")
	}

	// 4) Issue JWT tokens
	// Access token: 5 minutes
	accessToken, err := utils.CreateToken(foundAuth.SellerID.Hex(), 5*time.Minute)
	if err != nil {
		return nil, "", "", errors.New("failed to generate access token")
	}
	// Refresh token: 7 days
	refreshToken, err := utils.CreateToken(foundAuth.SellerID.Hex(), 7*24*time.Hour)
	if err != nil {
		return nil, "", "", errors.New("failed to generate refresh token")
	}

	return sellerData, accessToken, refreshToken, nil
}
