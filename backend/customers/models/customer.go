package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Customer represents the detailed data for a customer.
type Customer struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Email   string             `bson:"email" json:"email"`
	Phone   string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Address string             `bson:"address,omitempty" json:"address,omitempty"`
	
}
