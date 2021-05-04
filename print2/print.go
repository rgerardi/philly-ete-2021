package main

import "fmt"

// START OMIT
type Jedi struct {
	FirstName string
	LastName  string
	Age       int
	Rank      string
}

func (j Jedi) String() string {
	return j.FirstName + " " + j.LastName
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

// END OMIT
