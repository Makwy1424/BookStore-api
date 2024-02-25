package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const ConnectionString string = "mongodb+srv://srr:123@cluster0.zocfczn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const DatabaseName string = "book_store"
const CollectionName string = "book"

var Collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(ConnectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Db connection Successful")

	Collection = client.Database(DatabaseName).Collection(CollectionName)

	fmt.Println("Book collection is ready.")
}
