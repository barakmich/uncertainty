package uncertainty

type Gaussian struct {
	mean   float64
	stddev float64
	i      int
}

var _ Uncertain = &Gaussian{}

func NewGaussian(mean, stddev float64) *Gaussian {
	return &Gaussian{
		mean:   mean,
		stddev: stddev,
		i:      newID(),
	}
}

func NewNormal(mean, stddev float64) *Gaussian {
	return NewGaussian(mean, stddev)
}

func (g *Gaussian) sample() float64 {
	r := randNormalFloat64()
	return (r * g.stddev) + g.mean
}

func (g *Gaussian) sampleWithTrace() *sample {
	val := g.sample()
	s := newSample(val)
	s.addTrace(g.i, val)
	return s
}

func (g *Gaussian) id() int {
	return g.i
}
