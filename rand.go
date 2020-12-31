package uncertainty

import (
	"math/rand"
	"sync"
	"time"
)

var (
	globalRand     *rand.Rand
	globalRandLock sync.Mutex
)

func randFloat64() float64 {
	globalRandLock.Lock()
	defer globalRandLock.Unlock()
	if globalRand == nil {
		globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return globalRand.Float64()
}

func randNormalFloat64() float64 {
	globalRandLock.Lock()
	defer globalRandLock.Unlock()
	if globalRand == nil {
		globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return globalRand.NormFloat64()
}
