package userqueries

import (
	"fmt"
	"strings"
	"wedeal/context"
	"wedeal/modules/user"

	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
)

// ListUsers query
var ListUsers = &graphql.Field{
	Type:        graphql.NewList(user.UserType),
	Description: "List all users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var (
			keys    dataloader.Keys
			v       = p.Context.Value
			loaders = v(context.Key("loaders")).(map[string]*dataloader.Loader)

			handleErrors = func(errors []error) error {
				var errs []string
				for _, e := range errors {
					errs = append(errs, e.Error())
				}
				return fmt.Errorf(strings.Join(errs, "\n"))
			}
		)

		thunk := loaders["ListUsers"].LoadMany(p.Context, keys)
		return func() (interface{}, error) {
			users, errs := thunk()
			if len(errs) > 0 {
				return nil, handleErrors(errs)
			}
			return users, nil
		}, nil
	},
}
