package counter

import (
	"math"
	"math/big"
	"sync"
	"testing"
)

func testFloatCorrectness(t *testing.T, counter *CASFloatCounter) {
	wg := &sync.WaitGroup{}
	counter.Add(0.8)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter *CASFloatCounter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter *CASFloatCounter) {
				counter.Add(1.1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter *CASFloatCounter) {
				counter.Add(2.3)
				wg.Done()
			}(counter)
		}
	}

	wg.Wait()
	// 2.3 * 33 + 1.1 * 33  + 0.8= 113
	if big.NewFloat(math.Round(counter.Read())).Cmp(big.NewFloat(113)) != 0 {
		t.Errorf("counter should be %d and was %f", 113, counter.Read())
	}
}

func TestCASFloatCounter(t *testing.T) {
	testFloatCorrectness(t, NewCASFloatCounter())
}
