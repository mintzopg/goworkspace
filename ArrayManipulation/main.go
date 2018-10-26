package main

import (
	"fmt"
	"sort"
)

func arrayManipulation(n int, queries [][]int) int64 {
	// queries is a 2-D list of operation Arr[m][3]
	m := len(queries)
	v := make([]int, n)
	for i := range v {
		v[i] = 0
	}

	for y := 0; y < m; y++ {
		a := queries[y][0]
		b := queries[y][1]
		k := queries[y][2]
		// fmt.Println(a, b, k)
		for j := a - 1; j <= b-1; j++ {
			v[j] += k
		}
	}
	sort.Ints(v)
	fmt.Println(v)
	return int64(v[n-1])
}

func main() {
	n := 5
	arr := [][]int{[]int{1, 2, 100}, []int{2, 5, 100}, []int{3, 4, 100}}
	fmt.Println(arrayManipulation(n, arr))
}
