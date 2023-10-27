package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{duration: sleepTime, sleep: spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v, but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

type SpyOperations struct {
	CallSequence []string
}

func (s *SpyOperations) Sleep() {
	s.CallSequence = append(s.CallSequence, sleepOperation)
}

func (s *SpyOperations) Write(p []byte) (n int, err error) {
	s.CallSequence = append(s.CallSequence, printOperation)
	return 0, nil
}

const printOperation = "print"
const sleepOperation = "sleep"

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

	t.Run("it runs the operations in the correct order", func(t *testing.T) {
		spy := &SpyOperations{}

		writerMock := spy
		sleeperMock := spy

		Countdown(writerMock, sleeperMock)

		want := []string{
			"print",
			"sleep",
			"print",
			"sleep",
			"print",
			"sleep",
			"print",
		}

		if !reflect.DeepEqual(spy.CallSequence, want) {
			t.Errorf("got %v, but want %v", spy.CallSequence, want)
		}
	})
}
