package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	// start server
	log.Fatal(http.ListenAndServe(":12345", router))
}
