package di

import (
	"fmt"
	"io"
)

// Greet accepts an io.Writer to write name to.
//
// Use os.Stdout, bytes.Buffer or even http.ResponseWriter.
func Greet(w io.Writer, name string) {
	// Fprintf is like Printf but instead takes a Writer to send the string to,
	// whereas Printf defaults to stdout.
	fmt.Fprintf(w, "Hello, %s", name)
}
