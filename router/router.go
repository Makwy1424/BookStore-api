package router

import (
	"BookStore_Api/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/addBook", controller.CreateBook).Methods("POST")
	router.HandleFunc("/api/updateBook/{id}", controller.UpdateBook).Methods("POST")
	router.HandleFunc("/api/getBook/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/api/getAllBook", controller.GetAllBook).Methods("GET")
	router.HandleFunc("/api/deleteBook/{id}", controller.DeleteBook).Methods("DELETE")
	router.HandleFunc("/api/deleteAllBook", controller.DeleteAllBook).Methods("DELETE")

	return router
}
