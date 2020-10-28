package usermutations

import (
	u "wedeal/modules/user"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/graphql-go/graphql"
)

var addUserReturnType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AddUserReturnType",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: u.UserType,
			},
			"error": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// AddUserField returns the field for the add user mutation
var AddUser = &graphql.Field{
	Type:        addUserReturnType,
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
					"confirmPassword": &graphql.InputObjectFieldConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			})),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})

		name := input["name"].(string)
		email := input["email"].(string)
		password := input["password"].(string)
		confirmPassword := input["confirmPassword"].(string)

		result := make(map[string]interface{})

		result["user"] = nil

		// Check if any field is missing
		if name == "" || email == "" || password == "" || confirmPassword == "" {
			result["error"] = "You must provide all fields."
			return result, nil
		}

		// filter := bson.M{"email": email}
		// var userExists u.User
		// err := u.UsersCollection.FindOne(context.TODO(), filter).Decode(&userExists)
		// if err == nil {
		// 	// It means it could find a valid user
		// 	result["error"] = "This email is already registered."
		// 	return result, nil
		// }

		if password != confirmPassword {
			result["error"] = "Your password should match your confirm password."
			return result, nil
		}

		// TODO: perform password validations

		// Validate email
		notValidEmail := validation.Validate(email, is.Email)
		if notValidEmail != nil {
			result["error"] = "Email is not valid."
			return result, nil
		}

		// passwordHash, err := db.HashPassword(password)
		// if err != nil {
		// 	result["error"] = "Something went wrong. Please try again."
		// 	return result, nil
		// }

		// user := u.User{
		// 	Name:      name,
		// 	Email:     email,
		// 	Password:  passwordHash,
		// 	CreatedAt: time.Now(),
		// 	UpdatedAt: time.Now(),
		// }

		// createdUser, err := u.UsersCollection.InsertOne(context.TODO(), user)
		// if err != nil {
		// 	result["error"] = "Failed to create user. Please try again."
		// 	return result, nil
		// }

		result["user"] = map[string]interface{}{
			"id": "",
			// "id":       createdUser.InsertedID.(primitive.ObjectID).Hex(),
			"name":     name,
			"email":    email,
			"password": password,
		}

		return result, nil
	},
}
