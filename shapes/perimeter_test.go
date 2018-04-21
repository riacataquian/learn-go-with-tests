package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		s    Shape
		want float64
	}{
		{Rectangle{12, 6}, 72.00},
		{Circle{10}, 314.16},
		// {Triangle{12, 6}, 36.0},
	}

	for _, test := range tests {
		got := test.s.Area()

		if math.Round(got) != math.Round(test.want) {
			t.Errorf("got %.2f, want %.2f", got, test.want)
		}
	}
}
