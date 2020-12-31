package uncertainty

type arithmetic interface {
	Add(other Uncertain) Uncertain
	Sub(other Uncertain) Uncertain
	Mul(other Uncertain) Uncertain
	Div(other Uncertain) Uncertain
}

type equality interface {
	// Equals seems like a logical choice, but we're dealing with probablity
	// density functions. We know where there's no overlap under the curve, but we
	// don't know when there's exactly the same value
	NotEquals(other Uncertain) *Bernoulli
	LessThan(other Uncertain) *Bernoulli
	GreaterThan(other Uncertain) *Bernoulli
}

type Uncertain interface {
	sample() float64
}
