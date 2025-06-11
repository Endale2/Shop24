package services

import (
	"errors"
	"time"

	adminModels "github.com/Endale2/DRPS/admin/models"
	adminRepo "github.com/Endale2/DRPS/admin/repositories"
	authModels "github.com/Endale2/DRPS/auth/models"
	authRepo "github.com/Endale2/DRPS/auth/repositories"
	"github.com/Endale2/DRPS/auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterLocalAdmin creates both a profile + local-credential record.
func RegisterLocalAdmin(email, password, firstName, lastName string) error {
    if email == "" || password == "" {
        return errors.New("email and password required")
    }
    if _, err := authRepo.FindAuthAdminByEmail(email); err == nil {
        return errors.New("admin already exists")
    }
    // 1) Hash password
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    // 2) Create profile
    now := time.Now()
    prof := &adminModels.Admin{
        FirstName: firstName,
        LastName:  lastName,
        Email:     email,
        Role:      "admin",
        CreatedAt: now,
        UpdatedAt: now,
    }
    res, err := adminRepo.CreateAdmin(prof)
    if err != nil {
        return err
    }
    prof.ID = res.InsertedID.(primitive.ObjectID)
    // 3) Link auth
    cred := &authModels.AuthAdmin{
        Email:        email,
        PasswordHash: string(hash),
        Provider:     "local",
        ProviderID:   "",
        AdminID:      prof.ID,
    }
    _, err = authRepo.CreateAuthAdmin(cred)
    return err
}

// LoginLocalAdmin checks credentials and issues tokens.
func LoginLocalAdmin(email, password string) (*authModels.AuthAdmin, *adminModels.Admin, error) {
    cred, err := authRepo.FindAuthAdminByEmail(email)
    if err != nil {
        return nil, nil, errors.New("invalid credentials")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(cred.PasswordHash), []byte(password)); err != nil {
        return nil, nil, errors.New("invalid credentials")
    }
    // fetch profile
    prof, err := adminRepo.GetAdminByID(cred.AdminID)
    if err != nil {
        return nil, nil, err
    }
    // issue tokens
    at, err := utils.CreateToken(cred.AdminID.Hex(), 5*time.Minute)
    if err != nil {
        return nil, nil, err
    }
    rt, err := utils.CreateToken(cred.AdminID.Hex(), 7*24*time.Hour)
    if err != nil {
        return nil, nil, err
    }
    cred.AccessToken, cred.RefreshToken = at, rt
    return cred, prof, nil
}

// LoginOAuthAdmin finds or provisions a social‐login admin.
func LoginOAuthAdmin(provider, providerID, email, firstName, lastName string) (*authModels.AuthAdmin, *adminModels.Admin, error) {
    // 1) Try existing link
    cred, err := authRepo.FindAuthAdminByProvider(provider, providerID)
    if err == mongo.ErrNoDocuments {
        // None → try by email
        cred, _ = authRepo.FindAuthAdminByEmail(email)
    }
    var prof *adminModels.Admin
    if cred == nil {
        // provision new profile
        now := time.Now()
        p := &adminModels.Admin{
            FirstName: firstName,
            LastName:  lastName,
            Email:     email,
            Role:      "admin",
            CreatedAt: now,
            UpdatedAt: now,
        }
        res, err := adminRepo.CreateAdmin(p)
        if err != nil {
            return nil, nil, err
        }
        p.ID = res.InsertedID.(primitive.ObjectID)

        cred = &authModels.AuthAdmin{
            Email:      email,
            Provider:   provider,
            ProviderID: providerID,
            AdminID:    p.ID,
        }
        if _, err := authRepo.CreateAuthAdmin(cred); err != nil {
            return nil, nil, err
        }
        prof = p
    } else {
        // existing → load profile
        prof, err = adminRepo.GetAdminByID(cred.AdminID)
        if err != nil {
            return nil, nil, err
        }
    }
    // issue tokens
    at, _ := utils.CreateToken(cred.AdminID.Hex(), 5*time.Minute)
    rt, _ := utils.CreateToken(cred.AdminID.Hex(), 7*24*time.Hour)
    cred.AccessToken, cred.RefreshToken = at, rt
    return cred, prof, nil
}
