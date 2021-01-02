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
	isburglary := ProbGivenCondition(burglary, called)
	t.Log(ExpectedValueWithConfidence(isburglary))
	if Equals(isburglary, NewConstant(1.0)).Pr() {
		t.Error("Burglary is abnormally true")
	}
	if !ProbTrueAtLeast(Equals(isburglary, NewConstant(0.0)), 0.9) {
		t.Error("Burglary is too likely")
	}
}
