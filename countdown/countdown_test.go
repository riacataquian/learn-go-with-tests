package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	b := &bytes.Buffer{}

	Countdown(b, &ConfigurableSleeper{})

	got := b.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
