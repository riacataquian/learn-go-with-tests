package greeter

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// The `Buffer` type from the `bytes` package implements the `Writer` interface.
	//
	// Relevant docs:
	// https://golang.org/pkg/io/#Writer
	// https://golang.org/src/bytes/buffer.go, line 240
	b := bytes.Buffer{}
	Greet(&b, "Steve")

	got := b.String()
	want := "Hello, Steve"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
