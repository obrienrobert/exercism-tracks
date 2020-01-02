// Package isogram contains the solution to isogram exercise of the Go track.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if the given string is an isogram
func IsIsogram(input string) bool {

	input = strings.ToLower(input)

	for i, j := range input {
		if strings.ContainsRune(input[i+1:], j) && unicode.IsLetter(j) {
			return false
		}
	}

	return true
}