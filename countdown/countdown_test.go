package main

import (
	"bytes"
	"reflect"
	"testing"
)

// CountdownOperationsSpy holds the series of the operations performed on itself.
type CountdownOperationsSpy struct {
	// Calls be used to test against what is being performed on the object and its order.
	Calls []string
}

// Sleep appends `sleep` to its `Calls` field.
//
// It is CountdownOperationsSpy' implementation of the `Sleeper` interface.
func (o *CountdownOperationsSpy) Sleep() {
	o.Calls = append(o.Calls, sleep)
}

// Write appends every operation performed to `o`.
//
// It accepts p []byte which implements io.Writer.
func (o *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	o.Calls = append(o.Calls, write)
	return
}

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
		spyPrinter := &CountdownOperationsSpy{}
		Countdown(spyPrinter, spyPrinter)

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

		if !reflect.DeepEqual(want, spyPrinter.Calls) {
			t.Errorf("got %v want %v", want, spyPrinter.Calls)
		}
	})
}
