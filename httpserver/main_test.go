package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn-go-with-tests/httpserver/handlers"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	server := handlers.PlayerServer{NewInMemoryPlayerStore()}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	if response.Code != http.StatusOK {
		t.Errorf("it returns the correct status code\ngot %d, want %d", response.Code, http.StatusOK)
	}

	got := response.Body.String()
	want := "3"
	if got != want {
		t.Errorf("it returns the correct score for %s\ngot '%s', want '%s'", player, got, want)
	}
}

func newPostWinRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func newGetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}
