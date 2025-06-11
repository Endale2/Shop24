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
	"golang.org/x/crypto/bcrypt"
)

var ErrSellerNotFound = errors.New("seller not found")

// SellerRegisterService wraps the local registration (hashing + profile creation).
func SellerRegisterService(auth *models.AuthSeller) error {
	if auth.Email == "" || auth.Password == "" {
		return errors.New("email and password required")
	}
	if ex, _ := authRepo.FindAuthSellerByEmail(auth.Email); ex != nil {
		return errors.New("email already in use")
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	auth.Password = string(hash)
	auth.Provider = "local"
	// create seller profile
	now := time.Now()
	prof := &sellerModels.Seller{
		Email:     auth.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
	res, err := sellerRepo.CreateSeller(prof)
	if err != nil {
		return err
	}
	prof.ID = res.InsertedID.(primitive.ObjectID)
	auth.SellerID = prof.ID
	_, err = authRepo.CreateAuthSeller(auth)
	return err
}

// SellerLoginService logs in a local seller.
func SellerLoginService(auth *models.AuthSeller) (*sellerModels.Seller, string, string, error) {
	rec, _ := authRepo.FindAuthSellerByEmail(auth.Email)
	if rec == nil {
		return nil, "", "", errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(rec.Password), []byte(auth.Password)); err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}
	prof, err := sellerRepo.GetSellerByID(rec.SellerID)
	if err != nil || prof == nil {
		return nil, "", "", ErrSellerNotFound
	}
	at, _ := utils.CreateToken(prof.ID.Hex(), 5*time.Minute)
	rt, _ := utils.CreateToken(prof.ID.Hex(), 7*24*time.Hour)
	return prof, at, rt, nil
}

// SellerLoginOAuth handles Google (or other) OAuth login.
func SellerLoginOAuth(provider, idToken string) (*sellerModels.Seller, string, string, error) {
	payload, err := utils.VerifyGoogleIDToken(idToken)
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
