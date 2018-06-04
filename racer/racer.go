package racer

import (
	"fmt"
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

// Approach #1: Synchronous
// func Racer(a, b string) (winner string) {
// 	aDuration := measureDuration(a)
// 	bDuration := measureDuration(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// Racer determine the faster website between a and b url with a defaultTimeout.
// Approach #2: Asynchronous
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

// ConfigurableRacer determine the faster website between a and b url given a defaultTimeout.
func ConfigurableRacer(a, b string, timeout time.Duration) (faster string, err error) {
	// In our concurrency chapter, it shows that we can wait for values to be sent to the channel with: `myVar := <-ch`, this is a _blocking_ call, as you are waiting for a value.
	// What `select` lets us do is wait on `multiple` channels.
	// The first one to send a value "wins" and the code underneath the case is executed.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		// time.After returns a `chan` and will send a signal down
		// it after the amount of time you defined.
		return "", fmt.Errorf("timed out waiting for '%s' and '%s'", a, b)
	}
}

// measureDuration returns the difference between the time http.Get started until the time it finished.
func measureDuration(url string) time.Duration {
	// time.Now to record just before we try and get the URL's response.
	start := time.Now()

	// http.Get returns an http.Response though we don't handle that now
	// since we don't care about it now.
	http.Get(url)

	// time.Since takes the start time and returns a time.Duration of the difference.
	return time.Since(start)
}

// ping starts a goroutine that will send ch a signal that http.Get request is finished.
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		// blocking until http.Get is done.
		http.Get(url)
		// sends a signal to the channel that we are finished.
		ch <- true
	}()
	return ch
}
