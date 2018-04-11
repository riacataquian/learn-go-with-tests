package iteration

// Repeat repeats char 5 times.
func Repeat(char string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += char
	}

	return repeated
}
