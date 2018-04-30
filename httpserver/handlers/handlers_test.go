package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	// scores map[string]int
	scores map[string]string
}

func (s *StubPlayerStore) GetPlayerScore(name string) string {
	score := s.scores[name]
	return score
}

func TestPlayerHandler(t *testing.T) {
	// desc holds the test suite description.
	var desc string
	store := &StubPlayerStore{
		map[string]string{
			"Pepper": "20",
			"Floyd":  "10",
		},
	}
	server := &PlayerServer{store}

	desc = "returns Pepper's score"
	t.Run(desc, func(t *testing.T) {
		request := newGetRequest("Pepper")
		response := httptest.NewRecorder() // so we can spy on what is written on response.

		server.ServeHTTP(response, request)

		want := "20"
		assertBodyResponse(t, desc, response.Body.String(), want)
	})

	desc = "returns Floyd's score"
	t.Run(desc, func(t *testing.T) {
		request := newGetRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "10"
		assertBodyResponse(t, desc, response.Body.String(), want)
	})

	desc = "returns 404 on missing players"
	t.Run(desc, func(t *testing.T) {
		request := newGetRequest("Some Missing Player")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := http.StatusNotFound
		assertStatus(t, desc, response.Code, want)
	})

	desc = "returns 200 if a player is found"
	t.Run(desc, func(t *testing.T) {
		request := newGetRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := http.StatusOK
		assertStatus(t, desc, response.Code, want)
	})
}

func TestStoreWins(t *testing.T) {
	var desc string
	store := &StubPlayerStore{
		map[string]string{},
	}
	server := &PlayerServer{store}

	desc = "it returns accepted on POST"
	t.Run(desc, func(t *testing.T) {
		request := newPostRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, desc, response.Code, http.StatusAccepted)
	})
}

func newPostRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func newGetRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func assertStatus(t *testing.T, desc string, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("PlayerHandler(_, _): %s\ngot %d, want %d", desc, got, want)
	}
}

func assertBodyResponse(t *testing.T, desc, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("PlayerHandler(_, _): %s\ngot '%s', want '%s'", desc, got, want)
	}
}
