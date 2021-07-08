package counter

import (
	"math"
	"sync/atomic"
)

type CASFloatCounter struct {
	number uint64
}

func NewCASFloatCounter() *CASFloatCounter {
	return &CASFloatCounter{0}
}

func (c *CASFloatCounter) Add(num float64) {
	for {
		v := atomic.LoadUint64(&c.number)
		newValue := math.Float64bits(math.Float64frombits(v) + num)
		if atomic.CompareAndSwapUint64(&c.number, v, newValue) {
			return
		}
	}
}

func (c *CASFloatCounter) Read() float64 {
	return math.Float64frombits(atomic.LoadUint64(&c.number))
}
