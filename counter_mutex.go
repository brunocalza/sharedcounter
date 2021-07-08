package counter

import "sync"

type MutexCounter struct {
	mu     *sync.RWMutex
	number uint64
}

func NewMutexCounter() Counter {
	return &MutexCounter{&sync.RWMutex{}, 0}
}

func (c *MutexCounter) Add(num uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.number = c.number + num
}

func (c *MutexCounter) Read() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.number
}
