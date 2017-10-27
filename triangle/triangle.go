package triangle

import (
	"math"
)

const testVersion = 3

// Kind represents the type of triangle
type Kind int

// Triangle types
const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// KindFromSides determines what kind of triangle is given by three sides
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return 0
	}
	if a == b && b == c && a == c {
		return 1
	}
	if a == b || b == c || a == c {
		return 2
	}
	return 3
}

// helper function to determine if a triangle can be made from the three side values
func isTriangle(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	if a == 0 && b == 0 && c == 0 {
		return false
	}
	if (a+b) < c || (b+c) < a || (a+c) < b {
		return false
	}
	return true
}
