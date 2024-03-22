package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"myproject/resources"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define MongoDB connection parameters
const (
	uri        = "mongodb://localhost:27017" // MongoDB URI
	dbName     = "sample"                    // Name of the database
	collection = "users"                     // Name of the collection
)

var mongoInit *resources.MongoInstance // Global variable to hold the MongoDB client instance

func main() {
	// Parse command line flags to specify the listening port
	port := flag.String("port", "5000", "specify a listening port")
	flag.Parse()

	// Connect to MongoDB
	if err := connect(); err != nil {
		log.Fatal(err)
	}

	// Create a new Fiber app instance
	app := fiber.New()

	// Define API routes
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/users", handleGetAllUser)     // Get all users route
	apiv1.Post("/users", handleAddUsers)      // Add user route
	apiv1.Put("/users/:id", handlePutReq)     // Update user route
	apiv1.Delete("/users/:id", handleDeleteUser) // Delete user route

	// Start the server and listen on the specified port
	app.Listen(":" + *port)
}

// Function to connect to MongoDB
func connect() error {
	// Create a background context
	ctx := context.Background()

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	// Handle connection errors
	if err != nil {
		log.Fatal(err)
	}

	// Assign the MongoDB client instance to the global variable
	mongoInit = &resources.MongoInstance{
		Client: client,
	}

	return nil
}

// Function to handle GET request for retrieving all users
func handleGetAllUser(c fiber.Ctx) error {
	// Define an empty BSON query
	query := bson.D{}

	// Execute the query to find all users
	cursor, err := mongoInit.Client.Database(dbName).Collection(collection).Find(context.Background(), query)

	// Handle query execution errors
	if err != nil {
		log.Fatal(err)
	}

	// Declare a slice to store users
	var users []resources.User

	// Decode the results into the users slice
	if err := cursor.All(context.Background(), &users); err != nil {
		return c.Status(403).SendString(err.Error())
	}

	// Return the users as JSON response
	return c.Status(200).JSON(users)
}

// Function to handle POST request for adding a new user
func handleAddUsers(c fiber.Ctx) error {
	// Get the MongoDB collection
	collection := mongoInit.Client.Database(dbName).Collection(collection)

	// Retrieve the request body
	body := c.Body()

	// Declare a variable to store the decoded user object
	var user resources.User

	// Unmarshal the request body into the user object
	if err := json.Unmarshal(body, &user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Insert the user into the collection
	res, err := collection.InsertOne(context.Background(), user)

	// Handle insertion errors
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Return the insertion result as JSON response
	return c.JSON(res)
}

// Function to handle PUT request for updating a user
func handlePutReq(c fiber.Ctx) error {
	// Extract the user ID from the request parameters
	id := c.Params("id")

	// Parse the user ID into an ObjectID
	userID, err := primitive.ObjectIDFromHex(id)

	// Handle ObjectID parsing errors
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Declare a variable to store the decoded user object
	var user resources.User

	// Define a BSON query to find the user by ID
	query := bson.D{{Key: "_id", Value: userID}}

	// Retrieve the request body
	body := c.Body()

	// Unmarshal the request body into the user object
	if err := json.Unmarshal(body, &user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Define an update document to set the first name and last name fields
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "firstName", Value: user.FirstName},
				{Key: "lastName", Value: user.LastName},
			},
		},
	}

	// Set the user ID in the user object
	user.ID = id

	// Update the user document in the collection
	err = mongoInit.Client.Database(dbName).Collection(collection).FindOneAndUpdate(context.Background(), query, update).Err()

	// Handle update errors
	if err == mongo.ErrNoDocuments {
		return c.Status(500).SendString(err.Error())
	}

	return nil
}

// Function to handle DELETE request for deleting a user
func handleDeleteUser(c fiber.Ctx) error {
	// Extract the user ID from the request parameters
	id := c.Params("id")

	// Parse the user ID into an ObjectID
	userID, err := primitive.ObjectIDFromHex(id)

	// Handle ObjectID parsing errors
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Define a BSON query to find the user by ID
	query := bson.D{{Key: "_id", Value: userID}}

	// Delete the user document from the collection
	err = mongoInit.Client.Database(dbName).Collection(collection).FindOneAndDelete(context.Background(), query).Err()

	// Handle deletion errors
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Return success message as JSON response
	return c.Status(200).JSON("record deleted")
}
