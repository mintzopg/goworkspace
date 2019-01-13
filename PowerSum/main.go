package main

import (
	"fmt"
)

// calculate the power of int num raised to n (num^n)
func powInt(num, n int32) int32 {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return powInt(num, n/2) * powInt(num, n/2)
	} else {
		return num * powInt(num, n/2) * powInt(num, n/2)
	}
}

// Find the number of ways that a given integer x, can be expressed as the sum of the n powers of unique, natural numbers.
func powerSum(x int32, n int32) int32 {
	var count func(x, n, num int32) int32

	count = func(x, n, num int32) int32 {
		val := x - powInt(num, n)
		// base case
		if val == 0 {
			return 1
		}
		if val < 0 {
			return 0
		}
		// num included or not
		return count(val, n, num+1) + count(x, n, num+1)
	}
	return count(x, n, 1)
}

func main() {
	fmt.Println(powerSum(13, 2))
	fmt.Println(powerSum(13, 3))
}
