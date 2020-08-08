package usermutations

import (
	"github.com/graphql-go/graphql"
)

// GetMutationFields return all fields for users mutations
func GetMutationFields() graphql.Fields {
	addUserMutation := AddUserField()
	mutations := make(map[string]*graphql.Field)

	mutations["addUser"] = addUserMutation

	return mutations
}
