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

func Sub(a Uncertain, b Uncertain) Uncertain {
	return &SubVariable{
		a: a,
		b: b,
	}
}

type SubVariable struct {
	a, b Uncertain
}

func (sub *SubVariable) sample() float64 {
	aval := sub.a.sample()
	bval := sub.b.sample()
	return aval - bval
}

func Mul(a Uncertain, b Uncertain) Uncertain {
	return &MulVariable{
		a: a,
		b: b,
	}
}

type MulVariable struct {
	a, b Uncertain
}

func (mul *MulVariable) sample() float64 {
	aval := mul.a.sample()
	bval := mul.b.sample()
	return aval * bval
}

func Div(a Uncertain, b Uncertain) Uncertain {
	return &DivVariable{
		a: a,
		b: b,
	}
}

type DivVariable struct {
	a, b Uncertain
}

func (div *DivVariable) sample() float64 {
	aval := div.a.sample()
	bval := div.b.sample()
	return aval * (1.0 / bval)
}

func Materialize(uncertain Uncertain, n int) *Samples {
	var floats []float64
	for i := 0; i < n; i++ {
		s := uncertain.sample()
		floats = append(floats, s)
	}
	return FromSamples(floats)
}
