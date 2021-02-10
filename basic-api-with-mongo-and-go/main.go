package main

import (
	"basic-api-with-mongo-and-go/api/books"
	"basic-api-with-mongo-and-go/db"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	db.Connect()
}

func main() {

	r := mux.NewRouter()
	bc := new(books.Controller)
	r.HandleFunc("/books", bc.FindAll).Methods("GET")
	r.HandleFunc("/books/{id}", bc.FindByID).Methods("GET")
	r.HandleFunc("/books", bc.Create).Methods("POST")
	r.HandleFunc("/books/{id}", bc.FindAll).Methods("PUT")
	r.HandleFunc("/books/{id}", bc.FindAll).Methods("DELETE")
	// r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "Hello World")
	// }).Methods("GET")
	http.ListenAndServe("localhost:3007", r)
}
