package uncertainty

import "testing"

func TestGaussianMultiply(t *testing.T) {
	x := NewGaussian(5.0, 1.0)
	t.Log(ExpectedValueWithConfidence(x))
	y := NewGaussian(6.0, 1.0)
	t.Log(ExpectedValueWithConfidence(y))
	z := Mul(x, y)
	samplestats := ExpectedValueWithConfidence(z)
	t.Log(samplestats)
	mean := 30.0
	if mean < samplestats.Mean-samplestats.CI {
		t.Fatalf("True mean below sample mean interval")
	}

	if mean > samplestats.Mean+samplestats.CI {
		t.Fatalf("True mean above sample mean interval")
	}
}
