package models

import (
	

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type CartItem struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID  primitive.ObjectID `bson:"product_id" json:"product_id"`
	VariantID  primitive.ObjectID `bson:"variant_id,omitempty" json:"variant_id,omitempty"`
	ProductName string            `bson:"product_name" json:"product_name"`   // snapshot
	VariantOptions map[string]string `bson:"variant_options" json:"variant_options"` // snapshot
	Image      string              `bson:"image,omitempty" json:"image,omitempty"`   // primary image or variant image

	UnitPrice  float64           `bson:"unit_price" json:"unit_price"` // pre-discount
	Quantity   int               `bson:"quantity" json:"quantity"`
	LineTotal  float64           `bson:"line_total" json:"line_total"` // UnitPrice * Quantity (before discounts)

	DiscountAmount float64       `bson:"discount_amount,omitempty" json:"discount_amount,omitempty"` // optional line discount
	FinalLineTotal float64       `bson:"final_line_total" json:"final_line_total"` // LineTotal - DiscountAmount

	AppliedDiscountIDs []primitive.ObjectID `bson:"applied_discount_ids,omitempty" json:"applied_discount_ids,omitempty"`
}
