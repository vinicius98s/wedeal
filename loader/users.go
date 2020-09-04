package loader

import (
	"context"

	u "wedeal/modules/user"

	"github.com/graph-gophers/dataloader"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userLoader struct {
	db *mongo.Database
}

// func (client *Client) ListUsers(filter primitive.M) (u.User, error) {
// 	var user u.User
// 	err := client.Database.Collection("users").Find(context.TODO(), filter).Decode(&user)
// 	return user, err
// }

// NewListUsersBatchLoadFunc mano nao sei
func NewListUsersBatchLoadFunc(db *mongo.Database) dataloader.BatchFunc {
	return (&userLoader{db}).loadBatch
}

func (loader *userLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)
	_, err := loadKeys(keys)
	if err != nil {
		return loadBatchError(err, n)
	}

	var users []*u.User
	cursor, err := loader.db.Collection("users").Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return loadBatchError(err, n)
	}

	for cursor.Next(context.TODO()) {
		var user u.User
		err := cursor.Decode(&user)
		if err != nil {
			return loadBatchError(err, n)
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return loadBatchError(err, n)
	}

	cursor.Close(context.TODO())

	res := make([]*dataloader.Result, n)
	for i, c := range users {
		res[i] = &dataloader.Result{Data: c}
	}

	return res
}
