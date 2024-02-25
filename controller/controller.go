package controller

import (
	"BookStore_Api/db"
	"BookStore_Api/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
		return
	}

	db.InsertOneBook(book)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(params["id"], " ", book.Author, " ", book.Title, " ", book.ID, " ", book.Genre)
	db.UpdateOneBook(params["id"], book)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	book := db.GetOneBook(params["id"])

	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allBooks := db.GetBooks()

	err := json.NewEncoder(w).Encode(allBooks)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db.DeleteOneBook(params["id"])

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(params["id"])
	if err != nil {
		log.Fatal(err)
		return
	}
}

func DeleteAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db.DeleteAllBook()
	w.WriteHeader(http.StatusOK)
}
