package main

import (
	"fmt"
)

// Hacker Rank problem

func minimumSwaps(arr []int32) int32 {
	var i int32 // stupid int32 type
	var swaps int32

	// declare map to hold in realt-time the current index of each entry in "arr"
	m := make(map[int32]int)
	for j := range arr {
		m[arr[j]] = j
	}
	// fmt.Println(m)

	// main logic
	for i = 1; i <= int32(len(arr)); i++ {
		x := arr[i-1]         // temp value needed to update map entries
		if m[i] == int(i-1) { // if no swap is needed continue
			continue
		}
		arr[m[i]], arr[i-1] = arr[i-1], arr[m[i]] // do the swap
		m[x] = m[i]                               // update the map with the changed indexes
		m[i] = int(i - 1)
		swaps++
		if isSorted(arr) { // if sorted now get out
			break
		}
	}

	return swaps
}

// need to implement this for []int32 type
func isSorted(data []int32) bool {
	n := len(data)
	for i := n - 1; i > 0; i-- {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func main() {
	// q := []int32{2, 3, 4, 1, 5}
	q := []int32{1, 3, 5, 2, 4, 6, 7}
	// q := []int32{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(isSorted(q))
	fmt.Println(minimumSwaps(q))
}
