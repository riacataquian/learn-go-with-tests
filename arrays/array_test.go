package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// Sub-test example:
	t.Run("collection of any size", func(t *testing.T) {
		got := Sum([]int{1, 2, 3})
		want := 6

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestAll(t *testing.T) {
	got := All([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAllTails(t *testing.T) {
	got := AllTails([]int{1, 2}, []int{0, 9}, []int{4, 5, 7})
	want := []int{2, 9, 12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
