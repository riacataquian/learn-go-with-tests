package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	start = 3
	final = "Go!"
	sleep = "sleep"
	write = "write"
)

// Sleeper is what every normal person is.
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper sleeps repeatedly every duration.
type ConfigurableSleeper struct {
	duration time.Duration
}

// Sleep is what every normal person does.
func (s *ConfigurableSleeper) Sleep() {
	time.Sleep(s.duration)
}

// Countdown counts down from `start` to `final` and repeatedly sleep given a duration.
func Countdown(o io.Writer, s Sleeper) {
	for i := start; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(o, i)
	}

	s.Sleep()
	fmt.Fprintf(o, final)
}

// Ops holds the series of the operations performed.
type Ops struct {
	Calls []string
}

// Sleep is.. luxury.
func (o *Ops) Sleep() {
	o.Calls = append(o.Calls, sleep)
}

// Write appends every operation performed to Ops.Calls.
//
// It accepts p []byte which implements io.Writer.
func (o *Ops) Write(p []byte) (n int, err error) {
	o.Calls = append(o.Calls, write)
	return
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second})
}
