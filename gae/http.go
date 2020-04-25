package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func httpServer(ctx context.Context) error {
	{
		http.HandleFunc("/_ah/warmup", handlerWarmup)
		http.HandleFunc("/", handlerIndex)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		return err
	}
	return nil
}
