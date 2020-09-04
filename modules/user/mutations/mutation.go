package usermutations

import (
	"github.com/graphql-go/graphql"
)

// Mutations is all fields for users mutations
var Mutations = graphql.Fields{
	"addUser": AddUser,
}
