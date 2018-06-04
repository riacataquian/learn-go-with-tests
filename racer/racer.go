package racer

import (
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

// Racer determine the faster website between a and b url with a defaultTimeout.
func Racer(a, b string) (faster string, err error) {
	aDuration := measureDuration(a)
	bDuration := measureDuration(b)

	if aDuration < bDuration {
		return a, nil
	}

	return b, nil
}

func measureDuration(url string) time.Duration {
	// time.Now to record just before we try and get the URL's response.
	start := time.Now()

	// http.Get returns an http.Response though we don't handle that now
	// since we don't care about it now.
	http.Get(url)

	// time.Since takes the start time and returns a time.Duration of the difference.
	return time.Since(start)
}

// ConfigurableRacer determine the faster website between a and b url given a defaultTimeout.
func ConfigurableRacer(a, b string, timeout time.Duration) (faster string, err error) {
	return "", nil
}

// ping starts a goroutine that will send ch a signal that http.Get request is finished.
// func ping(url string) chan bool {
func ping(url string) bool {
	return true
}
