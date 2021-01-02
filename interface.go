package uncertainty

import "sync"

type Uncertain interface {
	sample() float64
	sampleWithTrace() *sample
	id() int
}

type UncertainBool interface {
	Uncertain
	sampleBool() bool
	Pr() bool
}

type sample struct {
	value float64
	trace map[int]float64
}

func newSample(val float64) *sample {
	return &sample{
		value: val,
		trace: make(map[int]float64),
	}
}

func (s *sample) addTrace(id int, val float64) {
	s.trace[id] = val
}

func (s *sample) combine(other *sample) *sample {
	out := newSample(s.value)
	for k, v := range s.trace {
		out.addTrace(k, v)
	}
	for k, v := range other.trace {
		out.addTrace(k, v)
	}
	return out
}

var (
	idPrinter     int
	idPrinterLock sync.Mutex
)

func newID() int {
	idPrinterLock.Lock()
	defer idPrinterLock.Unlock()
	idPrinter++
	return idPrinter
}
