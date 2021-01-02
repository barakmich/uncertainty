package uncertainty

type notOperation struct {
	b UncertainBool
	i int
}

func Not(booldist UncertainBool) UncertainBool {
	return &notOperation{booldist, newID()}
}

func (not *notOperation) Pr() bool {
	return Pr(not)
}

func (not *notOperation) sampleBool() bool {
	return !not.b.sampleBool()
}

func (not *notOperation) sample() float64 {
	return convertBoolSampleToFloat(not.sampleBool())
}

func (not *notOperation) sampleWithTrace() *sample {
	s := not.b.sampleWithTrace()
	s.value = 1.0 - s.value
	s.trace[not.i] = s.value
	return s
}

func (not *notOperation) id() int {
	return not.i
}

type orOperation struct {
	a, b UncertainBool
	i    int
}

func Or(a, b UncertainBool) UncertainBool {
	return &orOperation{a, b, newID()}
}

func (or *orOperation) sampleBool() bool {
	aval := or.a.sampleBool()
	bval := or.b.sampleBool()
	return aval || bval
}

func (or *orOperation) sample() float64 {
	return convertBoolSampleToFloat(or.sampleBool())
}

func (or *orOperation) sampleWithTrace() *sample {
	atrace := or.a.sampleWithTrace()
	btrace := or.b.sampleWithTrace()
	combined := atrace.combine(btrace)
	combined.value = 0.0
	if convertFloatSampleToBool(atrace.value + btrace.value) {
		combined.value = 1.0
	}
	combined.addTrace(or.i, combined.value)
	return combined
}

func (or *orOperation) Pr() bool {
	return Pr(or)
}

func (or *orOperation) id() int {
	return or.i
}

type andOperation struct {
	a, b UncertainBool
	i    int
}

func And(a, b UncertainBool) UncertainBool {
	return &andOperation{a, b, newID()}
}

func (and *andOperation) sampleBool() bool {
	aval := and.a.sampleBool()
	bval := and.b.sampleBool()
	return aval && bval
}

func (and *andOperation) Pr() bool {
	return Pr(and)
}

func (and *andOperation) sample() float64 {
	return convertBoolSampleToFloat(and.sampleBool())
}

func (and *andOperation) sampleWithTrace() *sample {
	atrace := and.a.sampleWithTrace()
	btrace := and.b.sampleWithTrace()
	combined := atrace.combine(btrace)
	combined.value = 0.0
	if convertFloatSampleToBool(atrace.value * btrace.value) {
		combined.value = 1.0
	}
	combined.addTrace(and.i, combined.value)
	return combined
}

func (and *andOperation) id() int {
	return and.i
}
