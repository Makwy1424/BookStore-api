package db

import (
	"BookStore_Api/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func InsertOneBook(book model.Book) {
	inserted, err := Collection.InsertOne(context.TODO(), book)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Id of insert movie: ", inserted.InsertedID)
}

func UpdateOneBook(bookId string, book model.Book) {
	id, err := primitive.ObjectIDFromHex(bookId)

	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.M{"_id": id}
	updateBook := bson.M{}

	if book.Title != "" {
		updateBook["title"] = book.Title
	}
	if book.Author != "" {
		updateBook["author"] = book.Author
	}
	if book.Price != 0 {
		updateBook["price"] = book.Price
	}
	if book.Quantity != 0 {
		updateBook["quantity"] = book.Quantity
	}
	if book.ISBN != "" {
		updateBook["ISBN"] = book.ISBN
	}
	if book.Genre != "" {
		updateBook["Genre"] = book.Genre
	}
	if book.Desc != "" {
		updateBook["Desc"] = book.Desc
	}
	update := bson.M{"$set": updateBook}
	result, err := Collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func GetBooks() []primitive.M {
	cursor, err := Collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	var books []primitive.M
	for cursor.Next(context.TODO()) {
		var book bson.M
		err := cursor.Decode(&book)

		if err != nil {
			log.Fatal(err)
			return nil
		}

		books = append(books, book)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, context.Background())

	return books
}

func GetOneBook(bookId string) primitive.M {
	id, err := primitive.ObjectIDFromHex(bookId)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	filter := bson.M{"_id": id}
	cursor, err := Collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	cursor.Next(context.TODO())
	var book bson.M
	err = cursor.Decode(&book)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, context.TODO())

	return book
}

func DeleteOneBook(bookId string) {
	id, err := primitive.ObjectIDFromHex(bookId)

	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.M{"_id": id}
	result, err := Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Deleted %v documents\n", result.DeletedCount)

}

func DeleteAllBook() {
	filter := bson.M{}
	result, err := Collection.DeleteMany(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("deleted %v documents \n", result.DeletedCount)
}
