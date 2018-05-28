package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		b := &bytes.Buffer{}

		Countdown(b, &CountdownOperationsSpy{})

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

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

// SpyTime allows spies to test against the durationSlept in place of a ConfigurableSleeper object.
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep is SpyTime's implementation of Sleeper interface.
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

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
