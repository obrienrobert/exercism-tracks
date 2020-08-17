package luhn

import (
	"strconv"
	"strings"
)

// Valid validates if the supplied string is a valid Luhn number
// https://www.wikiwand.com/en/Luhn_algorithm#/Pseudocode_implementation
func Valid(sin string) bool {
	str := strings.Replace(sin, " ", "", -1)

	// return if the string isn't of sufficient length
	if len(str) < 2 {
		return false
	}

	mod := len(str)%2 == 0
	total := 0

	for _, v := range str {
		val, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		if mod {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}

		mod = !mod
		total += val
	}

	return total%10 == 0
}
