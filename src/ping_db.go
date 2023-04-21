package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func get_secret(secret_file string) (string, string) {
	// Read the secret file
	// The secret file should contain the username and password
	// separated by a newline
	// Example:
	// username
	// password
	// The secret file should be in the same directory as the executable
	// or in the same directory as the source code
	// The secret file should not be uploaded to GitHub

	// Read the secret file
	secret, err := os.ReadFile(secret_file)
	if err != nil {
		panic(err)
	}

	// Split the secret file into lines
	lines := strings.Split(string(secret), "\n")

	// Get the username and password and remove the newline character
	user := lines[0]
	user = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(user, "")
	pass := lines[1]
	pass = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(pass, "")

	return user, pass
}

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Set the credentials
	secret_file := "secret.txt"
	user, pass := get_secret(secret_file)

	uri := "mongodb+srv://" + user + ":" + pass + "@chatbot-cluster.4axocdy.mongodb.net/?retryWrites=true&w=majority"

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(
		context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
