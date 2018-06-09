package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// StubPlayerStore ...
type StubPlayerStore struct {
	// Lets us mock a data source for our players score.
	scores map[string]int
	// Lets us spy wins of our players.
	winCalls []string
}

// GetPlayerScore is the StubPlayerStore implementation of PlayerStore interface.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestPlayerServer(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
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

	t.Run("it returns accepted status code on POST", func(t *testing.T) {
		req := newPostScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseStatus(t, res.Code, http.StatusAccepted)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{store}

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		req := newPostScoreRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		want := 1
		if len(store.winCalls) != want {
			t.Errorf("got %d record wins, want %d", len(store.winCalls), want)
		}

		if store.winCalls[0] != player {
			t.Errorf("got %s player, want %s", store.winCalls[0], player)
		}
	})
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

// newGetScoreRequest creates a new GET http request.
func newGetScoreRequest(name string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return r
}

// newPostScoreRequest creates a new POST http request.
func newPostScoreRequest(name string) *http.Request {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return r
}
