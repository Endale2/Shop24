package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

// Admin holds full profile data for an administrator.
type Admin struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    FirstName  string             `bson:"first_name"    json:"firstName"`
    LastName   string             `bson:"last_name"     json:"lastName"`
    Email      string             `bson:"email"         json:"email"`
    Phone      string             `bson:"phone,omitempty" json:"phone,omitempty"`
    Address    string             `bson:"address,omitempty" json:"address,omitempty"`
    Role       string             `bson:"role"          json:"role"` // e.g. "admin"
    CreatedAt  time.Time          `bson:"created_at"    json:"createdAt"`
    UpdatedAt  time.Time          `bson:"updated_at"    json:"updatedAt"`
}
