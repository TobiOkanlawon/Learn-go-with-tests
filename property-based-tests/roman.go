package main

import (
	"errors"
	"strings"
)

type RomanNumerals []RomanNumeral

type RomanNumeral struct {
	Value        Arabic
	Symbol       string
	isSubtractor bool
}

type Arabic int

func NewArabic(number int) (Arabic, error) {
	if (number < 4000) && (number > 0) {
		return Arabic(number), nil
	} else {
		return 0, errors.New("invalid int to arabic conversion")
	}
}

func (r RomanNumerals) ValueOf(symbols ...byte) Arabic {
	symbol := string(symbols)

	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r RomanNumerals) CheckIsSubtractor(symbol byte) bool {
	for _, s := range r {
		if (s.Symbol == string(symbol)) && s.isSubtractor {
			return true
		}
	}
	return false
}

var allRomanNumerals = RomanNumerals{
	{1000, "M", false},
	{900, "CM", false},
	{500, "D", false},
	{400, "CD", false},
	{100, "C", true},
	{90, "XC", false},
	{50, "L", false},
	{40, "XL", false},
	{10, "X", true},
	{9, "IX", false},
	{5, "V", false},
	{4, "IV", false},
	{1, "I", true},
}

func ConvertToRoman(arabic Arabic) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) Arabic {
	var total Arabic

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if i+1 < len(roman) && allRomanNumerals.CheckIsSubtractor(symbol) {
			nextSymbol := roman[i+1]

			value, status := isTwoCharacterSymbol(symbol, nextSymbol)

			if status == true {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}

		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}

	return total
}

func isTwoCharacterSymbol(char1, char2 byte) (value Arabic, status bool) {
	val := allRomanNumerals.ValueOf(char1, char2)

	if val != 0 {
		return val, true
	}
	return 0, false
}

func main() {
	ConvertToArabic("III")
}
