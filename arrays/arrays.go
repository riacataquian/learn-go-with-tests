package array

// Sum returns the sum of number slices.
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// SumAll returns the sum of all number slices.
// Demonstrates the use of a variadic function.
func SumAll(numSlices ...[]int) []int {
	sums := make([]int, len(numSlices))
	for _, numbers := range numSlices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails calculates the totals of the "tails" of each slice.
// The tail of a collection is all the items apart from the first one (the "head").
func SumAllTails(numToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // get all except the 0th item.
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
