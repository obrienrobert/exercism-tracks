// Package raindrops contains the solution to raindrops exercise of the Go track.
package raindrops

import (
	"strconv"
)

// Convert returns strings based on modulas of the given integer value
// 3 -> Pling
// 5 -> Plang
// 7 -> Plong
func Convert(raindrop int) string {

	sb := ""

	if raindrop%3 == 0 {
		sb += "Pling"
	}

	if raindrop%5 == 0 {
		sb += "Plang"
	}

	if raindrop%7 == 0 {
		sb += "Plong"
	}

	if sb == "" {
		return strconv.Itoa(raindrop)
	}

	return sb
}
