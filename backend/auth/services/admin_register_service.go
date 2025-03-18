package services

import (
	"errors"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/repositories"
)


// AdminRegisterService handles new admin registration.
func AdminRegisterService(admin *models.Admin) error {
	// Validate required fields.
	if admin.Email == "" || admin.Password == "" || admin.Username == "" {
		return errors.New("missing required fields")
	}

	// Optional: Check if the admin already exists.
	if existingAdmin, err := repositories.FindAdminByEmail(admin.Email); err == nil && existingAdmin != nil {
		return errors.New("admin already exists")
	}

	// Insert the new admin.
	return repositories.InsertAdmin(admin)
}