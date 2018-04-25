package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}

	Countdown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	if sleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, got %d want 4", sleeper.Calls)
	}
}
