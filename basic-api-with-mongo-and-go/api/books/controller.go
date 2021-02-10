package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct{}

//find all
func (Controller) FindAll(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello From Find All")
}

//find by id...
func (Controller) FindByID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	fmt.Fprint(res, "Hello From Find By Id ", params["id"])

}

// create...
func (Controller) Create(res http.ResponseWriter, req *http.Request) {
	var book Book
	jsonBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	err = json.Unmarshal(jsonBytes, &book)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// log.Printf("Books: %v <<<<<<<<<<<<<<<<<", book)
	err = book.Save()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%v", book)
}

// delete...
func (Controller) Delete(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello From Find By Id")
}

// update...
func (Controller) Update(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello From Find By Id")
}
