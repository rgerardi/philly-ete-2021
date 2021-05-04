package methods // OMIT

// Methods START OMIT
type Jedi struct {
	FirstName string
	LastName  string
	Age       int
	Rank      string
}

func (j Jedi) ForcePush(p int) int { // HL
	// TODO Method implementation
	return p
} // HL

type power int

func (p power) unleash() { // HL
	// TODO Method implementation
}

// Methods END OMIT

type ForceUser interface {
	ForcePush(int) int
}
