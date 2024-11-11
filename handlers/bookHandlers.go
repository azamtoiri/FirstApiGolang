package handlers

import (
	"FirstApi/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend", err)
		msg := models.Message{Message: "do not user parameter ID as uncased to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	book, ok := models.FindBookById(id)

	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with ID does not exists int database"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}

}

func CreateBook(writer http.ResponseWriter, request *http.Request) {

}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {

}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {

}
