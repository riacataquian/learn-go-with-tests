package main

import (
	"log"
	"net/http"

	"learn-go-with-tests/httpserver/handlers"
)

func main() {
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
	// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
	handler := http.HandlerFunc(handlers.PlayerHandler)

	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen to port 5000 %v", err)
	}
}
