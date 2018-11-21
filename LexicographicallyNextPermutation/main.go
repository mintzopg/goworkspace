package main

import (
	"fmt"
)

func nextPermutation(str string) (string, bool) {
	// calculates in place the next lexicographically order permutation
	arr := []rune(str)
	// find longest non-increasing suffix
	i := len(arr) - 1
	for i > 0 && arr[i-1] >= arr[i] {
		i--
	} // Now i is the head index of the suffix

	// Are we at the last permutation?
	if i <= 0 {
		return "No answer", false
	}

	// Let arr[i-1] be the pivot
	// find rightmost element that exceeds the pivot
	j := len(arr) - 1
	for arr[j] <= arr[i-1] {
		j--
	}
	// Now the value of arr[j] will become the new pivot
	// Assertion: j >= i

	// swap pivot with j
	arr[i-1], arr[j] = arr[j], arr[i-1]

	// reverse the suffix
	j = len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	// succesfully computed next permutation
	return string(arr), true
}

func main() {
	str := "bbbc"

	next, ok := nextPermutation(str)
	fmt.Println(ok, next)

}
