package shapes

import (
	"math"
)

// Shape is any structs that implements an Area() method.
type Shape interface {
	Area() float64
}

// Rectangle describes the properties of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a Rectangle given a height and width.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Triangle describes the properties of a triangle.
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a Triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Circle describes the properties of a circle.
type Circle struct {
	Radius float64
}

// Area calculates the area of a Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
