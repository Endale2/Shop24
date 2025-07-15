package services

import (
	"errors"
	"time"

	authModels "github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	"github.com/Endale2/DRPS/auth/utils" // Ensure utils package has VerifyGoogleIDToken and CreateToken
	customerModels "github.com/Endale2/DRPS/customers/models"
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OptionalProfile holds non-required profile fields at registration.
type OptionalProfile struct {
	FirstName, LastName, Phone, Address, City, State, PostalCode, Country string
}

// CustomerLoginOAuth handles Google (or other) OAuth login for customers.
func CustomerLoginOAuth(provider, idToken string) (*customerModels.Customer, string, string, error) {
	// Use the correct client ID for customers
	payload, err := utils.VerifyGoogleIDToken(idToken, utils.GoogleCustomerClientID())
	if err != nil {
		return nil, "", "", err
	}

	rec, _ := authRepo.FindAuthCustomerByProvider(provider, payload.Sub)
	if rec == nil {
		// First-time login: create customer profile + auth entry
		now := time.Now()
		cust := &customerModels.Customer{
			FirstName:    payload.GivenName,
			LastName:     payload.FamilyName,
			Email:        payload.Email,
			ProfileImage: payload.Picture,
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		res, err := customerRepo.CreateCustomer(cust)
		if err != nil {
			return nil, "", "", err
		}
		cust.ID = res.InsertedID.(primitive.ObjectID)

		rec = &authModels.AuthCustomer{
			Email:      payload.Email,
			Username:   payload.Email, // Consistent with customer profile
			Provider:   provider,
			ProviderID: payload.Sub,
			CustomerID: cust.ID,
			Password:   "", // OAuth users don't have a local password
		}
		if _, err := authRepo.CreateAuthCustomer(rec); err != nil {
			return nil, "", "", err
		}
	}

	// Retrieve the customer profile
	cust, err := customerRepo.GetCustomerByID(rec.CustomerID)
	if err != nil || cust == nil {
		return nil, "", "", errors.New("customer profile not found")
	}

	// Generate tokens
	at, err := utils.CreateToken(cust.ID.Hex(), 5*time.Minute)
	if err != nil {
		return nil, "", "", errors.New("failed to generate access token")
	}
	rt, err := utils.CreateToken(cust.ID.Hex(), 7*24*time.Hour)
	if err != nil {
		return nil, "", "", errors.New("failed to generate refresh token")
	}

	return cust, at, rt, nil
}

// TODO: Add CustomerLoginOAuth logic for Telegram here
