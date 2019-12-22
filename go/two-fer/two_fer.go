// Package twofer contains the solution to exercise 2 of the Go track.
package twofer

// ShareWith returns 'one for <name>, one for me.' if the name variable is not empty
// otherwise 'one for you, one for me.' is returned.
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
