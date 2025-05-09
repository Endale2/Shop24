package services

import (
	"errors"

	"github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	adminModels "github.com/Endale2/DRPS/admin/models"
	adminRepo "github.com/Endale2/DRPS/admin/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)




func AdminRegisterService(authAdmin *models.AuthAdmin) error {
	// Validate required fields.
	if authAdmin.Email == "" || authAdmin.Password == "" || authAdmin.Username == "" {
		return errors.New("missing required fields")
	}

	// Check if an authentication record already exists.
	if existing, _ := authRepo.FindAuthAdminByEmail(authAdmin.Email); existing != nil {
		return errors.New("admin already exists")
	}

	// Create a new detailed admin record.
	newAdmin := &adminModels.Admin{
		Name:    authAdmin.Username, // Use the username as the name by default.
		Phone:   "",
		Address: "",
		Role:    "admin",
	}

	// Insert the new admin into the Admin collection.
	res, err := adminRepo.CreateAdmin(newAdmin)
	if err != nil {
		return err
	}
	// Set the new admin's ID.
	newAdmin.ID = res.InsertedID.(primitive.ObjectID)
	// Link the AuthAdmin record to the new Admin.
	authAdmin.AdminID = newAdmin.ID

	// Create the AuthAdmin record.
	_, err = authRepo.CreateAuthAdmin(authAdmin)
	return err
}