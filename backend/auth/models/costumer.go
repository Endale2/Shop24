package models

// Customer represents a customer account.
type Customer struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	// additional customer-specific fields can be added here
}
