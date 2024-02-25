package main

import (
	"BookStore_Api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := router.Router()
	fmt.Println("Port is started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
