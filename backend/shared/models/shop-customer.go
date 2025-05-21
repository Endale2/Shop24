package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)



// ShopCustomer links a Customer to a Shop (many-to-many / membership).
type ShopCustomer struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    ShopID     primitive.ObjectID `bson:"shopId" json:"shopId"`
    CustomerID primitive.ObjectID `bson:"customerId" json:"customerId"`
    CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
    UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}
