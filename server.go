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

	log.Println("Server listening on port", port)

	log.Fatalln(http.ListenAndServe(port, router))
}
