package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Adds integers as expected", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %d, expected %d", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
