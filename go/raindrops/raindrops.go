package raindrops

import (
	"strconv"
)

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
