package handlers

import (
	"fmt"
	"net/http"
)

// PlayerStore ...
type PlayerStore interface {
	// GetPlayerScore(string) int
	GetPlayerScore(string) string
}

// PlayerServer ...
type PlayerServer struct {
	Store PlayerStore
}

// ServeHTTP ...
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.Store.GetPlayerScore(player)

	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, score)
}
