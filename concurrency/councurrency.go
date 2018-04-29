package concurrency

// WebsiteChecker ...
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites ...
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	resCh := make(chan result)

	for _, url := range urls {
		go func(u string) {
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

		// Then update the results map.
		res[r.string] = r.bool
	}

	return res
}
