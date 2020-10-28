package schema

import (
	u "wedeal/modules/user/mutations"

	"github.com/graphql-go/graphql"
)

func getMutationFields() graphql.Fields {
	return MergeFields(u.Mutations)
}

// Mutation exports the Mutation graphql object
var Mutation = getMutationFields()
