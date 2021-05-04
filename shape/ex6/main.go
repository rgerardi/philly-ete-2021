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

//START OMIT
type surface interface { // HL
	Area() float64 // HL
} // HL

type surfaceBorder interface {
	border
	surface
}

func main() {
	shapes := make([]surfaceBorder, 3)
	shapes[0] = shape.Rectangle{10, 5}
	shapes[1] = newSquare(5)
	shapes[2] = shape.Circle{3}
	for _, s := range shapes {
		printBorder(s)
		printArea(s)
	}
}

func printArea(s surface) {
	fmt.Println("Surface area is", s.Area())
}

//END OMIT
func printBorder(b border) {
	fmt.Println("Border length is", b.Perimeter())
}
