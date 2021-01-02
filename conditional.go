package uncertainty

type ConditionalDistribution struct {
	condition UncertainBool
	input     Uncertain
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
	t := ife.test.sampleBool()
	var s *sample
	if t {
		s = ife.trueBranch.sampleWithTrace()
	} else {
		s = ife.falseBranch.sampleWithTrace()
	}
	s.addTrace(ife.i, s.value)
	return s
}

func (ife *IfElseDistribution) sample() float64 {
	return ife.sampleWithTrace().value
}

func (ife *IfElseDistribution) id() int {
	return ife.i
}

func (ife *IfElseDistribution) ToBool() UncertainBool {
	return GreaterThan(ife, NewConstant(0.5))
}
