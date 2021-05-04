//START OMIT
package main

import (
	"fmt"

	"github.com/rgerardi/philly-ete-2021/shape/shape"
)

type border interface { // HL
	Perimeter() float64 // HL
} // HL

func main() {
	c := shape.Circle{5}
	r := shape.Rectangle{20, 10}

	printBorder(c)
	printBorder(r)
}

func printBorder(b border) { // HL
	fmt.Println("Border length is", b.Perimeter())
}

//END OMIT
