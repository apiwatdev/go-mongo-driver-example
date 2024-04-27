# Go MongoDB CRUD Operations Example

This Go application demonstrates CRUD (Create, Read, Update, Delete) operations using the MongoDB Go driver. It includes examples of inserting users and posts, updating user data with associated posts, querying data, performing aggregations, and deleting records.


## Prerequisites
- Go installed on your system
- MongoDB server running locally or accessible via connection string

## Setup
1. Clone the repository or download the code files.
2. Install the required Go packages:
```
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/mongo/options
```
3. Update the MongoDB connection string in the main function's connectionString variable. 

## Usage
1. Run the main program:
```
go run main.go
```
2. Follow the instructions in the console output to see the CRUD operations in action.
## Code Structure
main.go: Contains the main program with CRUD examples.

## Functionality
- Inserting a user and associated posts.
- Updating a user with new posts.
- Finding users by ID and retrieving associated posts.
- Performing aggregation operations to count posts per user.
- Purging users and associated posts by ID or purging all data in collections.

## Notes
- Ensure your MongoDB server is running and accessible with the provided connection string.
- This example covers basic CRUD operations; customize and expand the code as needed for your application.

## Example Result

```
--- Insert User ---
Inserted new post ID: ObjectID("662d0caded3d1942bf8b90e7")
--- Insert Posts ---
Inserted new post ID: [ObjectID("662d0caded3d1942bf8b90e8") ObjectID("662d0caded3d1942bf8b90e9")]
--- Update User with Posts ---
Updated 1 documents
--- Find Users---
[
    {
        "ID": "662d0caded3d1942bf8b90e7",
        "Name": "Bob",
        "Email": "bob@email.com",
        "Posts": [
            "662d0caded3d1942bf8b90e8",
            "662d0caded3d1942bf8b90e9"
        ]
    }
]
--- Find User by Id ---
{
    "ID": "662d0caded3d1942bf8b90e7",
    "Name": "Bob",
    "Email": "bob@email.com",
    "Posts": [
        "662d0caded3d1942bf8b90e8",
        "662d0caded3d1942bf8b90e9"
    ]
}
--- Find User by Id with Post---
{
    "_id": "662d0caded3d1942bf8b90e7",
    "name": "Bob",
    "email": "",
    "posts": [
        {
            "_id": "662d0caded3d1942bf8b90e8",
            "title": "New Post Title",
            "content": "New Post Content",
            "user_id": "662d0caded3d1942bf8b90e7"
        },
        {
            "_id": "662d0caded3d1942bf8b90e9",
            "title": "New Post Title2",
            "content": "New Post Content2",
            "user_id": "662d0caded3d1942bf8b90e7"
        }
    ]
}
--- Aggregation ---
[
    {
        "_id": "662d0caded3d1942bf8b90e7",
        "name": "Bob",
        "email": "bob@email.com",
        "posts": [
            {
                "_id": "662d0caded3d1942bf8b90e8",
                "title": "New Post Title",
                "content": "New Post Content",
                "user_id": "662d0caded3d1942bf8b90e7"
            },
            {
                "_id": "662d0caded3d1942bf8b90e9",
                "title": "New Post Title2",
                "content": "New Post Content2",
                "user_id": "662d0caded3d1942bf8b90e7"
            }
        ]
    }
]
--- Aggregation2 ---
[
    {
        "id": "662d0caded3d1942bf8b90e7",
        "count": 2
    }
]
--- Purge user and posts by id ---
User deleted successfully
2 posts deleted

```