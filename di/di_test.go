package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	b := bytes.Buffer{}
	Greet(&b, "Steve")

	got := b.String()
	want := "Hello, Steve"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
