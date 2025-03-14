package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthUser represents the minimal set of user information needed for authentication.
type AuthSeller struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email        string             `bson:"email" json:"email"`                     // Unique email address
	Password string             `bson:"password_hash" json:"-"`                   // Securely hashed password; not exposed in JSON responses
	Roles        []string           `bson:"roles" json:"roles"`                       // E.g., ["seller"], ["buyer"], ["admin"]
	MFAEnabled   bool               `bson:"mfa_enabled" json:"mfa_enabled"`           // Indicates if multi-factor authentication is enabled
	MFASecret    string             `bson:"mfa_secret,omitempty" json:"-"`            // MFA secret for TOTP; kept hidden
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	LastLogin    *time.Time         `bson:"last_login,omitempty" json:"last_login,omitempty"`
}
