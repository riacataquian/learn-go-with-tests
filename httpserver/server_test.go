package main

// This test file serves as the integration test for the entire server package.

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	w := httptest.NewRecorder()
	postReq, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	server.ServeHTTP(w, postReq)
	server.ServeHTTP(w, postReq)
	server.ServeHTTP(w, postReq)

	res := httptest.NewRecorder()
	getReq, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	server.ServeHTTP(res, getReq)

	if res.Body.String() != "3" {
		t.Errorf("Response: got '%s', want '%s'", res.Body.String(), "3")
	}
}
