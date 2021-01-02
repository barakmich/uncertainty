package uncertainty

type Uniform struct {
	lo, size float64
	i        int
}

func NewUniform(low, high float64) *Uniform {
	if high < low {
		panic("High value of range is lower than low value of range")
	}
	return &Uniform{
		lo:   low,
		size: high - low,
		i:    newID(),
	}
}

func (u *Uniform) sample() float64 {
	r := randFloat64()
	return r*u.size + u.lo
}

func (u *Uniform) id() int {
	return u.i
}

func (u *Uniform) sampleWithTrace() *sample {
	val := u.sample()
	s := &sample{
		value: val,
	}
	s.addTrace(u.i, val)
	return s
}
