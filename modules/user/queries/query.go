package userqueries

import (
	"github.com/graphql-go/graphql"
)

// Fields gets all query fields for the user module
var Fields = graphql.Fields{
	"user":  GetUser,
	"users": ListUsers,
}
