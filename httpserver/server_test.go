package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		// Create an http request.
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		// Spy on our response.
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("PlayerServer(_, _): got '%s', want '%s'", got, want)
		}
	})
}
