package uncertainty

type Samples struct {
	Samples []float64
	n       int
}

func FromSamples(samples []float64) *Samples {
	return &Samples{
		Samples: samples,
		n:       0,
	}
}

func (s *Samples) sample() float64 {
	if len(s.Samples) == 0 {
		panic("Must have at least some samples in a sampling distribution")
	}
	if s.n >= len(s.Samples) {
		s.n = 0
	}
	out := s.Samples[s.n]
	s.n += 1
	return out
}

func (s *Samples) addSample(sample float64) {
	s.Samples = append(s.Samples, sample)
}

func (s *Samples) Average() float64 {
	total := 0.0
	for _, v := range s.Samples {
		total += v
	}
	n := float64(len(s.Samples))
	return total / n
}

func (s *Samples) First() float64 {
	if len(s.Samples) == 0 {
		panic("No samples in the Sampling distribution")
	}
	return s.Samples[0]
}
