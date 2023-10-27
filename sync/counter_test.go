package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	assertCount := func(t testing.TB, got *Counter, want int) {
		t.Helper()

		if got.Value() != want {
			t.Errorf("expected counter value to be %d, got %d", want, got)
		}
	}

	t.Run("incrementing the counter works properly", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("works correctly concurrently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, want)
	})
}
