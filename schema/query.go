package schema

import (
	c "wedeal/modules/category/queries"
	u "wedeal/modules/user/queries"

	"github.com/graphql-go/graphql"
)

// MergeFields gets a variadic parameter of graphql fields and get them together
func MergeFields(fields ...graphql.Fields) graphql.Fields {
	result := make(map[string]*graphql.Field)
	for _, field := range fields {
		for key, value := range field {
			result[key] = value
		}
	}
	return result
}

func getRootQueryFields() graphql.Fields {
	return MergeFields(u.Fields, c.Fields)
}

// RootQuery exports the RootQuery graphql object
var RootQuery = getRootQueryFields()
