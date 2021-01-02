package uncertainty

import "testing"

func TestBernoulliMultinomial(t *testing.T) {
	x := NewBernoulli(0.6)
	m := NewMultinomial([]float64{0.0, 1.0}, []float64{0.4, 0.6})
	evx := ExpectedValueWithConfidence(x)
	evm := ExpectedValueWithConfidence(m)
	t.Log(evx)
	t.Log(evm)
	if !Within(evm.Mean, evx.Mean, epsilon) {
		t.Error("Multinomial doesn't track with similar Bernoulli")
	}
}

func TestDice(t *testing.T) {
	d6 := NewDice(6)
	b := GreaterThan(d6, NewConstant(4.1))
	if !Within(b.probability, 0.3333, epsilon) {
		t.Error("Probability of high rolls doesn't test")
	}
}
