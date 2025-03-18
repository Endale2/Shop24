package models

// Seller represents a seller account.
type Seller struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	// additional seller-specific fields can be added here
}
