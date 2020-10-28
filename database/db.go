package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect handles the connection with database
func Connect() (*mongo.Database, error) {
	uri := os.Getenv("MONGO_DB_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database!")

	dbName := os.Getenv("DB_NAME")
	db := client.Database(dbName)
	return db, nil
}
