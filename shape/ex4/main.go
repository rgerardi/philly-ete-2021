package main

import (
	"fmt"

	"github.com/rgerardi/philly-ete-2021/shape/shape"
)

//START OMIT
type border interface {
	Perimeter() float64
}

type square struct {
	shape.Rectangle
}

func newSquare(side float64) square {
	return square{shape.Rectangle{side, side}}
}

func main() {
	shapes := make([]border, 3)

	shapes[0] = shape.Rectangle{10, 5}
	shapes[1] = newSquare(5)
	shapes[2] = shape.Circle{3}

	for _, s := range shapes {
		printBorder(s)
	}
}

//END OMIT
func printBorder(b border) {
	fmt.Println("Border length is", b.Perimeter())
}
