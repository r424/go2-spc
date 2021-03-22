package main

import (
	"HW1/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port           string
	itemPrefix     string = apiPrefix + "/item"
	manyItemPrefix string = apiPrefix + "/items"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildItemResource(router, itemPrefix)
	utils.BuildManyItemsResourcePrefix(router, manyItemPrefix)

	log.Println("Router initalizing successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
