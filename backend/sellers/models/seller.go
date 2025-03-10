package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Seller struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Name        string             `bson:"name"`
    Email       string             `bson:"email"`
    PhoneNumber string             `bson:"phone_number"`
    Products    []primitive.ObjectID `bson:"products"` // List of product IDs
}
