package uncertainty

type Samples struct {
	samples []float64
	n       int
}

func FromSamples(samples []float64) *Samples {
	return &Samples{
		samples: samples,
		n:       0,
	}
}

func (s *Samples) sample() float64 {
	if len(s.samples) == 0 {
		panic("Must have at least some samples in a sampling distribution")
	}
	if s.n >= len(s.samples) {
		s.n = 0
	}
	out := s.samples[s.n]
	s.n += 1
	return out
}

func (s *Samples) addSample(sample float64) {
	s.samples = append(s.samples, sample)
}

func (s *Samples) Average() float64 {
	total := 0.0
	for _, v := range s.samples {
		total += v
	}
	n := float64(len(s.samples))
	return total / n
}

func (s *Samples) First() float64 {
	if len(s.samples) == 0 {
		panic("No samples in the Sampling distribution")
	}
	return s.samples[0]
}
