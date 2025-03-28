package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	adminModels "github.com/Endale2/DRPS/admin/models"
	adminRepo "github.com/Endale2/DRPS/admin/repositories"
	"github.com/golang-jwt/jwt/v4"
)

// jwtSecret is the secret key used to sign tokens.
// In production, load this from a secure configuration.
var jwtSecret = []byte("your_secret_key")

// AdminLoginService validates login credentials, generates JWT access and refresh tokens,
// and returns both authentication data (with tokens) and detailed admin data.
func AdminLoginService(authAdmin *models.AuthAdmin) (*models.AuthAdmin, *adminModels.Admin, error) {
	// Look up the AuthAdmin record by email.
	foundAuth, err := authRepo.FindAuthAdminByEmail(authAdmin.Email)
	if err != nil {
		return nil, nil, errors.New("admin not found")
	}

	// Compare passwords (in production, use a secure hash comparison).
	if authAdmin.Password != foundAuth.Password {
		return nil, nil, errors.New("password mismatch")
	}

	// Retrieve the detailed admin record using the stored AdminID.
	adminData, err := adminRepo.GetAdminByID(foundAuth.AdminID)
	if err != nil {
		return nil, nil, errors.New("admin details not found")
	}

	// Generate Access Token (expires in 15 minutes)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": foundAuth.AdminID,
		"email":    foundAuth.Email,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, nil, errors.New("failed to generate access token")
	}

	// Generate Refresh Token (expires in 7 days)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": foundAuth.AdminID,
		"email":    foundAuth.Email,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, nil, errors.New("failed to generate refresh token")
	}

	// Attach tokens to the auth response.
	// Note: These fields should be part of your models.AuthAdmin struct.
	foundAuth.AccessToken = accessTokenString
	foundAuth.RefreshToken = refreshTokenString

	return foundAuth, adminData, nil
}
