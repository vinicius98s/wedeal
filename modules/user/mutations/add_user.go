package usermutations

import (
	"wedeal/modules/user"

	"github.com/graphql-go/graphql"
)

type addUserInput struct {
	name     string
	email    string
	password string
}

// AddUserField returns the field for the add user mutation
func AddUserField() *graphql.Field {
	return &graphql.Field{
		Type:        user.UserType,
		Description: "Create a new user",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "AddUserInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"name": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"email": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
				})),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var name string
			var email string
			var password string

			for k, v := range p.Args["input"].(map[string]interface{}) {
				switch k {
				case "name":
					name = v.(string)
				case "email":
					email = v.(string)
				case "password":
					password = v.(string)
				}
			}

			return user.User{ID: 4, Name: name, Email: email, Password: password}, nil
		},
	}
}
