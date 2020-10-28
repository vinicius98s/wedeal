package main

import (
	"log"
	"wedeal/database"
	"wedeal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	db, err := database.Connect()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	server.Init(db)
}
