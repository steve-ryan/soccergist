package main

import (
	"fmt"
	"log"
	"net/http"

	"soccergist/implementations/go/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/webhook", handlers.WebHookHandler).Methods("GET")
	router.HandleFunc("/webhook", handlers.WebHookPostHandler).Methods("POST")

	fmt.Println("Server Starting up")
	log.Fatal(http.ListenAndServe(":3000", router))
}
