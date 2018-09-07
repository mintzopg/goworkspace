package kata_test

import (
	. "codewarrior/SimpsonRule/kata"
	"fmt"
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func assertFuzzyEquals(act float64, exp float64) {
	var inrange bool
	var merr float64 = 1e-10
	var e float64
	if exp == 0.0 {
		e = math.Abs(act)
	} else {
		e = math.Abs((act - exp) / exp)
	}
	inrange = (e <= merr)
	if inrange == false {
		fmt.Printf("Expected should be near: %1.10e , but got: %1.10e\n", exp, act)
	}
	Expect(inrange).To(Equal(true))
}
func dotest(n int, exp float64) {
	assertFuzzyEquals(Simpson(n), exp)
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases dist", func() {
		dotest(290, 1.9999999986)
		dotest(72, 1.9999996367)
		dotest(252, 1.9999999975)
		dotest(40, 1.9999961668)
	})
})
