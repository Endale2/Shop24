package services

import (
	"errors"
	"time"

	authModels "github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	customerModels "github.com/Endale2/DRPS/customers/models"
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	"github.com/Endale2/DRPS/auth/utils" // Ensure utils package has VerifyGoogleIDToken and CreateToken
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt" // Add bcrypt for local password hashing if not already used
)

// OptionalProfile holds non-required profile fields at registration.
type OptionalProfile struct {
	FirstName, LastName, Phone, Address, City, State, PostalCode, Country string
}

// CustomerRegisterService creates both Customer & AuthCustomer and returns full Customer.
func CustomerRegisterService(
	authCust *authModels.AuthCustomer,
	profile *OptionalProfile,
) (*customerModels.Customer, error) {
	// 1) Basic required
	if authCust.Username == "" || authCust.Email == "" || authCust.Password == "" {
		return nil, errors.New("username, email and password are required")
	}

	// 2) Dup check
	if existing, _ := authRepo.FindAuthCustomerByEmail(authCust.Email); existing != nil {
		return nil, errors.New("customer already exists")
	}

	// Hash password for local registration
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authCust.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	authCust.Password = string(hashedPassword)
	authCust.Provider = "local" // Set provider for local registration

	// 3) Build Customer record, merge optional profile
	cust := &customerModels.Customer{
		Email: authCust.Email,
		FirstName: profile.FirstName,
		LastName: profile.LastName,
		Phone: profile.Phone,
		Address: profile.Address,
		City: profile.City,
		State: profile.State,
		PostalCode: profile.PostalCode,
		Country: profile.Country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := customerRepo.CreateCustomer(cust)
	if err != nil {
		return nil, err
	}

	// 4) Link AuthCustomer → Customer
	oid := res.InsertedID.(primitive.ObjectID)
	authCust.CustomerID = oid
	if _, err := authRepo.CreateAuthCustomer(authCust); err != nil {
		return nil, err
	}

	// Return the complete customer
	cust.ID = oid
	return cust, nil
}

// CustomerLoginService authenticates a customer and returns Customer + tokens.
func CustomerLoginService(email, password string) (*customerModels.Customer, string, string, error) {
	// 1) Lookup
	foundAuth, err := authRepo.FindAuthCustomerByEmail(email)
	if err != nil || foundAuth == nil { // Handle nil foundAuth explicitly
		return nil, "", "", errors.New("invalid credentials")
	}

	// 2) Password check using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(foundAuth.Password), []byte(password)); err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// 3) Load Customer
	cust, err := customerRepo.GetCustomerByID(foundAuth.CustomerID)
	if err != nil || cust == nil {
		return nil, "", "", errors.New("failed to load customer data")
	}

	// 4) Tokens
	at, err := utils.CreateToken(foundAuth.CustomerID.Hex(), 5*time.Minute)
	if err != nil {
		return nil, "", "", errors.New("failed to generate access token")
	}
	rt, err := utils.CreateToken(foundAuth.CustomerID.Hex(), 7*24*time.Hour)
	if err != nil {
		return nil, "", "", errors.New("failed to generate refresh token")
	}
	return cust, at, rt, nil
}

// CustomerLoginOAuth handles Google (or other) OAuth login for customers.
func CustomerLoginOAuth(provider, idToken string) (*customerModels.Customer, string, string, error) {
	payload, err := utils.VerifyGoogleIDToken(idToken)
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