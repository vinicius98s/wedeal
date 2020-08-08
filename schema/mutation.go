package schema

import (
	u "wedeal/modules/user/mutations"

	"github.com/graphql-go/graphql"
)

// GetMutationFields gathers all packages mutations fields
func GetMutationFields() graphql.Fields {
	mutationFields := make(map[string]*graphql.Field)
	userMutationFields := u.GetMutationFields()

	for k, v := range userMutationFields {
		mutationFields[k] = v
	}

	return mutationFields
}

// Mutation exports the Mutation graphql object
var Mutation = GetMutationFields()
