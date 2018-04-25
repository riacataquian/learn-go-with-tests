package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
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
	})

	t.Run("sleep after every print", func(t *testing.T) {
		opsPrinter := &Ops{}
		Countdown(opsPrinter, opsPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, opsPrinter.Calls) {
			t.Errorf("got %v want %v", want, opsPrinter.Calls)
		}
	})
}
