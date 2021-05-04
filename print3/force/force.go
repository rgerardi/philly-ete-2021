package force

type Jedi struct {
	FirstName string
	LastName  string
	Age       int
	Rank      string
}

func (j Jedi) String() string {
	return j.FirstName + " " + j.LastName
}
