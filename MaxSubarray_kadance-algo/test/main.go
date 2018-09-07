package main

import (
	"fmt"
	"gimin/MaxSubarray_kadance-algo/maxsubarray"
)

func main() {
	var arr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	subarray, maxsum := maxsubarray.Sequence(arr)

	fmt.Printf("subarray: %v, maxsum: %d", subarray, maxsum)
}
