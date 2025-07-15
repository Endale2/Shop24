package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	"github.com/Endale2/DRPS/auth/utils"
	sellerModels "github.com/Endale2/DRPS/sellers/models"
	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrSellerNotFound = errors.New("seller not found")

// SellerLoginOAuth handles Google (or other) OAuth login.
func SellerLoginOAuth(provider, idToken string) (*sellerModels.Seller, string, string, error) {
	// Use the correct client ID for sellers
	payload, err := utils.VerifyGoogleIDToken(idToken, utils.GoogleSellerClientID())
	if err != nil {
		return nil, "", "", err
	}
	rec, _ := authRepo.FindAuthSellerByProvider(provider, payload.Sub)
	if rec == nil {
		// first‐time: create profile + auth
		now := time.Now()
		prof := &sellerModels.Seller{
			FirstName:    payload.GivenName,
			LastName:     payload.FamilyName,
			Email:        payload.Email,
			ProfileImage: payload.Picture,
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		res, err := sellerRepo.CreateSeller(prof)
		if err != nil {
			return nil, "", "", err
		}
		prof.ID = res.InsertedID.(primitive.ObjectID)
		rec = &models.AuthSeller{
			Email:      payload.Email,
			Provider:   provider,
			ProviderID: payload.Sub,
			SellerID:   prof.ID,
		}
		if _, err := authRepo.CreateAuthSeller(rec); err != nil {
			return nil, "", "", err
		}
	}
	prof, err := sellerRepo.GetSellerByID(rec.SellerID)
	if err != nil || prof == nil {
		return nil, "", "", ErrSellerNotFound
	}
	at, _ := utils.CreateToken(prof.ID.Hex(), 5*time.Minute)
	rt, _ := utils.CreateToken(prof.ID.Hex(), 7*24*time.Hour)
	return prof, at, rt, nil
}
