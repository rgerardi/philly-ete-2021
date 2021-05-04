package types // OMIT

import "fmt" // OMIT

var power int

type Jedi struct {
	FirstName string
	LastName  string
	Age       int
	Rank      string
}

func foo() {
	name := "Luke"    // string
	fmt.Println(name) // OMIT
}
