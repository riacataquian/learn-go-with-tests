// Package main ...
package main

import (
	"fmt"
	"log"
	"net/http"
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

	// We can pass PlayerServer as argument to http.ListenAndServe because it implements
	// ServeHTTP(http.ResponseWriter, http.Request) method.
	server := &PlayerServer{}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen to port 5000: %v", err)
	}
}

// ServeHTTP ...
//
// Third iteration:
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
