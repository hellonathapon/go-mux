package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Owner string `json:"owner"`
}

type Posts []Post

/**
 * * Method in Go is essentially Class like method
 */
func (p *Posts) pushPost(a Post) Posts {
	post = append(post, a)
	return post
}

// Mockup Data
// global variable   
var post Posts = Posts{ Post{"Good morning!", "Go is awesome language", "Nathapon"} }

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(post) // <-- w is http responseWriter
}

func createPost(w http.ResponseWriter, r *http.Request) {
	
	decoder := json.NewDecoder(r.Body)
	var t Post
	err := decoder.Decode(&t) // <-- pass address of t
	if err != nil {
		panic(err)
	}

	post.pushPost(t)
	fmt.Fprintf(w, "Successfully created")
}