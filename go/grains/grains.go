package grains

import (
	"errors"
	"math"
)

// Square calculates the number of grains on the given square (input)
func Square(input int) (uint64, error) {

	if input < 1 || input > 64 {
		return 0, errors.New("input must be between 1 and 64")
	}

	return 1 << (input - 1), nil
}

// Total returns the total possible number of grains that can exist on a chessboard (with 64 squares)
func Total() uint64 {
	return math.MaxUint64
}
