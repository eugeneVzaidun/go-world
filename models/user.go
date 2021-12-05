package models

// define user model with json tags
type User struct {
	// define mongo db id
	ID string `json:"id" bson:"_id"`
	Name string `json:"name"`
	Age int `json:"age"`
}