package models

// define user model with json tags
type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}