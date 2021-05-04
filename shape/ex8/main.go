package main

import (
	"fmt"

	"github.com/rgerardi/philly-ete-2021/shape/shape"
)

type border interface {
	Perimeter() float64
}

type square struct {
	shape.Rectangle
}

func newSquare(side float64) square {
	return square{shape.Rectangle{side, side}}
}

type surface interface {
	Area() float64
}

type surfaceBorder interface {
	border
	surface
}

//START OMIT
func main() {
	shapes := make([]surfaceBorder, 3)
	shapes[0] = shape.Rectangle{10, 5}
	shapes[1] = newSquare(5)
	shapes[2] = shape.Circle{3}
	for _, s := range shapes {
		switch c := s.(type) { // HL
		case shape.Circle: // HL
			fmt.Println("Radius is:", c.Radius) // HL
		case shape.Rectangle: // HL
			fmt.Println("Length is:", c.Length) // HL
		default: // HL
			fmt.Println("Not a rectangle or circle") // HL
		} // HL
	}
}

//END OMIT

func printArea(s surface) {
	fmt.Println("Surface area is", s.Area())
}

func printBorder(b border) {
	fmt.Println("Border length is", b.Perimeter())
}
