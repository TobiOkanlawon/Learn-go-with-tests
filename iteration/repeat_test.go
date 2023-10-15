package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats a certain amount of time", func(t *testing.T) {

		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("defaults to 5 if supplied a non-positive number", func(t *testing.T) {
		repeated := Repeat("a", -100)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func ExampleRepeat() {
	repeatedString := Repeat("b", 5)
	fmt.Println(repeatedString)
	// Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
