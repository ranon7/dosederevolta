package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	api := API{}

	port := 8081

	s := &http.Server{
		Handler: api.Routes(ctx),
		Addr:    fmt.Sprintf(":%d", port),
	}

	log.Printf("Listening to port %d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(fmt.Errorf("unable to start server: %w", err))
	}
}
