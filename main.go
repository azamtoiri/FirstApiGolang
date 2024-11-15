package main

import (
	"FirstApi/utils"
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
	port                   string
	bookResourcePrefix     string = apiPrefix + "/book"
	manyBookResourcePrefix string = apiPrefix + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env file", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResourcePrefix(router, manyBookResourcePrefix)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
