package kata

import (
	"math"
)

// Simpson function "Simpson(n int) float64" returns the approximation of an integral
func Simpson(n int) float64 {
	a := 0.0
	b := math.Pi
	h := (b - a) / float64(n)

	var s1, s2 float64

	// calculate s1
	for i := 1; i <= n/2; i++ {
		s1 += f(a + (2*float64(i)-1)*h)
	}

	// calculate s2
	for i := 1; i <= n/2-1; i++ {
		s2 += f(a + 2*float64(i)*h)
	}

	res := ((b - a) / (3 * float64(n))) * (f(a) + f(b) + 4*s1 + 2*s2)

	return res

}

func f(x float64) float64 {
	return (3.0 / 2.0) * math.Pow(math.Sin(x), 3)
	// if you simply put (3/2) at the beginning it fails because it makes 3/2 an integer
}
