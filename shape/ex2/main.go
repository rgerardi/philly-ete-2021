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
	c := shape.Circle{5}
	r := shape.Rectangle{20, 10}
	s := newSquare(5)

	printBorder(c)
	printBorder(r)
	printBorder(s)
}

//END OMIT
func printBorder(b border) {
	fmt.Println("Border length is", b.Perimeter())
}
