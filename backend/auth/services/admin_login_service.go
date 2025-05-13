package services

import (
	"errors"
	"time"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	adminModels "github.com/Endale2/DRPS/admin/models"
	adminRepo "github.com/Endale2/DRPS/admin/repositories"
	"github.com/Endale2/DRPS/auth/utils"
)

// AdminLoginService authenticates an admin and issues tokens.
func AdminLoginService(req *models.AuthAdmin) (*models.AuthAdmin, *adminModels.Admin, error) {
	// 1. Retrieve credential record
	cred, err := authRepo.FindAuthAdminByEmail(req.Email)
	if err != nil {
		return nil, nil, errors.New("admin not found")
	}
	// 2. Validate password
	if req.Password != cred.Password {
		return nil, nil, errors.New("password mismatch")
	}
	// 3. Fetch admin profile
	adminData, err := adminRepo.GetAdminByID(cred.AdminID)
	if err != nil {
		return nil, nil, errors.New("admin details not found")
	}
	// 4. Issue tokens
	accessToken, err := utils.CreateToken(cred.AdminID.Hex(), 5*time.Minute)
	if err != nil {
		return nil, nil, errors.New("could not create access token")
	}
	refreshToken, err := utils.CreateToken(cred.AdminID.Hex(), 7*24*time.Hour)
	if err != nil {
		return nil, nil, errors.New("could not create refresh token")
	}
	// 5. Attach tokens
	cred.AccessToken = accessToken
	cred.RefreshToken = refreshToken
	return cred, adminData, nil
}
