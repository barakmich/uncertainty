package uncertainty

type Multinomial struct {
	values  []float64
	cutoffs []float64
	i       int
}

var _ Uncertain = &Multinomial{}

func NewMultinomial(values []float64, probabilities []float64) *Multinomial {
	multinomialEpsilon := 0.0001
	tally := 0.0
	cutoffs := make([]float64, len(values)-1)
	for i, p := range probabilities {
		tally += p
		if i == len(values)-1 {
			if !Within(tally, 1.0, multinomialEpsilon) {
				panic("Sum of probabilities for this multinomial do not add up to 1.0")
			}
		} else {
			cutoffs[i] = tally
		}
	}
	return &Multinomial{
		values:  values,
		cutoffs: cutoffs,
		i:       newID(),
	}
}

func NewEvenMultinomial(values []float64) *Multinomial {
	probs := make([]float64, len(values))
	for i := range probs {
		probs[i] = 1.0 / float64(len(values))
	}
	return NewMultinomial(values, probs)
}

func NewDice(sides int) *Multinomial {
	vals := make([]float64, sides)
	for i := 0; i < sides; i++ {
		vals[i] = float64(i + 1)
	}
	return NewEvenMultinomial(vals)
}

func (m *Multinomial) sample() float64 {
	r := randFloat64()
	for i, v := range m.cutoffs {
		if r < v {
			return m.values[i]
		}
	}
	return m.values[len(m.values)-1]
}

func (m *Multinomial) id() int {
	return m.i
}

func (m *Multinomial) sampleWithTrace() *sample {
	val := m.sample()
	t := newSample(val)
	t.addTrace(m.i, val)
	return t
}
