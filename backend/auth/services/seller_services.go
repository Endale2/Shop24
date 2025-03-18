package services

import (
	"errors"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	sellerModels "github.com/Endale2/DRPS/sellers/models"
	sellerRepo "github.com/Endale2/DRPS/sellers/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SellerRegisterService handles new seller registration by creating both a Seller record
// (detailed data) and an AuthSeller record (authentication data), linking them via SellerID.
func SellerRegisterService(authSeller *models.AuthSeller) error {
	// Validate required fields.
	if authSeller.Email == "" || authSeller.Password == "" || authSeller.Username == "" {
		return errors.New("missing required fields")
	}

	// Check if an authentication record already exists.
	if existing, _ := authRepo.FindAuthSellerByEmail(authSeller.Email); existing != nil {
		return errors.New("seller already exists")
	}

	// Create a new detailed Seller record.
	newSeller := &sellerModels.Seller{
		Name:  authSeller.Username, // Use username as default display name.
		Email: authSeller.Email,
		// Additional fields (Phone, Address, etc.) can be set later.
	}

	// Insert the Seller record.
	res, err := sellerRepo.CreateSeller(newSeller)
	if err != nil {
		return err
	}
	newSeller.ID = res.InsertedID.(primitive.ObjectID)

	// Link the AuthSeller record to the new Seller.
	authSeller.SellerID = newSeller.ID

	// Insert the AuthSeller record.
	_, err = authRepo.CreateAuthSeller(authSeller)
	return err
}

// SellerLoginService validates login credentials and retrieves both authentication data and the detailed Seller data.
func SellerLoginService(authSeller *models.AuthSeller) (*models.AuthSeller, *sellerModels.Seller, error) {
	// Look up the AuthSeller record by email.
	foundAuth, err := authRepo.FindAuthSellerByEmail(authSeller.Email)
	if err != nil {
		return nil, nil, errors.New("seller not found")
	}

	// Compare passwords (for simplicity, using plain text; update this with secure comparisons in production).
	if authSeller.Password != foundAuth.Password {
		return nil, nil, errors.New("password mismatch")
	}

	// Retrieve the detailed Seller record using the stored SellerID.
	sellerData, err := sellerRepo.GetSellerByID(foundAuth.SellerID)
	if err != nil {
		return nil, nil, errors.New("seller details not found")
	}

	return foundAuth, sellerData, nil
}
