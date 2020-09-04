package context

import (
	"context"
	"net/http"
	"wedeal/loader"

	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

// Key handles the context key values
type Key string

// HandleRequest adds the context for the graphql resolvers
func HandleRequest(w http.ResponseWriter, r *http.Request, graphqlHandler *handler.Handler, db *mongo.Database) {
	var loaders = make(map[string]*dataloader.Loader, 1)
	var client = loader.Client{
		Database: db,
	}

	listUsers := loader.NewListUsersBatchLoadFunc(db)

	loaders["GetCategory"] = dataloader.NewBatchedLoader(loader.GetCategoryBatchFn)
	loaders["ListUsers"] = dataloader.NewBatchedLoader(listUsers)
	// loaders["ListUser"] = dataloader.NewBatchedLoader(loader.GetCategoryBatchFn)

	ctx := context.WithValue(r.Context(), Key("loaders"), loaders)
	ctx = context.WithValue(ctx, Key("client"), &client)

	graphqlHandler.ContextHandler(ctx, w, r)
}
