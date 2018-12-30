package main

import "gimin/HamiltonianCycle/hamiltonian"

func main() {
	// adjMat := [][]int{
	// 	{0, 1, 0}, {1, 0, 1}, {0, 1, 0},
	// }

	adjMat := [][]int{
		{0, 1, 1, 1, 0, 0}, {1, 0, 1, 0, 1, 0}, {1, 1, 1, 1, 0, 1}, {1, 0, 1, 0, 0, 1}, {0, 1, 0, 0, 0, 1}, {0, 1, 1, 1, 1, 1},
	}
	graph := hamiltonian.NewGraph(adjMat)
	graph.Algorithm()
}
