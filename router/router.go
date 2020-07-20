package router

import (
	"fmt"
	"log"
	"net/http"
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

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	fmt.Println("Running on port 8080")

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
