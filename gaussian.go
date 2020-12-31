package uncertainty

type Gaussian struct {
	mean   float64
	stddev float64
}

var _ Uncertain = &Gaussian{}

func NewGaussian(mean, stddev float64) *Gaussian {
	return &Gaussian{
		mean:   mean,
		stddev: stddev,
	}
}

func NewNormal(mean, stddev float64) *Gaussian {
	return NewGaussian(mean, stddev)
}

func (g *Gaussian) sample() float64 {
	r := randNormalFloat64()
	return (r * g.stddev) + g.mean
}
