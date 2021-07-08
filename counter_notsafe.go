package counter

type NotSafeCounter struct {
	number uint64
}

func NewNotSafeCounter() Counter {
	return &NotSafeCounter{0}
}

func (c *NotSafeCounter) Add(num uint64) {
	c.number = c.number + num
}

func (c *NotSafeCounter) Read() uint64 {
	return c.number
}
