package main

import (
	"fmt"
	"gimin/nQueensProblem/nq"
)

func main() {
	n := 16
	problem := nq.NewQueen(n)
	fmt.Printf("%d-Queens problem\n", n)
	problem.Solve()

	// The algorithm is slow
}
