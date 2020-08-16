package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"wedeal/modules/user"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect handles the connection with database
func Connect() {
	uri := os.Getenv("MONGO_DB_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error checking the database connection: %v\n", err)
	}

	fmt.Println("Connected to database!")

	dbName := os.Getenv("DB_NAME")
	db := client.Database(dbName)

	user.SetupCollection(db)
}
