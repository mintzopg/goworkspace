package main

import "fmt"

func queensAttack(n int32, k int32, qr int32, qc int32, obstacles [][]int32) int32 {
	var moves int32
	obstaclesMap := make(map[[2]int32]int32, k)
	directions := [][]int32{
		{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
	}
	for _, obs := range obstacles { // populate map with obstacles map[{obstacle.row, obstacle.col}] = 1
		obstaclesMap[[2]int32{obs[0], obs[1]}] = 1
	}

	// Algorithm
	for _, d := range directions { // for each direction
		row, col := qr, qc // start from queen's position
		for {              // increment per direction until limits reach
			row += d[0]
			col += d[1]
			if row > n || row < 1 || col > n || col < 1 { // check board limits and break
				break
			}
			if _, ok := obstaclesMap[[2]int32{row, col}]; ok { // check for obstacle and break
				break
			}
			moves++ // increase moves count
		}
	}
	return moves
}

func main() {
	var n, k, qr, qc int32

	// n = 8
	// k = 0
	// qr = 4
	// qc = 5
	// obstacles := [][]int32{{6, 3}, {8, 1}, {5, 5}, {6, 7}, {1, 5}}

	n = 4
	k = 0
	qr = 4
	qc = 4
	obstacles := [][]int32{}

	fmt.Println(queensAttack(n, k, qr, qc, obstacles))
}
