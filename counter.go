package counter

type Counter interface {
	Add(uint64)
	Read() uint64
}

type FloatCounter interface {
	Add(float64)
	Read() float64
}
