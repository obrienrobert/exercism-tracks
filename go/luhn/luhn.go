package luhn

import (
	"strings"
	"unicode"
)

// Valid validates if the supplied string is a valid Luhn number
// https://www.wikiwand.com/en/Luhn_algorithm#/Pseudocode_implementation
func Valid(sin string) bool {
	str := strings.ReplaceAll(sin, " ", "")

	// return if the string isn't of sufficient length
	if len(str) < 2 {
		return false
	}

	mod := len(str)%2 == 0
	total := 0

	for _, v := range str {
		if !unicode.IsDigit(v) {
			return false
		}
		v := int(v - '0')

		if mod {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}

		mod = !mod
		total += v
	}

	return total%10 == 0
}
