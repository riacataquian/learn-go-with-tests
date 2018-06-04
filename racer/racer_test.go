package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Approach #1
// We don't want to be relying on external services to test our code because they can be:
// - Slow
// - Flaky
// - Can't test edge case
// func TestRacer(t *testing.T) {
// 	slow := "http://slow.com"
// 	fast := "http://fast.com"

// 	want := fast
// 	got := Racer(slow, fast)

// 	if got != want {
// 		t.Errorf("got '%s', want '%s'", got, want)
// 	}
// }

// Approach #2
// Let's change our tests to use mocks so we have reliable servers to test against that we can control.
//
// `httptest.NewServer` makes it easier to use it with testing,
// as it finds an open port to listen on and then you can close it when you're done with the test.
func TestRacer(t *testing.T) {
	t.Run("returns the fastest URL to reponse", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// Close the test servers so that it does not continue to listen to a port.
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error, but got one: %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server dosn't response within 10s", func(t *testing.T) {
		server := makeDelayedServer(20 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		if err != nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep to simulate a delay.
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
