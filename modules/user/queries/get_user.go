package userqueries

import (
	"wedeal/modules/user"

	"github.com/graphql-go/graphql"
)

// GetUser query
var GetUser = &graphql.Field{
	Type:        user.UserType,
	Description: "Get user by id or email",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var users []user.User
		user := user.User{ID: "1", Name: "Vinicius", Email: "vinicius@sales.com", Password: "vinicius"}
		users = append(users, user)

		email, emailOk := p.Args["email"].(string)
		id, idOk := p.Args["id"].(string)

		if idOk || emailOk {
			for _, u := range users {
				if u.ID == id || u.Email == email {
					return u, nil
				}
			}
		}

		return nil, nil
	},
}
