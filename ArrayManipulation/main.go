package main

import "fmt"

// using difference array of an array concept, to do some array manipulations

func arrayManipulation(n int32, queries [][]int32) int64 {
	A := make([]int, n)   // initial array of zeros for this case, to be updated
	D := make([]int, n+1) // this will also be an array of zeroes for this specific case

	for y := 0; y < len(queries); y++ {
		l := queries[y][0] - 1
		r := queries[y][1] - 1
		x := queries[y][2]
		update(D, int(l), int(r), int(x))
	}
	reconstruct(A, D)
	max := A[0]
	for _, val := range A[1:] { // O(n)
		if val > max {
			max = val
		}
	}
	return int64(max)
}

func reconstruct(A []int, D []int) { // O(n)
	n := len(A)
	for i := 0; i < n; i++ {
		if i == 0 {
			A[i] = D[i]
		} else {
			A[i] = D[i] + A[i-1]
		}
	}
	fmt.Println(A)
}

func update(D []int, l, r, x int) {
	D[l] += x
	D[r+1] -= x
}

func main() {
	var n int32 = 5
	queries := [][]int32{[]int32{1, 2, 100}, []int32{2, 5, 100}, []int32{3, 4, 100}}
	res := arrayManipulation(n, queries)
	fmt.Println(res)
}
