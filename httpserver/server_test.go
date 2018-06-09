package main

// This test file serves as the integration test for the entire server package.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn-go-with-tests/httpserver/memstore"

	"github.com/kylelemons/godebug/pretty"
)

func TestPlayerServer(t *testing.T) {
	store := memstore.New()
	server := NewPlayerServer(store)
	player := "Pepper"

	w := httptest.NewRecorder()
	postReq, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	server.ServeHTTP(w, postReq)
	server.ServeHTTP(w, postReq)
	server.ServeHTTP(w, postReq)

	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		getReq, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
		server.ServeHTTP(res, getReq)

		if res.Body.String() != "3" {
			t.Errorf("Player score: got '%s', want '%s'", res.Body.String(), "3")
		}
	})

	t.Run("get league", func(t *testing.T) {
		res := httptest.NewRecorder()
		getReq, _ := http.NewRequest(http.MethodGet, "/league", nil)
		server.ServeHTTP(res, getReq)

		var got []memstore.Player
		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatalf("Unable to parse response from server: %s, '%v'", res.Body, err)
		}
		want := []memstore.Player{
			{"Pepper", 3},
		}

		if s := pretty.Compare(got, want); s != "" {
			t.Errorf("League: diff +want -got: %s", s)

		}
	})
}
