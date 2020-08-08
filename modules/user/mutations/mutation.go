package usermutations

import (
	"github.com/graphql-go/graphql"
)

// GetMutationFields return all fields for users mutations
func GetMutationFields() graphql.Fields {
	addUserMutation := AddUserField()

	return graphql.Fields{
		"addUser": addUserMutation,
	}
}
