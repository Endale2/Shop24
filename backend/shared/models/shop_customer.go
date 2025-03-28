package models


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)



// Shop-Customer :  stores customer's data in the specific shop
type ShopCustomer struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	ShopID         primitive.ObjectID   `bson:"_id,omitempty"`
	CustomerID     primitive.ObjectID   `bson:"_id,omitempty"`
	CreatedAt      time.Time            `bson:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
	
}