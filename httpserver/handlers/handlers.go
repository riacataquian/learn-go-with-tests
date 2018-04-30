package handlers

import (
	"fmt"
	"net/http"
)

// PlayerStore describes the transaction allowed for the storage.
type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(name string)
}

// PlayerServer describes the HTTP server.
type PlayerServer struct {
	// Store is the object for storage.
	Store PlayerStore
}

// ServeHTTP process and serves HTTP requests and responses.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.getScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) getScore(w http.ResponseWriter, player string) {
	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
