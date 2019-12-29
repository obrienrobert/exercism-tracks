// Package etl contains the solution to etl exercise of the Go track.
package etl

import (
	"fmt"
	"strings"
)

// Transform takes a list of letters and their corresponding scrabble scores and applies the score to each individual letter while also converting the each letter into lowercase
func Transform(input map[int][]string) map[string]int {

	m := make(map[string]int)

	for key, value := range input {
		fmt.Println("Key:", key, "Value:", value)

		for i := range value {
			m[strings.ToLower(value[i])] = key
		}
	}

	return m
}
