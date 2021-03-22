package handlers

import (
	"HW1/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.Error{Error: "Do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	item, ok := models.FindItemById(id)

	if !ok {
		writer.WriteHeader(404)
		msg := models.Error{Error: "Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(item)
}

func CreateItem(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.Error{Error: "Do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newItem models.Item
	err = json.NewDecoder(request.Body).Decode(&newItem)

	if err != nil {
		msg := models.Message{Message: "Provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok := models.FindItemById(id)

	if ok {
		writer.WriteHeader(400)
		msg := models.Error{Error: "Item with that id already exists"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	models.SetItemById(id, newItem)
	writer.WriteHeader(201)
	msg := models.Message{Message: "Item created"}
	json.NewEncoder(writer).Encode(msg)
}

func UpdateItemById(writer http.ResponseWriter, request *http.Request) {

	initHeaders(writer)

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.Error{Error: "Do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newItem models.Item
	err = json.NewDecoder(request.Body).Decode(&newItem)

	if err != nil {
		msg := models.Message{Message: "Provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok := models.FindItemById(id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Error{Error: "Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	models.SetItemById(id, newItem)

	writer.WriteHeader(202)
	json.NewEncoder(writer).Encode(newItem)
}

func DeleteItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.Error{Error: "Do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok := models.FindItemById(id)

	if !ok {
		writer.WriteHeader(404)
		msg := models.Error{Error: "Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	models.DeleteItemById(id)
	writer.WriteHeader(202)
	msg := models.Message{Message: "Item deleted"}
	json.NewEncoder(writer).Encode(msg)
}