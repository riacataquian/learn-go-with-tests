package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	startCount = 3
	final      = "Go!"
	sleep      = "sleep"
	write      = "write"
)

// Sleeper describes a sleeper object.
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper interface with a defined delay.
type ConfigurableSleeper struct {
	// duration is the duration of its `Sleep` operation.
	duration time.Duration
	// sleep describes time.Sleep's function signature.
	// See TestConfigurableSleeper.
	sleep func(time.Duration)
}

// Sleep pause the timer given the duration.
//
// It is ConfigurableSleeper's implementation of the `Sleeper` interface.
func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

// Countdown counts down from `startCount` to `final` and repeatedly sleep given a duration.
//
// In `main` we will send to `os.Stdout` so users can see the printed output,
// in test we will send to `bytes.Buffer` for later capturing.
func Countdown(o io.Writer, s Sleeper) {
	for i := startCount; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(o, i)
	}

	s.Sleep()
	fmt.Fprintf(o, final)
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second, time.Sleep})
}
