// Package accumulate contains the solution to accumulate exercise of the Go track.
package accumulate

// Accumulate applies the supplied function (converter) to supplied array elements (given)
func Accumulate(given []string, converter func(string) string) []string {

	var result []string

	for _, s := range given {
		result = append(result, converter(s))
	}

	return result
}
