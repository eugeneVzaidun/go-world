package main

import (
	"context"
	"encoding/json"
	"go-world/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
// init function that starts the server
func init() {
	// println is a function that prints to the console
	println("Starting server...")
	var connectionError error
	// init connects to the mongodb server
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, connectionError = mongo.Connect(context.TODO(), clientOptions)
	// print message if connection to mongo is successful
	println("Connected to MongoDB!")

	if connectionError != nil {
		log.Fatal(connectionError)
	}
}



// function that stores user in monogo db and return the user
func storeUser(user models.User) models.User {
	// connect to the database
	db := client.Database("test")
	// create a collection
	collection := db.Collection("users")
	var err error
	// insert the user
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	// return the user

	return user
}

// function that handles post requests to the server and create users
func handlePost(w http.ResponseWriter, r *http.Request) {
	// user logger to log the request
	println("POST request received")

	// decode the json body
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	storeUser(user)
	// write the user to the response
	json, _ := json.Marshal(user)
	w.Write(json)
}

// handler that rertieves the user from the database
func handleGet(w http.ResponseWriter, r *http.Request) {
	// user logger to log the request
	println("GET request received")
	// connect to the database
	db := client.Database("test")
	// create a collection
	collection := db.Collection("users")
	// get the user from the database
	var user models.User
	collection.FindOne(context.TODO(), bson.D{{"name","mayank"}}).Decode(&user)
	// write the user to the response
	json, _ := json.Marshal(user)
	w.Write(json)
}

// handler that retrivies all documents from the database
func handleGetAll(w http.ResponseWriter, r *http.Request) {
	// user logger to log the request
	println("GET request received")
	// connect to the database
	db := client.Database("test")
	// create a collection
	collection := db.Collection("users")
	// get the user from the database
	var users []models.User
	opts := options.Find().SetSort(bson.D{})
	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}
	// write the user to the response
	json, _ := json.Marshal(users)
	w.Write(json)
}

func startServer() {
	http.HandleFunc("/user", handlePost)
	http.HandleFunc("/users/get", handleGet)
	http.HandleFunc("/users", handleGetAll)
	http.ListenAndServe(":8080", nil)
}

// function that starts the server
func main() {
	startServer()
	defer client.Disconnect(context.TODO())
}