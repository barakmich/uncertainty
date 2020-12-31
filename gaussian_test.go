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

func TestGaussianSample(t *testing.T) {
	x := NewNormal(5.0, 2.0)
	m := Materialize(x, 100)
	for _, v := range m.Samples {
		if v < -3.0 || v > 13.0 {
			t.Error("Gaussian sample way out of range")
		}
	}
}

func TestGaussianMean(t *testing.T) {
	fails := 0
	for i := 0; i < 10; i++ {
		x := NewGaussian(5.0, 1.0)
		m := Materialize(x, 100)
		avg := m.Average()
		// If everything is working, this has about a 0.003% chance of a false positive
		// (99.9997% confidence interval with n=100, sigma=1.0 is +/- 0.4)
		t.Log(avg)
		if avg <= 4.6 || avg >= 5.4 {
			t.Log("Mean outside expected bounds (small chance of error)")
			fails += 1
		}
	}
	if fails > 1 {
		t.Error("Mean repeatedly outside expected bounds")
	}
}
