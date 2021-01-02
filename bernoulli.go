package uncertainty

import (
	"fmt"
	"math"
)

type Bernoulli struct {
	probability float64
}

var _ Uncertain = &Bernoulli{}

func NewBernoulli(probability float64) *Bernoulli {
	if probability > 1.0 || probability < 0.0 {
		panic("Trying to create a bernoulli probability outside [0.0, 1.0], got " + fmt.Sprintf("%0.7f", probability))
	}
	return &Bernoulli{
		probability: probability,
	}
}

func (b *Bernoulli) sample() float64 {
	r := randFloat64()
	if r < b.probability {
		return 1.0
	}
	return 0.0
}

func (b *Bernoulli) Pr() bool {
	return b.ProbTrueAtLeast(0.5)
}

func (b *Bernoulli) ProbTrueAtLeast(prob float64, opts ...Option) bool {
	errorPercent := getPercentError(opts, 0.05)
	return b.sequentialProbabilityRatioTest(prob, errorPercent, 0.03, opts...)
}

// sequentialProbabilityRatioTest implements
// https://en.wikipedia.org/wiki/Sequential_probability_ratio_test.
// prob is the threshhold that this binary random variable has a true
// probability at least prob.
// confidence is the p value for how much error we accept (for 95% confidence, this is 5% or 0.05)
// indifference is the size of the indifference region (where we're not sure)
func (b *Bernoulli) sequentialProbabilityRatioTest(prob, confidence, indifference float64, opts ...Option) bool {
	maxSampleSize := getSampleSize(opts, 10_000)
	initSampleSize := 10
	sampleSizeStep := 10

	nSamples := 0

	h0 := prob - indifference
	h1 := prob + indifference

	alpha := confidence
	beta := confidence
	alphaLog := math.Log(beta / (1.0 - alpha))
	betaLog := math.Log((1.0 - beta) / alpha)

	k := 0

	wSum := 0.0
	wSumTrue := 0.0

	for nSamples = 0; nSamples < initSampleSize; nSamples++ {
		sample := b.sample()
		if sample == 1.0 {
			k += 1
			wSumTrue += b.probability
			wSum += b.probability
		} else {
			wSum += 1.0 - b.probability
		}
	}

	for nSamples <= maxSampleSize {
		logLikelihood := wSumTrue*math.Log(h1/h0) + (wSum-wSumTrue)*math.Log((1-h1)/(1-h0))
		if logLikelihood >= betaLog {
			return true
		}
		if logLikelihood <= alphaLog {
			return false
		}

		for i := 0; i < sampleSizeStep; i++ {
			sample := b.sample()
			if sample == 1.0 {
				k += 1
				wSumTrue += b.probability
				wSum += b.probability
			} else {
				wSum += 1.0 - b.probability
			}
		}
		nSamples += sampleSizeStep
	}

	// From the original implementation...
	//
	// If the maximum sample size is reached, assume the answer is false. This is an
	// (mostly unjustified) assumption that false positives are more damaging.
	//
	// It's an okay assumption, but compared to sample size steps, explaining as a
	// function input whether I'd like to return a false positive or a false negative
	// is perhaps more useful. Ultimately, though, this is a TODO.
	return false
}
