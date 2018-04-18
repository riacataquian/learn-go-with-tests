package iteration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got '%s', want '%s'", repeated, expected)
	}
}

func TestSquareX(t *testing.T) {
	squares := SquareX([]int{1, 2, 3, 4, 5})
	expected := []int{1, 4, 9, 16, 25}

	if !reflect.DeepEqual(expected, squares) {
		t.Errorf("got %v, want %v", squares, expected)
	}
}

func TestSquare(t *testing.T) {
	squares := Square([]int{1, 2, 3, 4, 5})
	expected := []int{1, 4, 9, 16, 25}

	if !reflect.DeepEqual(expected, squares) {
		t.Errorf("got %v, want %v", squares, expected)
	}
}

// By default, Benchmark runs synchronously.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square([]int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSquareX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareX([]int{1, 2, 3, 4, 5})
	}
}

func ExampleRepeat() {
	repeated := Repeat("jar", 2)
	fmt.Println(repeated)
	// Output: jarjar
}
