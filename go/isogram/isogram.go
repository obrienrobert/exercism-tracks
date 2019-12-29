// Package isogram contains the solution to isogram exercise of the Go track.
package isogram

import (
	"strings"
)

// IsIsogram returns true if the given string is an isogram
func IsIsogram(input string) bool {

	s := strings.Replace(input, "-", "", -1)
	s = strings.Replace(s, " ", "", -1)
	split := strings.Split(s, "")

	for i := range s {

		for j := i + 1; j < len(s); j++ {
			if strings.ToLower(split[i]) == strings.ToLower(split[j]) {
				return false
			}
		}
	}

	return true
}
