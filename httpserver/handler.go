package main

import (
	"net/http"

	"learn-go-with-tests/httpserver/memstore"
)

// PlayerStore describes a PlayerServer's persistence layer.
type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(string)
	GetLeague() []memstore.Player
}

// PlayerServer encapsulate the server's persistence layer.
// It also implements Handler interface to be able to start an HTTP server.
type PlayerServer struct {
	store PlayerStore
	// By embedding http.Handler, PlayerServer now has a ServeHTTP method.
	// Be careful tho, https://github.com/quii/learn-go-with-tests/blob/master/json.md#any-downsides.
	http.Handler
}

// Second iteration:
// PlayerServer ...
// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	// First iteration:
// 	// ResponseWriter also implements io.Writer so we can use fmt.Fprintf to send strings
// 	// as HTTP responses.
// 	// fmt.Fprintf(w, "20")

// 	player := r.URL.Path[len("/players/"):]
// 	s := GetPlayerScore(player)
// 	fmt.Fprintf(w, s)
// }

// GetPlayerScore retrieves a player's score.
// func (p *PlayerServer) GetPlayerScore(name string) string {
// 	if name == "Pepper" {
// 		return "20"
// 	}

// 	if name == "Floyd" {
// 		return "10"
// 	}

// 	return ""
// }
