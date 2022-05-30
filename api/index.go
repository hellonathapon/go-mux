package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter().StrictSlash(true)
	/**
	 * * Map all routes
	 */
	r.HandleFunc("/", home)
	r.HandleFunc("/posts", getAllPosts)
	r.HandleFunc("/post", createPost).Methods("POST")


	/**
	 * * log.Fatal is essentially logging function
	 * * works when fails to starting the server.
	 */
	log.Fatal(http.ListenAndServe(":8081", r))
}