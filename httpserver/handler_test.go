package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn-go-with-tests/httpserver/memstore"

	"github.com/kylelemons/godebug/pretty"
)

// StubPlayerStore captures and encapsulates data for testing.
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []memstore.Player
}

// GetPlayerScore is the StubPlayerStore implementation of PlayerStore interface.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

// RecordWin is the StubPlayerStore implementation of PlayerStore interface.
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []memstore.Player {
	return s.league
}

func TestGetPlayerScore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
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

func TestLeague(t *testing.T) {
	server := NewPlayerServer(&StubPlayerStore{})

	t.Run("it returns 200 on GET", func(t *testing.T) {
		req := newGetLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		// To parse JSON in our data model, we create a `Decoder` then call its `Decode` method.
		// To create a `Decoder`, it needs an `io.Reader` to read _from_ which in our case is response spy's `Body`.
		var got []memstore.Player
		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatalf("Unable to parse response from server: %s, '%v'", res.Body, err)
		}

		assertResponseStatus(t, res.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		want := []memstore.Player{
			{"Pins", 32},
			{"Pongpong", 20},
			{"Pingping", 14},
			{"Piupiu", 10},
		}
		store := &StubPlayerStore{nil, nil, want}
		server := NewPlayerServer(store)

		req := newGetLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := getLeagueFromResponse(t, res.Body)
		assertLeague(t, got, want)
	})

	t.Run("it has the proper response type", func(t *testing.T) {
		store := &StubPlayerStore{}
		server := NewPlayerServer(store)

		req := newGetLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		want := "application/json"
		if res.Header().Get("content-type") != want {
			t.Errorf("incorrect response type, got %s, want %s", res.HeaderMap, want)
		}
	})
}

func assertLeague(t *testing.T, got, want []memstore.Player) {
	t.Helper()

	if s := pretty.Compare(got, want); s != "" {
		t.Errorf("GET league: diff +want -got: %s", s)
	}
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

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []memstore.Player) {
	t.Helper()

	if err := json.NewDecoder(body).Decode(&league); err != nil {
		t.Fatalf("Unable to parse response from server: %s, '%v'", body, err)
	}

	return
}

func newGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return r
}

func newPostScoreRequest(name string) *http.Request {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return r
}
