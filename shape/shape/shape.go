// STARTEX1 OMIT
package shape

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// ENDEX1 OMIT

// STARTEX5 OMIT
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// ENDEX5 OMIT
