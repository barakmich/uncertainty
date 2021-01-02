package uncertainty

type arithmeticOperation struct {
	a, b    Uncertain
	i       int
	combine combineFunc
}

type combineFunc func(x float64, y float64) float64

func Add(a Uncertain, b Uncertain) Uncertain {
	return newArithmetic(a, b, func(x, y float64) float64 {
		return x + y
	})
}

func Sub(a Uncertain, b Uncertain) Uncertain {
	return newArithmetic(a, b, func(x, y float64) float64 {
		return x - y
	})
}

func Mul(a Uncertain, b Uncertain) Uncertain {
	return newArithmetic(a, b, func(x, y float64) float64 {
		return x * y
	})
}

func Div(a Uncertain, b Uncertain) Uncertain {
	return newArithmetic(a, b, func(x, y float64) float64 {
		return x / y
	})
}

func newArithmetic(a, b Uncertain, op combineFunc) *arithmeticOperation {
	return &arithmeticOperation{
		a:       a,
		b:       b,
		combine: op,
		i:       newID(),
	}
}

func (ar *arithmeticOperation) sampleWithTrace() *sample {
	as := ar.a.sampleWithTrace()
	bs := ar.b.sampleWithTrace()
	v := ar.combine(as.value, bs.value)
	s := as.combine(bs)
	s.value = v
	s.addTrace(ar.i, v)
	return s
}

func (ar *arithmeticOperation) sample() float64 {
	a := ar.a.sample()
	b := ar.b.sample()
	return ar.combine(a, b)
}

func (ar *arithmeticOperation) id() int {
	return ar.i
}
