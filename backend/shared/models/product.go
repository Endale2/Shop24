package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Variant struct {
	Color     string `bson:"color"`
	Size      string `bson:"size"`
	Capacity  string `bson:"capacity"`
	Price     float64 `bson:"price"`
	Stock     int     `bson:"stock"`
	Image     string  `bson:"image"`
}

type Product struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	Description   string             `bson:"description"`
	Images        []string           `bson:"images"` 
	Category      string             `bson:"category"`
	Source        string             `bson:"source"` // 'user' or 'aliexpress'
	CreatedBy     primitive.ObjectID `bson:"createdBy,omitempty"` // If the product was added by a user
	Variants      []Variant          `bson:"variants"`
	AliExpressData AliExpressData    `bson:"aliexpressData,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt     time.Time          `bson:"updatedAt,omitempty"`
}

type AliExpressData struct {
	ProductID   string `bson:"productId"`
	SellerName  string `bson:"sellerName"`
	AliUrl      string `bson:"aliUrl"`
}
