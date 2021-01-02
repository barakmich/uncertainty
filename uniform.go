package uncertainty

type Uniform struct {
	lo, size float64
}

func NewUniform(low, high float64) *Uniform {
	if high < low {
		panic("High value of range is lower than low value of range")
	}
	return &Uniform{
		lo:   low,
		size: high - low,
	}
}

func (u *Uniform) sample() float64 {
	r := randFloat64()
	return r*u.size + u.lo
}
