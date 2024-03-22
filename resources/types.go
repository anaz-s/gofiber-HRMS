package resources

import "go.mongodb.org/mongo-driver/mongo"

// User struct represents the schema of a user document in MongoDB
type User struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"` // ID field of the user document
	FirstName string `bson:"firstName" json:"firstName"`         // First name of the user
	LastName  string `bson:"lastName" json:"lastName"`           // Last name of the user
}

// MongoInstance struct holds the MongoDB client instance
type MongoInstance struct {
	Client *mongo.Client // MongoDB client instance
}
