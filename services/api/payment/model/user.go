package model

// User data entity for users
type User struct {
	Name    string                 `json:"name"`
	Email   string                 `json:"email"`
	Payment map[string]interface{} `json:"payment" bson:"payment"`
}
