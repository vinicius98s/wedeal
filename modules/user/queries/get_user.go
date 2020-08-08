package userqueries

import (
	"wedeal/modules/user"

	"github.com/graphql-go/graphql"
)

// GetUserQuery generates the fields for the user module
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type:        user.UserType,
		Description: "Get user by id or email",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var users []user.User
			user := user.User{ID: 1, Name: "Vinicius", Email: "vinicius@email.com", Password: "vinicius"}
			users = append(users, user)

			email, emailOk := p.Args["email"].(string)
			id, idOk := p.Args["id"].(int)

			if idOk || emailOk {
				for _, u := range users {
					if int(u.ID) == id || u.Email == email {
						return u, nil
					}
				}
			}

			return nil, nil
		},
	}
}
