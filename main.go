// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Person struct {
// 	Name  string `bson:"name"`
// 	Age   int    `bson:"age"`
// 	Email string `bson:"email"`
// }

// // Define the structures for orders and order_items
// type Order struct {
// 	ID           primitive.ObjectID `bson:"_id,omitempty"`
// 	OrderNumber  string             `bson:"orderNumber"`
// 	CustomerName string             `bson:"customerName"`
// 	// Other order fields
// }

// type OrderItem struct {
// 	ID       primitive.ObjectID `bson:"_id,omitempty"`
// 	OrderID  primitive.ObjectID `bson:"order_id,omitempty"`
// 	ItemName string             `bson:"itemName"`
// 	Quantity int                `bson:"quantity"`
// 	// Other item fields
// }

// func main() {
// 	// // MongoDB connection string
// 	connectionString := "mongodb://user:P%40ssword@localhost:27017/?retryWrites=true"

// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	// Create a MongoDB client
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer func() {
// 		if err = client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	// Ping the MongoDB server
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Access the orders and order_items collections
// 	orderCollection := client.Database("your_database_name").Collection("orders")
// 	orderItemCollection := client.Database("your_database_name").Collection("order_items")

// 	// Insert sample data into the orders collection
// 	order1 := Order{
// 		ID:           primitive.NewObjectID(),
// 		OrderNumber:  "ORD001",
// 		CustomerName: "Alice",
// 	}
// 	order2 := Order{
// 		ID:           primitive.NewObjectID(),
// 		OrderNumber:  "ORD002",
// 		CustomerName: "Bob",
// 	}

// 	_, err = orderCollection.InsertMany(context.Background(), []interface{}{order1, order2})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Insert sample data into the order_items collection
// 	item1 := OrderItem{
// 		ID:       primitive.NewObjectID(),
// 		OrderID:  order1.ID,
// 		ItemName: "Item A",
// 		Quantity: 2,
// 	}
// 	item2 := OrderItem{
// 		ID:       primitive.NewObjectID(),
// 		OrderID:  order1.ID,
// 		ItemName: "Item B",
// 		Quantity: 3,
// 	}

// 	_, err = orderItemCollection.InsertMany(context.Background(), []interface{}{item1, item2})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Define the aggregation pipeline with foreignField as order_id
// 	pipeline := bson.A{
// 		bson.D{
// 			{"$lookup", bson.D{
// 				{"from", "order_items"},
// 				{"localField", "_id"},
// 				{"foreignField", "order_id"},
// 				{"as", "items"},
// 			}},
// 		},
// 	}

// 	// Perform aggregation
// 	cursor, err := orderCollection.Aggregate(context.Background(), pipeline)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cursor.Close(context.Background())

// 	// Iterate through the cursor to access the joined data
// 	var results []bson.M
// 	if err := cursor.All(context.Background(), &results); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Process and print the results
// 	for _, result := range results {
// 		fmt.Println(result)
// 	}

// 	// fmt.Println("Connected to MongoDB successfully!")

// 	// database := client.Database("test")
// 	// collection := database.Collection("people")

// 	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	// defer cancel()
// 	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

// 	// defer func() {
// 	// 	if err = client.Disconnect(ctx); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }()

// 	// // Ping the MongoDB server
// 	// err = client.Ping(ctx, nil)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// fmt.Println("Connected to MongoDB successfully!")
// }

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post represents the Post document in MongoDB
type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title   string             `bson:"title" json:"title"`
	Content string             `bson:"content" json:"content"`
	UserID  primitive.ObjectID `bson:"user_id" json:"user_id"`
}

// User represents the User document in MongoDB
type User struct {
	ID    primitive.ObjectID   `bson:"_id,omitempty"`
	Name  string               `bson:"name"`
	Email string               `bson:"email"`
	Posts []primitive.ObjectID `bson:"posts"`
}

type UserWithPosts struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
	Posts []Post             `json:"posts"`
}

type UserPostCount struct {
	ID    string `bson:"_id" json:"id"`
	Count int    `bson:"count" json:"count"`
}

func main() {
	runJoinTwoCollection()
	// // MongoDB connection settings
	// connectionString := "mongodb://user:P%40ssword@localhost:27017/?retryWrites=true"
	// clientOptions := options.Client().ApplyURI(connectionString)
	// client, err := mongo.Connect(context.Background(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err := client.Disconnect(context.Background()); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// // Access the database and collection
	// database := client.Database("your_database_name")
	// usersCollection := database.Collection("users")
	// postsCollection := database.Collection("posts")

	// // MongoDB aggregation pipeline to fetch user and their posts
	// pipeline := bson.D{
	// 	{"$match", bson.D{{"name", "John Doe"}}}, // Match the user by name
	// 	{"$lookup", bson.D{
	// 		{"from", "posts"},       // Lookup in the 'posts' collection
	// 		{"localField", "posts"}, // Local field in 'users' collection
	// 		{"foreignField", "_id"}, // Foreign field in 'posts' collection
	// 		{"as", "posts"},         // Name for the merged array field
	// 	}},
	// }
	// // Perform the aggregation
	// // Perform the aggregation
	// cursor, err := usersCollection.Aggregate(context.Background(), mongo.Pipeline{pipeline})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(context.Background())

	// // Decode the result
	// var user User
	// for cursor.Next(context.Background()) {
	// 	if err := cursor.Decode(&user); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// // Print the user and their posts
	// fmt.Println("User:", user)

	// // Insert a new post for the user
	// newPost := Post{
	// 	Title:   "New Post Title",
	// 	Content: "New Post Content",
	// }
	// insertResult, err := postsCollection.InsertOne(context.Background(), newPost)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted new post ID:", insertResult.InsertedID)

	// // Update the user's posts array with the new post ID
	// user.Posts = append(user.Posts, insertResult.InsertedID.(primitive.ObjectID))
	// updateResult, err := usersCollection.UpdateOne(
	// 	context.Background(),
	// 	bson.M{"_id": user.ID},
	// 	bson.D{
	// 		{"$set", bson.D{{"posts", user.Posts}}},
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Updated user's posts:", updateResult.ModifiedCount)
}

func runJoinTwoCollection() {
	connectionString := "mongodb://user:P%40ssword@localhost:27017/?retryWrites=true"
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	database := client.Database("your_database_name")
	usersCollection := database.Collection("users")
	postsCollection := database.Collection("posts")

	var userId primitive.ObjectID
	{
		fmt.Println("--- Insert User ---")
		newUser := User{
			Name:  "Bob",
			Email: "bob@email.com",
		}

		insertResult, err := usersCollection.InsertOne(context.Background(), newUser)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted new post ID:", insertResult.InsertedID)

		userId = insertResult.InsertedID.(primitive.ObjectID)

	}

	var postIDs []primitive.ObjectID
	{
		fmt.Println("--- Insert Posts ---")
		var interfacePosts []interface{}

		newPost := Post{
			Title:   "New Post Title",
			Content: "New Post Content",
			UserID:  primitive.ObjectID(userId),
		}

		newPost2 := Post{
			Title:   "New Post Title2",
			Content: "New Post Content2",
			UserID:  primitive.ObjectID(userId),
		}

		interfacePosts = append(interfacePosts, newPost, newPost2)
		ctx := context.Background()
		insertResult, err := postsCollection.InsertMany(ctx, interfacePosts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted new post ID:", insertResult.InsertedIDs)

		for _, id := range insertResult.InsertedIDs {
			postIDs = append(postIDs, id.(primitive.ObjectID))
		}

	}

	{
		fmt.Println("--- Update User with Posts ---")
		updateResult, err := usersCollection.UpdateOne(
			context.Background(),
			bson.M{"_id": userId},
			bson.D{
				{Key: "$set", Value: bson.D{{"posts", postIDs}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		// Print the update result
		fmt.Printf("Updated %v documents\n", updateResult.ModifiedCount)

	}

	{
		fmt.Println("--- Find Users---")
		ctx := context.Background()
		var users []User
		cursor, err := usersCollection.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}

		defer cursor.Close(ctx)

		// Decode the results into the users slice
		err = cursor.All(ctx, &users)
		if err != nil {
			log.Fatal(err)
		}

		PrintPretiseJson(users)

	}

	{
		fmt.Println("--- Find User by Id ---")
		ctx := context.Background()
		var user User
		filter := bson.D{{"_id", userId}}
		err := usersCollection.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		PrintPretiseJson(user)

	}

	{
		fmt.Println("--- Find User by Id with Post---")
		ctx := context.Background()
		var user User
		filter := bson.D{{"_id", userId}}
		err := usersCollection.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		var userPosts []Post
		postsFilter := bson.M{"_id": bson.M{"$in": user.Posts}}
		cursor, err := postsCollection.Find(context.Background(), postsFilter)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var post Post
			if err := cursor.Decode(&post); err != nil {
				log.Fatal(err)
			}
			userPosts = append(userPosts, post)
		}

		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}

		// Combine user and posts data into one response
		userWithPosts := UserWithPosts{
			ID:    user.ID,
			Name:  user.Name,
			Posts: userPosts,
		}
		PrintPretiseJson(userWithPosts)

	}

	{
		fmt.Println("--- Aggregation ---")

		pipeline := mongo.Pipeline{
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "posts"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "user_id"},
				{Key: "as", Value: "posts"},
			}}},
			{{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "name", Value: 1},
				{Key: "email", Value: 1},
				{Key: "posts._id", Value: 1},
				{Key: "posts.title", Value: 1},
				{Key: "posts.content", Value: 1},
				{Key: "posts.user_id", Value: 1},
			}}},
		}

		// Perform the aggregation
		cursor, err := usersCollection.Aggregate(context.Background(), pipeline)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())

		// Iterate over the results
		var results []UserWithPosts
		if err := cursor.All(context.Background(), &results); err != nil {
			log.Fatal(err)
		}

		PrintPretiseJson(results)

	}

	{
		fmt.Println("--- Aggregation2 ---")
		// Define the aggregation pipeline
		pipeline := mongo.Pipeline{
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "posts"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "user_id"},
				{Key: "as", Value: "posts"},
			}}},
			{{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "count", Value: bson.D{{Key: "$size", Value: "$posts"}}},
			}}},
		}

		// Perform the aggregation
		cursor, err := usersCollection.Aggregate(context.Background(), pipeline)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())

		// Iterate over the results
		var results []UserPostCount
		if err := cursor.All(context.Background(), &results); err != nil {
			log.Fatal(err)
		}

		PrintPretiseJson(results)
	}

	{
		fmt.Println("--- Purge user and posts by id ---")

		// Create a filter to find the user by ID
		filter := bson.M{"_id": userId}

		// Delete the user
		deleteResult, err := usersCollection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}

		// Check the number of deleted documents
		if deleteResult.DeletedCount == 1 {
			fmt.Println("User deleted successfully")
		} else {
			fmt.Println("User not found or not deleted")
		}

		filter = bson.M{"user_id": userId}

		// Delete posts associated with the user
		deleteResult, err = postsCollection.DeleteMany(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}

		// Check the number of deleted documents
		fmt.Printf("%d posts deleted\n", deleteResult.DeletedCount)
	}

	{

		deleteResult, err := usersCollection.DeleteMany(context.Background(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Purged %d documents from collection %s\n", deleteResult.DeletedCount, "User")

		deleteResult, err = postsCollection.DeleteMany(context.Background(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Purged %d documents from collection %s\n", deleteResult.DeletedCount, "Post")
	}

}

func PrintPretiseJson(data interface{}) {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
