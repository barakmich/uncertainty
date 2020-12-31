package uncertainty

func Add(a Uncertain, b Uncertain) Uncertain {
	return &AddVariable{
		a: a,
		b: b,
	}
}

type AddVariable struct {
	a, b Uncertain
}

func (add *AddVariable) sample() float64 {
	aval := add.a.sample()
	bval := add.b.sample()
	return aval + bval
}

func Materialize(uncertain Uncertain, n int) *Samples {
	var floats []float64
	for i := 0; i < n; i++ {
		s := uncertain.sample()
		floats = append(floats, s)
	}
	return FromSamples(floats)
}
