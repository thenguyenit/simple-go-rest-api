package main

import (
	"log"
	"net/http"
	"os"

	"github.com/thenguyenit/simple-go-rest-api/store"

	"github.com/joho/godotenv"
)

func main() {

	//Read the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the "PORT" env variable
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := store.NewRouter()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
