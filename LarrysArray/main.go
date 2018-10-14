package main

import (
	"fmt"
)

func main() {
	arr := []int32{3, 1, 2}
	fmt.Println(larrysArray(arr)) //YES
	arr = []int32{1, 3, 4, 2}
	fmt.Println(larrysArray(arr)) //YES
	arr = []int32{1, 2, 3, 5, 4}
	fmt.Println(larrysArray(arr)) //NO

}

func larrysArray(A []int32) string {
	/* a sequence has an inversion if two of its elements are out of their natural order; the inversion number is the cardinality of inversion set. It is a common measure of the sortedness of a permutation or sequence.
	 */
	n := len(A)
	inv := 0 // number of inversions

	// O(N^2) loop used (small array)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if A[i] > A[j] {
				// there is an inversion => count it
				inv++
			}
		}
	}
	// decide upon inversions number: even => yes
	if inv%2 == 0 {
		return "YES"
	}
	return "NO"
}
