package racer

import (
	"fmt"
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

// Racer determine the faster website between a and b url with a defaultTimeout.
func Racer(a, b string) (faster string, err error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

// ConfigurableRacer determine the faster website between a and b url given a defaultTimeout.
func ConfigurableRacer(a, b string, timeout time.Duration) (faster string, err error) {
	select { // lets you wait on multiple channels.
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // https://golang.org/pkg/time/#After
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// ping starts a goroutine that will send ch a signal that http.Get request is finished.
func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)

		// blocks until http.Get is done.
		ch <- true
	}()

	return ch
}
