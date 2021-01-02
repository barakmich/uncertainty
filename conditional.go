package uncertainty

type ConditionalDistribution struct {
	condition UncertainBool
	input     Uncertain
	i         int
}

func ProbGivenCondition(input Uncertain, condition UncertainBool) Uncertain {
	return &ConditionalDistribution{
		condition: condition,
		input:     input,
		i:         newID(),
	}
}

func (c *ConditionalDistribution) sampleWithTrace() *sample {
	for {
		s := c.condition.sampleWithTrace()
		inputval, ok := s.trace[c.input.id()]
		if !ok {
			panic("not in trace")
			///inputval = c.input.sample()
		}
		if convertFloatSampleToBool(s.value) {
			s.value = inputval
			s.addTrace(c.input.id(), inputval)
			return s
		}
	}
}

func (c *ConditionalDistribution) sample() float64 {
	return c.sampleWithTrace().value
}

func (c *ConditionalDistribution) id() int {
	return c.i
}

type IfElseDistribution struct {
	test        UncertainBool
	trueBranch  Uncertain
	falseBranch Uncertain
	i           int
}

func IfElse(condition UncertainBool, trueCond Uncertain, falseCond Uncertain) *IfElseDistribution {
	return &IfElseDistribution{
		test:        condition,
		trueBranch:  trueCond,
		falseBranch: falseCond,
		i:           newID(),
	}
}

func (ife *IfElseDistribution) sampleWithTrace() *sample {
	t := ife.test.sampleWithTrace()
	var s *sample
	if convertFloatSampleToBool(t.value) {
		s = ife.trueBranch.sampleWithTrace()
	} else {
		s = ife.falseBranch.sampleWithTrace()
	}
	return s.combine(t)
}

func (ife *IfElseDistribution) sample() float64 {
	return ife.sampleWithTrace().value
}

func (ife *IfElseDistribution) id() int {
	return ife.i
}

func (ife *IfElseDistribution) sampleBool() bool {
	return convertFloatSampleToBool(ife.sample())
}

func (ife *IfElseDistribution) Pr() bool {
	return Pr(ife.ToBool())
}

func (ife *IfElseDistribution) ToBool() UncertainBool {
	if _, ok := ife.trueBranch.(UncertainBool); ok {
		if _, ok := ife.falseBranch.(UncertainBool); ok {
			return ife
		}
	}
	return GreaterThan(ife, NewConstant(0.5))
}
