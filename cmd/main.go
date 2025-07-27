package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/hrishin/sifar/pkg/hello"
)

func main() {
	// change the port in Dockerfile as well
	wait := time.Second * 15
	port := 8000

	r := mux.NewRouter()
	hello.MountRoutes(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: r,
	}

	errChan := make(chan error, 1)
	go func() {
		log.Printf("Starting hello server on port %d \n", port)
		if err := srv.ListenAndServe(); err != nil {
			errChan <- fmt.Errorf("failed to start the server: %v", err)
			return
		}
		log.Printf("Started server on port 0.0.0.0:%d \n", port)
	}()

	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case osSig := <-osChan:
		log.Printf("Received signal: %s. Initiating shutdown...", osSig)
		ctx, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Server shutdown error: %v", err)
			os.Exit(1)
		}
		log.Print("Server shutdown complete")

	case err := <-errChan:
		if err != nil {
			log.Printf("Server error: %v", err)
			os.Exit(1)
		}
		log.Printf("Server exited normally")
	}
}
