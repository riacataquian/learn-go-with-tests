package racer

import (
	"net/http"
)

func Racer(a, b string) (faster string) {
	select { // lets you wait on multiple channels.
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
