package userqueries

import (
	"github.com/graphql-go/graphql"
)

// GetQueryFields generates the fields for the user module
func GetQueryFields() map[string]*graphql.Field {
	getUserQuery := GetUserQuery()
	usersMap := make(map[string]*graphql.Field)

	usersMap["user"] = getUserQuery

	return usersMap
}
