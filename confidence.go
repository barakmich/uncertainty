package uncertainty

import (
	"fmt"
	"math"
)

type MeanAndConfidenceInterval struct {
	Mean float64
	CI   float64
}

func (mci MeanAndConfidenceInterval) String() string {
	return fmt.Sprintf("%0.4f +- %0.4f", mci.Mean, mci.CI)
}

func ExpectedValueWithConfidence(u Uncertain, opts ...Option) MeanAndConfidenceInterval {
	sampleSize := getSampleSize(opts, 1000)
	zScore := getZScore(opts, zScore95)

	m := Materialize(u, sampleSize)

	mean := m.Average()

	squaredError := 0.0
	for _, s := range m.Samples {
		squaredError += math.Pow(s.value-mean, 2.0)
	}
	sdev := math.Sqrt(squaredError / float64(sampleSize-1))

	// Since this is a sample standard deviation we're doing half of a t-test error
	// estimation using the square root of the number of samples to guide the rande
	// of the z scores.
	ci := zScore * sdev / math.Sqrt(float64(sampleSize))
	return MeanAndConfidenceInterval{
		Mean: mean,
		CI:   ci,
	}
}
