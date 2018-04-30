package handlers

import (
	"fmt"
	"net/http"
)

// PlayerStore ...
type PlayerStore interface {
	// GetPlayerScore(string) int
	GetPlayerScore(name string) string
	RecordWin(name string)
}

// PlayerServer ...
type PlayerServer struct {
	Store PlayerStore
}

// ServeHTTP ...
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.getScore(w, r)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) getScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.Store.GetPlayerScore(player)

	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, score)
}
