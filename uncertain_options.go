package uncertainty

type OptionType int

const (
	sampleSizeOpt OptionType = iota
	zScoreOpt
	percentErrorOpt
)

type Option struct {
	optionType OptionType
	intVal     int
	floatVal   float64
}

func SampleSize(n int) Option {
	return Option{
		optionType: sampleSizeOpt,
		intVal:     n,
	}
}

func getSampleSize(opts []Option, def int) int {
	for _, v := range opts {
		if v.optionType == sampleSizeOpt {
			return v.intVal
		}
	}
	return def
}

const zScore95 = 1.96

func ZScore95() Option {
	return ZScore(zScore95)
}

const zScore99 = 2.58

func ZScore99() Option {
	return ZScore(zScore99)
}

const zScore90 = 1.64

func ZScore90() Option {
	return ZScore(zScore90)
}

func ZScore(score float64) Option {
	return Option{
		optionType: zScoreOpt,
		floatVal:   score,
	}
}

func getZScore(opts []Option, def float64) float64 {
	for _, v := range opts {
		if v.optionType == zScoreOpt {
			return v.floatVal
		}
	}
	return def
}

func PercentError(v float64) Option {
	return Option{
		optionType: percentErrorOpt,
		floatVal:   v,
	}
}

func getPercentError(opts []Option, def float64) float64 {
	for _, v := range opts {
		if v.optionType == percentErrorOpt {
			return v.floatVal
		}
	}
	return def
}
