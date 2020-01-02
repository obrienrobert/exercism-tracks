// Package diffsquares contains the solution to diffsquares exercise of the Go track.
package diffsquares

// SquareOfSum returns the square of the sum of the first ten natural numbers
func SquareOfSum(sum int) int {

	a := sum

	for i := 1; i < sum; i++ {
		a += i
	}
	return a * a
}

// SumOfSquares returns the sum of the squares of the first ten natural numbers
func SumOfSquares(sum int) int {

	b := 0

	for i := 1; i <= sum; i++ {
		b += i * i
	}
	return b
}

// Difference returns the difference between the values returned from SquareOfSum and SumOfSquares
func Difference(diff int) int {
	return SquareOfSum(diff) - SumOfSquares(diff)
}
