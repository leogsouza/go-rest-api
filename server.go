package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8085"
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running....")
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("Server listening on port", port)

	log.Fatalln(http.ListenAndServe(port, router))
}
