package uncertainty

// Conditional creates a new distribution based on a boolean distribution where, if the first is true, then consider the input, otherwise not.
//
// This assumes the distributions are independent.
func Conditional(condition *Bernoulli, input Uncertain, opts ...Option) Uncertain {
	outputSamples := getSampleSize(opts, 10_000)
	out := &Samples{}
	for outputSamples != 0 {
		r := randFloat64()
		cond := r < condition.probability
		val := input.sample()
		if cond {
			out.addSample(val)
		}
		outputSamples--
	}
	return out
}
