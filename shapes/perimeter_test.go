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
		desc string
		s    Shape
		want float64
	}{
		{
			desc: "Calculate for the area of a Rectangle",
			s:    Rectangle{12, 6},
			want: 72.00,
		},
		{
			desc: "Calculate for the area of a Circle",
			s:    Circle{10},
			want: 314.16,
		},
		{
			desc: "Calculate for the area of a Triangle",
			s:    Triangle{12, 6},
			want: 36.0,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := test.s.Area()

			if math.Round(got) != math.Round(test.want) {
				t.Errorf("%s: Area() = %.2f, want %.2f", test.desc, got, test.want)
			}
		})
	}
}
