package uncertainty

// numCompareSamples is the number of samples to materialize to generate a comparison.
// Follows the paper in choice of value.
// TODO(barakmich): maybe worth exposing in some broader way.
const numCompareSamples = 10_000

type compareFunc func(x float64, y float64) bool

// LessThan returns a Bernoulli distribution
// where the probability of a 1.0 is reflected by
// how often a < b.
func LessThan(a Uncertain, b Uncertain) *Bernoulli {
	return comparison(a, b, func(x, y float64) bool {
		return x < y
	})
}

func GreaterThan(a Uncertain, b Uncertain) *Bernoulli {
	return comparison(a, b, func(x, y float64) bool {
		return x > y
	})
}

func NotEquals(a Uncertain, b Uncertain) *Bernoulli {
	return comparison(a, b, func(x, y float64) bool {
		return x != y
	})
}

func comparison(a Uncertain, b Uncertain, compare compareFunc) *Bernoulli {
	count := 0.0
	for i := 0; i < numCompareSamples; i++ {
		val := 0.0
		x := a.sample()
		y := b.sample()
		if compare(x, y) {
			val = 1.0
		}
		count += val
	}
	return NewBernoulli(count / float64(numCompareSamples))
}
