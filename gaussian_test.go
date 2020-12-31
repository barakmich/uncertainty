package uncertainty

import "testing"

func TestGaussianExpectedValue(t *testing.T) {
	mean := 10.0
	x := NewGaussian(mean, 1.0)
	samplestats := ExpectedValueWithConfidence(x)
	t.Log(samplestats)
	if mean < samplestats.Mean-samplestats.CI {
		t.Fatalf("True mean below sample mean interval")
	}

	if mean > samplestats.Mean+samplestats.CI {
		t.Fatalf("True mean above sample mean interval")
	}
}
