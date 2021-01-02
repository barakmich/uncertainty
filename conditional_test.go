package uncertainty

import "testing"

func TestBurglary(t *testing.T) {
	earthquake := Flip(0.001)
	burglary := Flip(0.01)
	alarm := Or(earthquake, burglary)
	phoneWorking := IfElse(earthquake, Flip(0.6), Flip(0.99)).ToBool()
	maryWakes := IfElse(
		And(alarm, earthquake),
		Flip(0.8),
		IfElse(alarm, Flip(0.6), Flip(0.2)),
	).ToBool()

	called := And(maryWakes, phoneWorking)
	// Conditionalprob
	avg := Materialize(called, 1000).Average()
	t.Log(avg)
}
