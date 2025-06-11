package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthAdmin links credentials → profile
type AuthAdmin struct {
    ID           primitive.ObjectID `bson:"_id,omitempty"    json:"id"`
    Email        string             `bson:"email"            json:"email"`
    PasswordHash string             `bson:"password_hash"    json:"-"`
    Provider     string             `bson:"provider"         json:"provider"`    // "local" or "google"
    ProviderID   string             `bson:"provider_id"      json:"providerId"`  // e.g. Google sub
    AdminID      primitive.ObjectID `bson:"admin_id"         json:"adminId"`
    AccessToken  string             `bson:"-"                json:"accessToken,omitempty"`
    RefreshToken string             `bson:"-"                json:"refreshToken,omitempty"`
}
