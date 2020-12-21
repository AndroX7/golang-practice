package main

import (
	"fmt"
	"net/http"
	"rest-mongo-and-go/config"
	"rest-mongo-and-go/database"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	db := database.ConnectDB(conf.Mongo)
	fmt.Println(db)
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)
}
