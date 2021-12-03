package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// init function that starts the server
func init() {
	// println is a function that prints to the console
	println("Starting server...")
	
}
// define user model with json tags
type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

// function that stores user in monogo db and return the user
func storeUser(user User) User {
	// store user in mongo db
	// connect to mongo db
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// connect to the database
	db := client.Database("test")
	// create a collection
	collection := db.Collection("users")
	// insert the user
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	// return the user

	return user
}

func handler(w http.ResponseWriter, r *http.Request) {
	// create a new user
	user := User{
		Name: "John",
		Age: 30,
	}
	// encode the user to json
	json, _ := json.Marshal(user)
	// write the json to the response
	w.Write(json)
}
// function that handles post requests to the server and create users
func handlePost(w http.ResponseWriter, r *http.Request) {
	// user logger to log the request
	println("POST request received")

	// decode the json body
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	storeUser(user)
	// write the user to the response
	json, _ := json.Marshal(user)
	w.Write(json)
}

func startServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/post", handlePost)
	http.ListenAndServe(":8080", nil)
}

// function that starts the server
func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Connected to MongoDB!")
	
	// start the server
	startServer()
	defer client.Disconnect(context.TODO())
}
