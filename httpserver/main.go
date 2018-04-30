package main

import (
	"log"
	"net/http"

	"learn-go-with-tests/httpserver/handlers"
)

// InMemoryPlayerStore serves as the storage of players and their scores.
//
// Replace me with a real-world storage like Postgres, SQL, etc.
type InMemoryPlayerStore struct {
	Store map[string]int
}

// NewInMemoryPlayerStore is a convenient wrapper to initialize a store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// RecordWin persists a score for the player.
func (i *InMemoryPlayerStore) RecordWin(player string) {
	i.Store[player]++
}

// GetPlayerScore retrieve the player's score from the store.
func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return i.Store[player]
}

func main() {
	server := &handlers.PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen to port 5000 %v", err)
	}
}
