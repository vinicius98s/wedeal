package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// UsersCollection in a database context
var UsersCollection *mongo.Collection

// SetupCollection handles the user collection within database
func SetupCollection(c *mongo.Database) {
	UsersCollection = c.Collection("users")
}
