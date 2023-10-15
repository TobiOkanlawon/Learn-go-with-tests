package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

func TestCountdown(t *testing.T) {
	t.Run("outputs the correct text", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpySleeper{}

		Countdown(buffer, spy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("expected %q, but got %q", got, want)
		}

	})

	t.Run("it takes a second to sleep between tests", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpySleeper{}

		Countdown(buffer, spy)

		if spy.Calls != 3 {
			t.Errorf("Sleeper wasn't called appropriate amount of times (3), called %d", spy.Calls)
		}
	})
}
