package main

import (
	"fmt"

	"github.com/rgerardi/philly-ete-2021/print3/force"
)

func main() {
	luke := force.Jedi{
		FirstName: "Luke",
		LastName:  "Skywalker",
		Age:       22,
		Rank:      "Master",
	}

	fmt.Println(luke)
}
