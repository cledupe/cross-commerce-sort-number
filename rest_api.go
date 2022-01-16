package main

import (
	"log"
	"net/http"

	"github.com/cledupe/cross-commerce-sort-number/framework/load/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/numbers", handler.GetNumbers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
