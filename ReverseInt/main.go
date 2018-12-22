package main

import (
	"errors"
	"fmt"
	"math"
)

// To Do: check for overflows and return 0 if such a case

func reverse(x int) (int, error) {
	err := errors.New("Integer overflow")

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0, err
	}

	sign := 1
	if x < 0 {
		sign *= -1
	}
	x *= sign
	revX := 0
	for x > 0 {
		revX = revX*10 + x%10
		x /= 10
	}
	return revX * sign, nil
}

func main() {
	x := math.MinInt32 + 1
	fmt.Println(reverse(x))
}
