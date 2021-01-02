package uncertainty

type Constant struct {
	val float64
	i   int
}

var _ Uncertain = &Constant{}

func NewConstant(val float64) *Constant {
	return &Constant{
		val: val,
		i:   newID(),
	}
}

func (c *Constant) sample() float64 {
	return c.val
}

func (c *Constant) id() int {
	return c.i
}

func (c *Constant) sampleWithTrace() *sample {
	s := newSample(c.val)
	s.addTrace(c.i, c.val)
	return s
}
