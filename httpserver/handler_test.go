package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// StubPlayerStore captures and encapsulates data for testing.
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

// GetPlayerScore is the StubPlayerStore implementation of PlayerStore interface.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

// RecordWin is the StubPlayerStore implementation of PlayerStore interface.
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayerScore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := NewPlayerServer(store)

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

func TestPostPlayerScore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := NewPlayerServer(store)

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

	t.Run("it returns accepted status code on POST", func(t *testing.T) {
		req := newPostScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseStatus(t, res.Code, http.StatusAccepted)
	})
}

func TestLeage(t *testing.T) {
	server := NewPlayerServer(&StubPlayerStore{})

	t.Run("it returns 200 on GET", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseStatus(t, res.Code, http.StatusOK)
	})
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("StatusCode: got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Response: got '%s', want '%s'", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return r
}

func newPostScoreRequest(name string) *http.Request {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return r
}
