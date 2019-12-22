// Package hamming contains the solution to exercise 3 of the Go track.
package hamming

import (
	"errors"
)

// Distance returns the hamming distance between 2 DNA strands
func Distance(a, b string) (int, error) {
	err := errors.New("Strands must be of same length")
	if len(a) != len(b) {
		return 0, err
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
