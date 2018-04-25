package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

// Countdown ....
func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}
