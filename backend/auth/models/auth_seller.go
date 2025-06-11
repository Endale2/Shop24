package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthSeller holds both local and OAuth credentials.
type AuthSeller struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"        json:"id"`
    Email      string             `bson:"email,omitempty"      json:"email,omitempty"`
    Password   string             `bson:"password,omitempty"   json:"-"`
    Provider   string             `bson:"provider,omitempty"   json:"provider,omitempty"`   // "local" or "google"
    ProviderID string             `bson:"provider_id,omitempty" json:"providerId,omitempty"` // OAuth sub
    SellerID   primitive.ObjectID `bson:"seller_id,omitempty"  json:"seller_id,omitempty"`
}
