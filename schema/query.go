package schema

import (
	u "wedeal/modules/user/queries"

	"github.com/graphql-go/graphql"
)

// GetRootQueryFields gathers all packages query fields
func GetRootQueryFields() graphql.Fields {
	rootQueryFields := make(map[string]*graphql.Field)
	userFields := u.GetQueryFields()

	for k, v := range userFields {
		rootQueryFields[k] = v
	}

	return rootQueryFields
}

// RootQuery exports the RootQuery graphql object
var RootQuery = GetRootQueryFields()
