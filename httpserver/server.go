// Package main creates a server and accepts GET requests for a player's score
// and POST requests to increment a player's score.
//
// Perform a GET request:
// curl http://localhost:5000/players/Pepper
//
// Perform a POST request:
// curl -X POST http://localhost:5000/players/Pepper
package main

import (
	"fmt"
	"log"
	"net/http"

	"learn-go-with-tests/httpserver/memstore"
)

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

	s := memstore.New()
	// We can pass PlayerServer as argument to http.ListenAndServe because it implements
	// ServeHTTP(http.ResponseWriter, http.Request) method.
	server := &PlayerServer{s}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen to port 5000: %v", err)
	}
}

// ServeHTTP is PlayerServer's Handler interface implementation.
//
// Third iteration:
// Switch transaction between request methods.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

// processWin handle HTTP requests to increment a player's score.
func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// showScore handle HTTP requests to retrieve a player's score.
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	// Write status not found for missing players.
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
