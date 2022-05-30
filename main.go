package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)
	/**
	 * * Map all routes
	 */
	r.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	handleRequest()
}