package counter

import "sync/atomic"

type CASCounter struct {
	number uint64
}

func NewCASCounter() Counter {
	return &CASCounter{0}
}

func (c *CASCounter) Add(num uint64) {
	for {
		v := atomic.LoadUint64(&c.number)
		if atomic.CompareAndSwapUint64(&c.number, v, v+num) {
			return
		}
	}
}

func (c *CASCounter) Read() uint64 {
	return atomic.LoadUint64(&c.number)
}
