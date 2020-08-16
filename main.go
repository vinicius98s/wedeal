package main

import (
	"log"
	"wedeal/database"
	"wedeal/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	database.Connect()
	router.SetupRouter()
}
