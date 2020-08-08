package userqueries

import (
	"github.com/graphql-go/graphql"
)

// GetQueryFields generates the fields for the user module
func GetQueryFields() graphql.Fields {
	getUserQuery := GetUserQuery()

	return graphql.Fields{
		"user": getUserQuery,
	}
}
