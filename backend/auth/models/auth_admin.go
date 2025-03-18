package models


import "go.mongodb.org/mongo-driver/bson/primitive"

// Admin represents an administrator account.
type AuthAdmin struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	AdminID  primitive.ObjectID `bson:"admin_id,omitempty" json:"admin_id"` // Reference to Admin model
}
