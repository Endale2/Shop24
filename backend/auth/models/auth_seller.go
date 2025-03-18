package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthSeller represents the authentication record for a seller.
type AuthSeller struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	// SellerID links to the detailed Seller record in sellers/models/seller.go.
	SellerID primitive.ObjectID `bson:"seller_id,omitempty" json:"seller_id"`
}
