package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerHandler(t *testing.T) {
	tests := []struct {
		desc   string
		player string
		want   string
	}{
		{
			desc:   "returns Pepper's score",
			player: "Pepper",
			want:   "20",
		}, {
			desc:   "returns Floyd's score",
			player: "Floyd",
			want:   "10",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", test.player), nil)
			response := httptest.NewRecorder() // so we can spy on what is written on response.

			PlayerHandler(response, request)

			got := response.Body.String()

			if got != test.want {
				t.Errorf("PlayerHandler(_, _): %s\ngot '%s', want '%s'", test.desc, got, test.want)
			}
		})
	}
}
