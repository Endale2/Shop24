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

// SellerLoginService validates credentials, loads seller details,
// and issues access & refresh JWTs.
func SellerLoginService(authSeller *models.AuthSeller) (*models.AuthSeller, *sellerModels.Seller, error) {
	// 1) Find AuthSeller record
	foundAuth, err := authRepo.FindAuthSellerByEmail(authSeller.Email)
	if err != nil {
		return nil, nil, errors.New("seller not found")
	}
	// 2) Compare passwords
	if authSeller.Password != foundAuth.Password {
		return nil, nil, errors.New("password mismatch")
	}
	// 3) Load full Seller record
	sellerData, err := sellerRepo.GetSellerByID(foundAuth.SellerID)
	if err != nil {
		return nil, nil, errors.New("seller details not found")
	}
	// 4) Issue tokens via utils
	access, _ := utils.CreateToken(foundAuth.SellerID.Hex(), 5*time.Minute)
	refresh, _ := utils.CreateToken(foundAuth.SellerID.Hex(), 7*24*time.Hour)
	foundAuth.AccessToken = access
	foundAuth.RefreshToken = refresh

	return foundAuth, sellerData, nil
}
