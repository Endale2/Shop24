package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	sellerModels "github.com/Endale2/DRPS/sellers/models"
	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	"github.com/Endale2/DRPS/auth/utils"
)

// SellerLoginService validates login credentials, returns seller data and tokens.
func SellerLoginService(authSeller *models.AuthSeller) (*sellerModels.Seller, string, string, error) {
	// Find auth record
	foundAuth, err := authRepo.FindAuthSellerByEmail(authSeller.Email)
	if err != nil {
		return nil, "", "", errors.New("seller not found")
	}

	if authSeller.Password != foundAuth.Password {
		return nil, "", "", errors.New("password mismatch")
	}

	sellerData, err := sellerRepo.GetSellerByID(foundAuth.SellerID)
	if err != nil {
		return nil, "", "", errors.New("seller details not found")
	}

	// Generate tokens
	accessToken, err := utils.CreateToken(foundAuth.SellerID.Hex(), 5*time.Minute)
	if err != nil {
		return nil, "", "", errors.New("failed to generate access token")
	}

	refreshToken, err := utils.CreateToken(foundAuth.SellerID.Hex(), 7*24*time.Hour)
	if err != nil {
		return nil, "", "", errors.New("failed to generate refresh token")
	}

	return sellerData, accessToken, refreshToken, nil
}
