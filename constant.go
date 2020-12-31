package uncertainty

type Constant struct {
	val float64
}

var _ Uncertain = &Constant{}

func NewConstant(val float64) *Constant {
	return &Constant{
		val: val,
	}
}

func (c *Constant) sample() float64 {
	return c.val
}
