package schema

import (
	"github.com/graphql-go/graphql"
)

// CreateSchema generates the graphql schema necessary for application
func CreateSchema() (graphql.Schema, error) {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: RootQuery}
	mutation := graphql.ObjectConfig{Name: "Mutation", Fields: Mutation}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(mutation)}
	schema, err := graphql.NewSchema(schemaConfig)

	return schema, err
}
