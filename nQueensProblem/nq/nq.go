package nq

import (
	"fmt"
)

// Queens structure represent the N-Queen problem fields: {chesstable, num of queens}
type Queens struct {
	chessTable  [][]int
	numOfQueens int
}

// NewQueen constructor for Queens
func NewQueen(n int) *Queens {
	q := new(Queens)
	q.chessTable = make([][]int, n)
	for i := 0; i < n; i++ {
		q.chessTable[i] = make([]int, n)
		for j := 0; j < n; j++ {
			q.chessTable[i][j] = 0
		}
	}
	q.numOfQueens = n
	return q
}

// Solve function starts recursive algorithm
func (q Queens) Solve() {
	if ok := q.setQueens(0); ok { // steQueens starting fro position 0 returns true, we have a solution
		q.printQueens()
	} else {
		fmt.Println("There is no solution")
	}
}

func (q Queens) setQueens(colIndex int) bool {
	if colIndex == q.numOfQueens { // base case
		return true
	}
	for rowIndex := 0; rowIndex < q.numOfQueens; rowIndex++ {
		if q.isPlaceValid(rowIndex, colIndex) {
			q.chessTable[rowIndex][colIndex] = 1
			if q.setQueens(colIndex + 1) {
				return true
			}
			// BACKTRACKING starts at this point so need to reset square value to 0
			q.chessTable[rowIndex][colIndex] = 0
		}
	}
	return false
}

func (q Queens) isPlaceValid(rowIndex, colIndex int) bool {
	for i := 0; i < colIndex; i++ {
		if q.chessTable[rowIndex][i] == 1 {
			return false
		}
	}
	for i, j := rowIndex, colIndex; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if q.chessTable[i][j] == 1 {
			return false
		}
	}
	for i, j := rowIndex, colIndex; i < q.numOfQueens && j >= 0; i, j = i+1, j-1 {
		if q.chessTable[i][j] == 1 {
			return false
		}
	}

	return true
}

func (q Queens) printQueens() {
	for i := 0; i < len(q.chessTable); i++ {
		for j := 0; j < len(q.chessTable); j++ {
			if q.chessTable[i][j] == 1 {
				fmt.Print(" * ")
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Println()
	}
}
