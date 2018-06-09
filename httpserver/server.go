// Package main ...
package main

import (
	"fmt"
	"log"
	"net/http"
)

// InMemoryPlayerStore ...
type InMemoryPlayerStore struct{}

// GetPlayerScore ...
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// RecordWin ...
func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	// First iteration:
	// We need the `Handler` interface to be able to create a server.
	// Typically, we do that by creating a `struct` and make it implement the interface.
	//
	// However, the use-case for structs is for holding data but currently, we have no state,
	// so it doesn't feel right to be creating one.
	// http.HandlerFunc lets us avoid this: https://golang.org/pkg/net/http/#HandlerFunc.
	//
	// So we use this to wrap our `PlayerServer` so that it now conforms to the type `Handler`.
	// handler := http.HandlerFunc(PlayerServer)

	// We can pass PlayerServer as argument to http.ListenAndServe because it implements
	// ServeHTTP(http.ResponseWriter, http.Request) method.
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen to port 5000: %v", err)
	}
}

// ServeHTTP ...
//
// Third iteration:
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

// processWin ...
func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// showScore ...
func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)

	// Write status not found for missing players.
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
