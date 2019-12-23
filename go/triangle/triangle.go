package triangle

import (
	"math"
	"sort"
)

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {

	var arr []float64
	arr = append(arr, a, b, c)
	sort.Float64s(arr)

	if a == 0 || b == 0 || c == 0 || a < 0 || b < 0 || c < 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) || arr[0]+arr[1] < arr[2] {
		return NaT
	} else if a == b && b == c && a == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else if a != b && b != c && a != c {
		return Sca
	} else {
		return NaT
	}
}
