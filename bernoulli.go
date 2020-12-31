package uncertainty

import "math/rand"

type Bernoulli struct {
	probability float64
}

var _ Uncertain = &Bernoulli{}

func NewBernoulli(probability float64) *Bernoulli {
	if probability >= 1.0 {
		panic("Trying to create a bernoulli probability outside [0.0, 1.0))")
	}
	return &Bernoulli{
		probability: probability,
	}
}

func (b *Bernoulli) sample() float64 {
	r := rand.Float64()
	if r < b.probability {
		return 1.0
	}
	return 0.0
}
