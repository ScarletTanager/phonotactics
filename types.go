package phonotactics

type Autosegment struct {
	Name string
}

type Slot struct {
	Previous     *Slot
	Next         *Slot
	Associations []*Autosegment
}

type Constraint struct {
	// How would we represent this?
	Weight int
	Name   string
}

type ConstraintSet interface {
	Sort()
	Constraints() []Constraint
}

type constraintSet struct {
}

type Engine interface {
	// Should have methods for processing inputs, etc.
}

type simpleConstraintEngine struct {
}
