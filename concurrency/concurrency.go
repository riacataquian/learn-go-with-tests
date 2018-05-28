package concurrency

import (
	"fmt"
)

// WebsiteChecker is the function signature for checking if a website is working ok.
type WebsiteChecker func(string) bool

type result struct {
	string // example of an unnamed key, access it like: result.string
	bool
}

// CheckWebsites performs WebsiteChecker on urls.
//
// Returns a map of the websites checked and its state.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	resCh := make(chan result)

	// The order of the two for loops matter.
	// In the first for loop, we are sending to the resCh: resCh <- ...
	// In the second for loop, we are receiving from the resCh: r := <-resCh.
	//
	// A send operation on an unbuffered channel **blocks** the sending goroutine
	// until another goroutine executes a corresponding receive on the same channel,
	// at which point the value is transmitted and both goroutines may continue.
	//
	// Communication over an unbuffered channels causes the sending and receiving go routines to *synchronize*.

	for _, url := range urls {
		go func(u string) {
			fmt.Printf("sending %s..\n", u)
			// send statement: this uses the <- operator,
			// taking a channel on the left and a value on the right:
			resCh <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// receive statement: which assigns a value received from a channel to a variable.
		//
		// also uses the <- operator, but with the two operands now reversed:
		// the channel is now on the right and the variable we're assigning to is on the left:
		r := <-resCh

		fmt.Printf("received %s: %v..\n\n", r.string, r.bool)

		// Then update the results map.
		res[r.string] = r.bool
	}

	return res
}
