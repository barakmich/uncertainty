package uncertainty

import "fmt"

type Samples struct {
	Samples []*sample
	n       int
	i       int
}

func FromSamples(samples []*sample) *Samples {
	return &Samples{
		Samples: samples,
		n:       0,
		i:       newID(),
	}
}

func (s *Samples) sampleWithTrace() *sample {
	if len(s.Samples) == 0 {
		panic("Must have at least some samples in a sampling distribution")
	}
	if s.n >= len(s.Samples) {
		s.n = 0
	}
	out := s.Samples[s.n]
	s.n += 1
	out.addTrace(s.i, out.value)
	return out
}

func (s *Samples) sample() float64 {
	return s.sampleWithTrace().value
}

func (s *Samples) addSample(sample *sample) {
	s.Samples = append(s.Samples, sample)
}

func (s *Samples) Average() float64 {
	total := 0.0
	for _, v := range s.Samples {
		total += v.value
	}
	n := float64(len(s.Samples))
	return total / n
}

func (s *Samples) First() float64 {
	if len(s.Samples) == 0 {
		panic("No samples in the Sampling distribution")
	}
	return s.Samples[0].value
}

func (s *Samples) id() int {
	return s.i
}

func (s *Samples) String() string {
	var out string
	for i, v := range s.Samples {
		if i == 0 {
			out = v.String()
			continue
		}
		out = fmt.Sprintf("%s\n%s", out, v.String())
	}
	return out
}

func Materialize(u Uncertain, n int) *Samples {
	out := &Samples{
		i: newID(),
	}
	for i := 0; i < n; i++ {
		out.addSample(u.sampleWithTrace())
	}
	return out
}
