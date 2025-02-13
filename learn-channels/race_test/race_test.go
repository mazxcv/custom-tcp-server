package learnchannels_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestDataRaceConditionsMutex(t *testing.T) {
	var state int32
	var mu sync.RWMutex

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			state += int32(i)
			mu.Unlock()
		}(i)
	}
}

func TestDataRaceConditions2Atomic(t *testing.T) {
	var state int32
	// &state => memory address

	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}
