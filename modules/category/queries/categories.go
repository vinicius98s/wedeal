package categoryqueries

import (
	"fmt"
	"strings"
	"wedeal/context"
	"wedeal/loader"
	"wedeal/modules/category"

	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
)

var Fields = graphql.Fields{
	"categories": &graphql.Field{
		Type: graphql.NewList(category.CategoryType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var (
				// Same categories from `Client.ListCategories`
				// with ID (1, 2, 3).
				// `GetCategoryBatchFn` batch them.
				categoryIDs []uint = []uint{1, 2, 3}
				keys        dataloader.Keys
				v           = p.Context.Value
				loaders     = v(context.Key("loaders")).(map[string]*dataloader.Loader)
				c           = v(context.Key("client")).(*loader.Client)

				handleErrors = func(errors []error) error {
					var errs []string
					for _, e := range errors {
						errs = append(errs, e.Error())
					}
					return fmt.Errorf(strings.Join(errs, "\n"))
				}
			)

			for _, categoryID := range categoryIDs {
				key := loader.NewResolverKey(fmt.Sprintf("%d", categoryID), c)
				keys = append(keys, key)
			}

			thunk := loaders["GetCategory"].LoadMany(p.Context, keys)
			return func() (interface{}, error) {
				categories, errs := thunk()
				if len(errs) > 0 {
					return nil, handleErrors(errs)
				}
				return categories, nil
			}, nil
		},
	},
}
