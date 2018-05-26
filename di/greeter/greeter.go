package greeter

import (
	"fmt"
	"io"
)

// Greet accepts an io.Writer to write name to.
//
// Note, our function doesn't need to care where or how the printing happens,
// so we should accept a general-purporse interface rather than a concrete type.
func Greet(w io.Writer, name string) {
	// Fprintf is like Printf but instead takes a Writer to send the string to,
	// whereas Printf defaults to stdout.
	//
	// It allows you to pass in an `io.Writer` which is both implemented by
	// `os.Stdout` and `bytes.Buffer`.
	fmt.Fprintf(w, "Hello, %s", name)
}
