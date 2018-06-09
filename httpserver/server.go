package main

import (
	"fmt"
	"net/http"
)

// PlayerServer ...
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	// First iteration:
	// ResponseWriter also implements io.Writer so we can use fmt.Fprintf to send strings
	// as HTTP responses.
	// fmt.Fprintf(w, "20")

	player := r.URL.Path[len("/players/"):]

	if player == "Pepper" {
		fmt.Fprintf(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprintf(w, "10")
		return
	}
}
