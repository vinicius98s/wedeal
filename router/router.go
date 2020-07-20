package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wedeal/schema"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// SetupRouter : setup the routing
func SetupRouter() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	schema, err := schema.CreateSchema(fields)

	if err != nil {
		log.Fatalf("failed to create new Schema: %v", err)
	}

	isDevEnv := os.Getenv("ENV") == "development"

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   isDevEnv,
		GraphiQL: isDevEnv,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Running on port %s\n", port)

	http.Handle("/graphql", h)
	http.ListenAndServe(":"+port, nil)
}
