package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// StubPlayerStore ...
type StubPlayerStore struct {
	scores map[string]int
}

// GetPlayerScore is the StubPlayerStore implementation of PlayerStore interface.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func TestPlayerServer(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseBody(t, res.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseBody(t, res.Body.String(), "10")

	})

	t.Run("returns 200 for players found", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseStatus(t, res.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("Apollo")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseStatus(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
	}
	server := &PlayerServer{store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseStatus(t, res.Code, http.StatusAccepted)
	})
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("PlayerServer(_, _): got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("PlayerServer(_, _): got '%s', want '%s'", got, want)
	}
}

// Create a new http request.
func newGetScoreRequest(name string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return r
}
