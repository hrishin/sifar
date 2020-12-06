package hello

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MountRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloWorldHandler)
	return r
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")
}
