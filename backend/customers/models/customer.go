package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Customer struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Email   string             `bson:"email" json:"email"`
	Phone   string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Address string             `bson:"address,omitempty" json:"address,omitempty"`
	
}
