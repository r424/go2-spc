package handlers

import (
	"HW1/models"
	"encoding/json"
	"log"
	"net/http"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllItems(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get infos about all items in database")
	if len(models.DB) == 0 {
		writer.WriteHeader(403)
		msg := models.Error{Error: "No one items found in store back"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
