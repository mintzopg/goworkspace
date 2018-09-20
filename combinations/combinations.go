package combinations

import "fmt"

// FindCombis function([]int, int, int)[][]int    Given an array of size n, generate and print all possible combinations of m elements in array without repetition.
func FindCombis(arr []int, n int, m int) [][]int {
	// Method used: Fix elements and Recur
	// define an array tthat will hold the arrays of combinations
	out := [][]int{}
	// define temporary array to hold currect combination
	data := make([]int, m)
	// define a func variable to use anonymous function with recursion
	var combinationUtil func(arr []int, data []int, start, end, idx, m int)

	// the recursive function
	combinationUtil = func(arr []int, data []int, start int, end int, idx int, m int) {
		// current combination is ready store it
		if idx == m {
			fmt.Printf("DATA: %v\n", data)
			out = append(out, append([]int(nil), data...))
			fmt.Printf("OUT: %v\n", out)
			return
		}
		/*	replace index with all possible elements. The condition "end-i+1 >= r-index" makes sure that including one element at index will make a combination with remaining elements at remaining positions
		 */
		for i := start; i <= end && end-i+1 >= m-idx; i++ {
			data[idx] = arr[i]
			combinationUtil(arr, data, i+1, end, idx+1, m)
		}
	}
	combinationUtil(arr, data, 0, n-1, 0, m)
	return out
}
