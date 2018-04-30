package main

import (
	"log"
	"net/http"

	"learn-go-with-tests/httpserver/handlers"
)

// InMemoryPlayerStore ...
type InMemoryPlayerStore struct{}

// GetPlayerScore ...
func (i *InMemoryPlayerStore) GetPlayerScore(name string) string {
	return "123"
}

func main() {
	server := &handlers.PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen to port 5000 %v", err)
	}
}
