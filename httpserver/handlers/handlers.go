package handlers

import (
	"fmt"
	"net/http"
)

// PlayerHandler ...
func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprintf(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	return "10"
}
