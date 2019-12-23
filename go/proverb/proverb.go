// Package proverb contains the solution to proverb exercise of the Go track.
package proverb

// Proverb returns a string array containing a constructed proverb using the given values
func Proverb(rhyme []string) []string {

	var proverb []string

	if len(rhyme) == 0 {
		return rhyme
	}

	for i := 1; i < len(rhyme); i++ {
		proverb = append(proverb, "For want of a "+rhyme[i-1]+" the "+rhyme[i]+" was lost.")
	}

	for i := 0; i < 1; i++ {
		proverb = append(proverb, "And all for the want of a "+rhyme[i]+".")
	}

	return proverb
}
