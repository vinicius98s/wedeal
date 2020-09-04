package loader

import "go.mongodb.org/mongo-driver/mongo"

type Client struct {
	Database *mongo.Database
}

type ResolverKey struct {
	Key    string
	Client *Client
}

func (rk *ResolverKey) GetClient() *Client {
	return rk.Client
}

func NewResolverKey(key string, client *Client) *ResolverKey {
	return &ResolverKey{
		Key:    key,
		Client: client,
	}
}

func (rk *ResolverKey) String() string {
	return rk.Key
}

func (rk *ResolverKey) Raw() interface{} {
	return rk.Key
}
