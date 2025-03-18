package services

import (
	"errors"

	"github.com/Endale2/DRPS/auth/models"
	"github.com/Endale2/DRPS/auth/repositories"
)

// AdminLoginService validates admin credentials.
func AdminLoginService(admin *models.Admin) error {
	foundAdmin, err := repositories.FindAdminByEmail(admin.Email)
	if err != nil {
		return err
	}
	// For demo purposes, a plain text password comparison is shown.
	// In production, use hashed passwords and secure comparisons.
	if admin.Password != foundAdmin.Password {
		return errors.New("password mismatch")
	}
	return nil
}
