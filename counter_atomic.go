package counter

import "sync/atomic"

type AtomicCounter struct {
	number uint64
}

func NewAtomicCounter() Counter {
	return &AtomicCounter{0}
}

func (c *AtomicCounter) Add(num uint64) {
	atomic.AddUint64(&c.number, num)
}

func (c *AtomicCounter) Read() uint64 {
	return atomic.LoadUint64(&c.number)
}
