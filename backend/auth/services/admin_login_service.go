package services

import (
	"errors"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	adminModels "github.com/Endale2/DRPS/admin/models"
	adminRepo "github.com/Endale2/DRPS/admin/repositories"
	
)



// AdminLoginService validates login credentials and returns both authentication and detailed admin data.
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

	return foundAuth, adminData, nil
}
