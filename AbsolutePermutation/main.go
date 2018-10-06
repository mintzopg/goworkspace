package main

import "fmt"

func main() {
	// Absolute Permutation problem

	fmt.Printf("%v\n", absolutePermutation(100, 2))
	fmt.Printf("%v\n", absolutePermutation(10, 5))
	fmt.Printf("%v\n", absolutePermutation(10, 0))
	fmt.Printf("%v\n", absolutePermutation(6, 0))
	fmt.Printf("%v\n", absolutePermutation(2, 1))
	fmt.Printf("%v\n", absolutePermutation(3, 0))
	fmt.Printf("%v\n", absolutePermutation(3, 2))
}

func absolutePermutation(n int32, k int32) []int32 {
	cache := map[int32]string{}
	out := []int32{}

	checkCache := func(x int32) bool {
		// True if in cache False if not
		_, ok := cache[x]
		return ok
	}

	i := int32(1)
	for ; i <= n; i++ {
		// fmt.Println(i)
		// check i-k first
		if (i-k > 0) && (i-k <= n) && (!checkCache(i - k)) {
			out = append(out, i-k)
			cache[i-k] = "added"
		} else if (i+k <= n) && (!checkCache(i + k)) {
			out = append(out, i+k)
			cache[i+k] = "added"
		} else {
			return []int32{-1}
		}
	}
	return out
}
