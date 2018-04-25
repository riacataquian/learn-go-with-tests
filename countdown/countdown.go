package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	start     = 3
	finalWord = "Go!"
)

// Sleeper ...
type Sleeper interface {
	Sleep()
}

// SpySleeper ...
type SpySleeper struct {
	Calls int
}

// Sleep ...
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// ConfigurableSleeper ...
type ConfigurableSleeper struct {
	duration time.Duration
}

// Sleep ...
func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

// Countdown ....
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := start; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second})
}
