package models // in customers/models

import "go.mongodb.org/mongo-driver/bson/primitive"
import "time"

// Customer is the global customer record.
type Customer struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    FirstName  string             `bson:"firstName" json:"firstName"`
    LastName   string             `bson:"lastName" json:"lastName"`
	UserName   string             `bson:"username" json:"username"`
    Email      string             `bson:"email" json:"email"`
    Phone      string             `bson:"phone" json:"phone"`
    Address    string             `bson:"address" json:"address"`
    City       string             `bson:"city" json:"city"`
    State      string             `bson:"state" json:"state"`
    PostalCode string             `bson:"postalCode" json:"postalCode"`
    Country    string             `bson:"country" json:"country"`
    CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
    UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}