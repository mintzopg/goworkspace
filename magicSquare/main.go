package main

import "fmt"

func formMagicSquare(n int) {
	/* construct a 3x3 magic square with numbers from 1 to 9

	Three conditions hold:
	1. The position of next number is calculated by decrementing row number of previous number by 1, and incrementing the column number of previous number by 1.
	At any time, if the calculated row position becomes -1, it will wrap around to n-1. Similarly, if the calculated column position becomes n, it will wrap around to 0.

	2. If the magic square already contains a number at the calculated position, calculated column position will be decremented by 2, and calculated row position will be incremented by 1.

	3. If the calculated row position is -1 & calculated column position is n, the new position would be: (0, n-2).

	*/
	// initialize nxn, magic square to be
	magicSquare := make([][]int, n)
	for x := range magicSquare {
		magicSquare[x] = make([]int, n)
	}

	// initilaize positions i, j
	var i, j int = n / 2, n - 1

	// one by one put values in magic square
	for num := 1; num <= n*n; {
		if (i == -1) && (j == n) { // 3rd condition
			j = n - 2
			i = 0
		} else {
			// 1st condition helper, if next number goes to out of square's right size
			if j == n {
				j = 0
			}
			// 1st condition helper if next number goes out of square upper side
			if i < 0 {
				i = n - 1
			}
		}
		if magicSquare[i][j] != 0 { // 2nd condition
			j -= 2
			i++
			continue
		} else {
			magicSquare[i][j] = num
			num++
		}
		j++
		i--
	}

	// Print magic square
	fmt.Printf("The Magic Square for n=%d\nSum of each row or column %d:\n\n", n, n*(n*n+1)/2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%3d ", magicSquare[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	n := 3
	formMagicSquare(n)
}
