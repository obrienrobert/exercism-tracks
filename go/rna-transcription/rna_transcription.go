package strand

import (
	"strings"
)

// ToRNA returns the RNA compliment of the given DNA string sequence
func ToRNA(dna string) string {

	rna := ""
	for _, c := range strings.Split(dna, "") {

		if c == "G" {
			rna += "C"
		} else if c == "C" {
			rna += "G"
		} else if c == "T" {
			rna += "A"
		} else if c == "A" {
			rna += "U"
		} else {
			rna += c
		}
	}

	return rna
}
