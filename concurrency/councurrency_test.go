package concurrency

import (
	"reflect"
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

	got := CheckWebsites(mockWebsiteChecker, urls)
	if len(urls) != len(got) {
		t.Fatalf("got %v, want %v", len(got), len(urls))
	}

	expected := map[string]bool{
		"https://google.com":         false,
		"https://github.com":         false,
		"https://workingexample.com": true,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
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
