package main

import (
	"net/http"

	"learn-go-with-tests/di/greeter"
)

// MyGreeterHandler sends `world` to http.ResponseWriter.
//
// http.ResponseWriter implements io.Writer.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greeter.Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
