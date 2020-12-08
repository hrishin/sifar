package hello

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MountRoutes(router *mux.Router) {
	router.HandleFunc("/hello", HelloWorldHandler)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")
}
