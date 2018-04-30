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
	store := StubPlayerStore{
		map[string]string{
			"Pepper": "20",
			"Floyd":  "10",
		},
	}
	server := &PlayerServer{&store}

	desc = "returns Pepper's score"
	t.Run(desc, func(t *testing.T) {
		request := newGetRequest("Pepper")
		response := httptest.NewRecorder() // so we can spy on what is written on response.

		server.ServeHTTP(response, request)

		want := "20"
		assertBodyResponse(t, response.Body.String(), desc, want)
	})

	desc = "returns Floyd's score"
	t.Run(desc, func(t *testing.T) {
		request := newGetRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "10"
		assertBodyResponse(t, response.Body.String(), desc, want)
	})

	desc = "returns 404 on missing players"
	t.Run(desc, func(t *testing.T) {
		request := newGetRequest("Some Missing Player")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound
		// assertBodyResponse(t, response.Code, desc, want)
		if got != want {
			t.Errorf("PlayerHandler(_, _): %s\ngot %d, want %d", desc, got, want)
		}
	})
}

func newGetRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func assertBodyResponse(t *testing.T, got string, desc string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("PlayerHandler(_, _): %s\ngot '%s', want '%s'", desc, got, want)
	}
}
