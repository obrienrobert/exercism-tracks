// Package scrabble contains the solution to scrabble exercise of the Go track.
package scrabble

import (
	"strings"
)

// Score calculates and returns the scrabble score for the given string value
// A, E, I, O, U, L, N, R, S, T       1
// D, G                               2
// B, C, M, P                         3
// F, H, V, W, Y                      4
// K                                  5
// J, X                               8
// Q, Z                               10
func Score(word string) int {

	s := strings.Split(word, "")
	sum := 0

	for i := 0; i < len(s); i++ {
		s[i] = strings.ToLower(s[i])

		switch s[i] {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			sum = sum + 1
		case "d", "g":
			sum = sum + 2
		case "b", "c", "m", "p":
			sum = sum + 3
		case "f", "h", "v", "w", "y":
			sum = sum + 4
		case "k":
			sum = sum + 5
		case "j", "x":
			sum = sum + 8
		case "q", "z":
			sum = sum + 10
		}
	}

	return sum
}
