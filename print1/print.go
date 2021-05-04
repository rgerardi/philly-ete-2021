package main

import "fmt"

type Jedi struct {
	FirstName string
	LastName  string
	Age       int
	Rank      string
}

func main() {
	luke := Jedi{
		FirstName: "Luke",
		LastName:  "Skywalker",
		Age:       22,
		Rank:      "Master",
	}

	fmt.Println(luke)
}
