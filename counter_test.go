package counter

import (
	"sync"
	"testing"
)

func testCorrectness(t *testing.T, counter Counter) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter Counter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter Counter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter Counter) {
				counter.Add(1)
				wg.Done()
			}(counter)
		}
	}

	wg.Wait()

	if counter.Read() != 66 {
		t.Errorf("counter should be %d and was %d", 66, counter.Read())
	}
}

func benchmark(b *testing.B, counter Counter, concurrency int) {
	b.StopTimer()
	start, end := &sync.WaitGroup{}, &sync.WaitGroup{}
	start.Add(1)
	for i := 0; i < concurrency; i++ {
		end.Add(1)
		go func(counter Counter) {
			start.Wait()
			counter.Add(1)
			counter.Read()
			end.Done()
		}(counter)
	}

	b.StartTimer()
	start.Done()
	end.Wait()
}

func TestNotSafeCounter(t *testing.T) {
	testCorrectness(t, NewNotSafeCounter())
}

func TestMutexCounter(t *testing.T) {
	testCorrectness(t, NewMutexCounter())
}

func TestChannelCounter(t *testing.T) {
	testCorrectness(t, NewChannelCounter())
}

func TestCASCounter(t *testing.T) {
	testCorrectness(t, NewCASCounter())
}

func TestAtomicCounter(t *testing.T) {
	testCorrectness(t, NewAtomicCounter())
}

func BenchmarkNotSafeCounter1(b *testing.B) {
	benchmark(b, NewNotSafeCounter(), 1)
}

func BenchmarkMutexCounter1(b *testing.B) {
	benchmark(b, NewMutexCounter(), 1)
}

func BenchmarkChannelCounter1(b *testing.B) {
	benchmark(b, NewChannelCounter(), 1)
}

func BenchmarkCASCounter1(b *testing.B) {
	benchmark(b, NewCASCounter(), 1)
}

func BenchmarkAtomicCounter1(b *testing.B) {
	benchmark(b, NewAtomicCounter(), 1)
}

func BenchmarkNotSafeCounter10(b *testing.B) {
	benchmark(b, NewNotSafeCounter(), 10)
}

func BenchmarkMutexCounter10(b *testing.B) {
	benchmark(b, NewMutexCounter(), 10)
}

func BenchmarkChannelCounter10(b *testing.B) {
	benchmark(b, NewChannelCounter(), 10)
}

func BenchmarkCASCounter10(b *testing.B) {
	benchmark(b, NewCASCounter(), 10)
}

func BenchmarkAtomicCounter10(b *testing.B) {
	benchmark(b, NewAtomicCounter(), 10)
}

func BenchmarkNotSafeCounter100(b *testing.B) {
	benchmark(b, NewNotSafeCounter(), 100)
}

func BenchmarkMutexCounter100(b *testing.B) {
	benchmark(b, NewMutexCounter(), 100)
}

func BenchmarkChannelCounter100(b *testing.B) {
	benchmark(b, NewChannelCounter(), 100)
}

func BenchmarkCASCounter100(b *testing.B) {
	benchmark(b, NewCASCounter(), 100)
}

func BenchmarkAtomicCounter100(b *testing.B) {
	benchmark(b, NewAtomicCounter(), 100)
}
