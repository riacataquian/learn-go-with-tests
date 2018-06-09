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
	s := GetPlayerScore(player)
	fmt.Fprintf(w, s)
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
