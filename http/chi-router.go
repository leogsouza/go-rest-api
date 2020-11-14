package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

// NewChiRouter creates a new Router instance based on chi library
func NewChiRouter() Router {
	return &chiRouter{}
}

func (cr *chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (cr *chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (cr *chiRouter) SERVE(port string) {
	log.Printf("Chi HTTP runing on %s", port)
	http.ListenAndServe(port, chiDispatcher)
}
