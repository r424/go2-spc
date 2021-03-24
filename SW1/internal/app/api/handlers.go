package api

import (
	"SW1/internal/app/models"
	"encoding/json"
	"net/http"
)

var c models.Coefficients

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func Grab(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	err := json.NewDecoder(request.Body).Decode(&c)

	if err != nil {
		msg := models.Message{Message: "Provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(201)
}

func Solve(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	r := models.CoefficientsAndSolution{
		Coefficients: c,
		Nroots:       findRoots(c),
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(r)
}
