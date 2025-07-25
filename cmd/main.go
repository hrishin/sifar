package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hrishin/sifar/pkg/hello"
)

func main() {
	// change the port in Dockerfile as well
	port := 8000

	r := mux.NewRouter()
	hello.MountRoutes(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: r,
	}

	log.Printf("Starting hello server on port %d \n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Failed to start the sever: error: %v", err)
	}
}
