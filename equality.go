package uncertainty

// numCompareSamples is the number of samples to materialize to generate a comparison.
// Follows the paper in choice of value.
// TODO(barakmich): maybe worth exposing in some broader way.
const numCompareSamples = 10_000

type compareFunc func(x float64, y float64) bool

type comparisonOperation struct {
	a, b     Uncertain
	i        int
	comparer compareFunc
}

// TODO(barakmich): These comparisons do a simple sampling comparison by
// default, however, as an extension for speed and correctness, it's relatively
// straightforward to type-inspect the uncertainty types and create a
// better/faster/more-accruate implementation that matches.

// LessThan returns a Bernoulli distribution
// where the probability of a 1.0 is reflected by
// how often a < b.
func LessThan(a Uncertain, b Uncertain) UncertainBool {
	return newComparison(a, b, func(x, y float64) bool {
		return x < y
	})
}

func GreaterThan(a Uncertain, b Uncertain) UncertainBool {
	return newComparison(a, b, func(x, y float64) bool {
		return x > y
	})
}

func NotEquals(a Uncertain, b Uncertain) UncertainBool {
	return newComparison(a, b, func(x, y float64) bool {
		return x != y
	})
}

func Equals(a Uncertain, b Uncertain) UncertainBool {
	return newComparison(a, b, func(x, y float64) bool {
		return x == y
	})
}

func newComparison(a Uncertain, b Uncertain, compare compareFunc) *comparisonOperation {
	return &comparisonOperation{
		a:        a,
		b:        b,
		i:        newID(),
		comparer: compare,
	}
}

func (comp *comparisonOperation) sampleBool() bool {
	return comp.comparer(comp.a.sample(), comp.b.sample())
}

func (comp *comparisonOperation) sample() float64 {
	return convertBoolSampleToFloat(comp.sampleBool())
}

func (comp *comparisonOperation) id() int {
	return comp.i
}

func (comp *comparisonOperation) sampleWithTrace() *sample {
	asample := comp.a.sampleWithTrace()
	bsample := comp.b.sampleWithTrace()
	out := asample.combine(bsample)
	out.value = convertBoolSampleToFloat(
		comp.comparer(
			asample.value,
			bsample.value,
		),
	)
	out.addTrace(comp.i, out.value)
	return out
}

func (comp *comparisonOperation) Pr() bool {
	return Pr(comp)
}
