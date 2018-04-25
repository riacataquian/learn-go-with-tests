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

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (s *ConfigurableSleeper) Sleep() {
	time.Sleep(s.duration)
}

func Countdown(o io.Writer, s Sleeper) {
	for i := start; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(o, i)
	}

	s.Sleep()
	fmt.Fprintf(o, final)
}

type OperationsSpy struct {
	Calls []string
}

func (o *OperationsSpy) Sleep() {
	o.Calls = append(o.Calls, sleep)
}

func (o *OperationsSpy) Write(p []byte) (n int, err error) {
	o.Calls = append(o.Calls, write)
	return
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second})
}
