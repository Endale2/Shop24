package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthCustomer represents the authentication details for a customer.
type AuthCustomer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string             `bson:"username" json:"username" binding:"required"`
	Email      string             `bson:"email" json:"email" binding:"required"`
	Password   string             `bson:"password" json:"password" binding:"required"`
	CustomerID primitive.ObjectID `bson:"customer_id,omitempty" json:"customer_id"` // Reference to Customer model
}
