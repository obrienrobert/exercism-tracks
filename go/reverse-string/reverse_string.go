// Package reverse contains the solution to reverse exercise of the Go track.
package reverse

// Reverse returns the reverse of the given string
func Reverse(input string) string {

	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
