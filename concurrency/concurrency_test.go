package concurrency

import (
	"github.com/kylelemons/godebug/pretty"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://workingexample.com" {
		return true
	}

	return false
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://workingexample.com",
	}

	res := CheckWebsites(mockWebsiteChecker, urls)

	want := len(urls)
	got := len(res)
	if want != got {
		t.Fatalf("got %v, want %v", got, want)
	}

	expected := map[string]bool{
		"https://google.com":         false,
		"https://github.com":         false,
		"https://workingexample.com": true,
	}
	if s := pretty.Compare(expected, res); s != "" {
		t.Fatalf("CheckWebsites(_, %v): Diff -got +want:\n %s", urls, s)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true

}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"

	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
