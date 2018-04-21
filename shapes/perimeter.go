package shapes

import (
	"math"
)

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a Rectangle given a height and width.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle ...
type Circle struct {
	Radius float64
}

// Shape ...
type Shape interface {
	Area() float64
}

// Area calculates the area of a Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculates the area of a Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
