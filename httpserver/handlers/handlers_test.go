package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestPlayerHandler(t *testing.T) {
	var desc string

	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &PlayerServer{store}

	desc = "returns Pepper's score"
	t.Run(desc, func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder() // so we can spy on what is written on response.

		server.ServeHTTP(response, request)

		want := "20"
		assertResponseBody(t, desc, response.Body.String(), want)
	})

	desc = "returns Floyd's score"
	t.Run(desc, func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "10"
		assertResponseBody(t, desc, response.Body.String(), want)
	})

	desc = "returns 404 on missing players"
	t.Run(desc, func(t *testing.T) {
		request := newGetScoreRequest("Some Missing Player")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := http.StatusNotFound
		assertStatus(t, desc, response.Code, want)
	})

	desc = "returns 200 for valid players"
	t.Run(desc, func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := http.StatusOK
		assertStatus(t, desc, response.Code, want)
	})
}

func TestStoreWins(t *testing.T) {
	var desc string
	store := &StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{store}

	t.Run("it records win when POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d, want %d", len(store.winCalls), 1)
		}

		assertResponseBody(t, "it records win for player when POST", store.winCalls[0], player)
	})

	desc = "it returns http.StatusAccepted response"
	t.Run(desc, func(t *testing.T) {
		request := newPostWinRequest("Strange")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, desc, response.Code, http.StatusAccepted)
	})
}

func newPostWinRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func newGetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func assertStatus(t *testing.T, desc string, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("PlayerHandler(_, _): %s\ngot %d, want %d", desc, got, want)
	}
}

func assertResponseBody(t *testing.T, desc string, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("PlayerHandler(_, _): %s\ngot '%s', want '%s'", desc, got, want)
	}
}
