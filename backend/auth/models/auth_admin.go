package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthAdmin represents administrator credentials and issued tokens.
type AuthAdmin struct {
	ID           string             `json:"id" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	AdminID      primitive.ObjectID `json:"admin_id" bson:"admin_id"`

	// Tokens issued on successful login
	AccessToken  string             `json:"access_token,omitempty" bson:"-"`
	RefreshToken string             `json:"refresh_token,omitempty" bson:"-"`
}