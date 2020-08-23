package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a string of nucleotides
type DNA string

// Counts generates a histogram of valid nucleotides, incrementing the corresponding
// nucleotide nucleotide value by 1 for every occurrence in the given DNA string
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = map[rune]int{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, c := range strings.Split(string(d), "") {

		if c == "A" {
			h['A'] += 1
		} else if c == "C" {
			h['C'] += 1
		} else if c == "G" {
			h['G'] += 1
		} else if c == "T" {
			h['T'] += 1
		} else {
			return h, errors.New("invalid nucleotides")
		}
	}

	return h, nil
}
