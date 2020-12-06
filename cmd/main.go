package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hrishin/sifar/pkg/hello"
)

func main() {
	// change the port in Dockerfile as well
	port := 5000

	r := hello.MountRoutes()
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: r,
	}

	log.Printf("Starting hello server on port %d \n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Failed to start the sever: error: %v", err)
	}
}
