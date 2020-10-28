package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"wedeal/context"
	graphqlSchema "wedeal/schema"

	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

// Init setups the server for the application
func Init(db *mongo.Database) {
	schema, err := graphqlSchema.CreateSchema()

	if err != nil {
		log.Fatalf("Failed to create new Schema: %v", err)
	}

	isDevEnv := os.Getenv("ENV") == "development"

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   isDevEnv,
		GraphiQL: isDevEnv,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Running on port %s\n", port)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.HandleRequest(w, r, graphqlHandler, db)
	})

	http.Handle("/graphql", handler)

	e := http.ListenAndServe(":"+port, nil)
	if e != nil {
		log.Fatalf("Failed to run http server: %v\n", e)
	}
}
