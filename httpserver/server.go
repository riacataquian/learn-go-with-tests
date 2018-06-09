package main

import (
	"fmt"
	"net/http"
)

// PlayerServer ...
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	// ResponseWriter also implements io.Writer so we can use fmt.Fprintf to send strings
	// as HTTP responses.
	fmt.Fprintf(w, "20")
}
