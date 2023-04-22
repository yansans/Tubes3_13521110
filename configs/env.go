package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("user_db")
	pass := os.Getenv("pass_db")
	uri := "mongodb+srv://" + user + ":" + pass + "@chatbot-cluster.4axocdy.mongodb.net/?retryWrites=true&w=majority"
	return uri
}
