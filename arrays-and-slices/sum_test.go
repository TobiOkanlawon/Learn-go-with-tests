package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("returns the sum of a collection", func(t *testing.T) {
		numbers := []int{1, 2}

		got := Sum(numbers)
		want := 3

		assertNumbersEqual(t, got, want, numbers)
	})
}

func TestSumEach(t *testing.T) {
	t.Run("returns the sum of each collection", func(t *testing.T) {
		got := SumEach([]int{1, 2, 3, 4}, []int{2, 3, 4, 5})
		want := []int{10, 14}

		assertDeepEquality(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("returns the sum of all but the heads of each collection", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
		want := []int{5, 7}

		assertDeepEquality(t, got, want)
	})

	t.Run("safely sums empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3, 4})
		want := []int{0, 9}

		assertDeepEquality(t, got, want)
	})
}

func assertNumbersEqual(t *testing.T, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d with arguments %v", got, want, numbers)
	}
}

func assertDeepEquality(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, expected %v", got, want)
	}
}
