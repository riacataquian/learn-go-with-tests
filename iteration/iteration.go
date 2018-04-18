package iteration

// Repeat repeats char 5 times.
func Repeat(char string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += char
	}

	return repeated
}

// Square accepts a variable number of ints then square each.
func Square(nums []int) []int {
	cap := len(nums)
	squares := make([]int, cap)

	for i, num := range nums {
		squares[i] = num * num
	}

	return squares
}

// SquareX accepts a variable number of ints then square each.
func SquareX(nums []int) []int {
	var squares []int

	for _, num := range nums {
		squares = append(squares, num*num)
	}

	return squares
}
