package schema

import "github.com/graphql-go/graphql"

// CreateSchema generates the schema necessary for application
func CreateSchema(fields interface{}) (graphql.Schema, error) {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	return schema, err
}
