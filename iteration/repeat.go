package iteration

import (
	"strings"
)

const defaultAmount = 5

func Repeat(char string, amount int) string {
	if amount < 1 {
		amount = defaultAmount
	}
	return strings.Repeat(char, amount)
}
