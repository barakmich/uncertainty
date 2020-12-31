package uncertainty

import "testing"

func TestExpectedValueGeneric(t *testing.T) {
	p := 0.5
	x := NewBernoulli(p)
	meanCI := ExpectedValueWithConfidence(x)
	t.Log(meanCI)
	Within(meanCI.Mean, p, epsilon)
}
