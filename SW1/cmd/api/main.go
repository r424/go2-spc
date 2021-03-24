package main

import (
	"SW1/internal/app/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	port string = "8080"
)


func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()
	router.HandleFunc("/grab", api.Grab).Methods("POST")
	router.HandleFunc("/solve", api.Solve).Methods("GET")

	log.Println("Router initalizing successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
