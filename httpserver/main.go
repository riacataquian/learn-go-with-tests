package main

import (
	"log"
	"net/http"

	"learn-go-with-tests/httpserver/handlers"
)

// InMemoryPlayerStore ...
type InMemoryPlayerStore struct {
	Store map[string]int
}

// NewInMemoryPlayerStore ...
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// RecordWin ...
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.Store[name]++
}

// GetPlayerScore ...
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.Store[name]
}

func main() {
	server := &handlers.PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen to port 5000 %v", err)
	}
}
