package main

// PlayerStore ...
type PlayerStore interface {
	GetPlayerScore(string) int
}

// PlayerServer ...
type PlayerServer struct {
	store PlayerStore
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

// GetPlayerScore ...
func (p *PlayerServer) GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
