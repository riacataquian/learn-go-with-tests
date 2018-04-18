package array

// Sum returns the sum of number slices.
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// All returns the sum of all number slices.
// Demonstrates the use of a variadic function.
func All(numSlices ...[]int) []int {
	var sums []int

	for _, numbers := range numSlices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// AllTails calculates the totals of the "tails" of each slice.
// The tail of a collection is all the items apart from the first one (the "head").
func AllTails(numToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numToSum {
		tail := numbers[1:] // get all except the 0th item.
		sums = append(sums, Sum(tail))
	}

	return sums
}
