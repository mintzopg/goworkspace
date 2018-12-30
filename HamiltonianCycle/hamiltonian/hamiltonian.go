package hamiltonian

import "fmt"

// Graph representation
type Graph struct {
	numVertices int     // number of vertices
	path        []int   // slice of nomdes that form a hamiltonian path
	adjMatrix   [][]int // g adjacency matrix representation
}

// NewGraph constructor function for the g
func NewGraph(adjMatrix [][]int) *Graph {
	g := &Graph{}
	g.adjMatrix = adjMatrix
	g.numVertices = len(adjMatrix[0])
	g.path = make([]int, g.numVertices)
	for i := 0; i < g.numVertices; i++ {
		g.path[i] = 0
	}
	return g
}

func (g Graph) showHamiltonianPath() {
	fmt.Println("A Hamiltonian cycle exists: ")
	for i := 0; i < g.numVertices; i++ {
		fmt.Printf("%d -> ", g.path[i])
	}
	fmt.Print(g.path[0]) // print last one to complete the cycle
}

// Algorithm function to do the algorithm
func (g Graph) Algorithm() {
	var findCycle func(v int) bool

	findCycle = func(v int) bool {

		if v == g.numVertices { //base case
			if g.adjMatrix[g.path[v-1]][g.path[0]] == 1 {
				return true
			}
			return false
		}

		for vertex := 1; vertex < g.numVertices; vertex++ {
			if g.isValid(vertex, v) {
				g.path[v] = vertex
				if findCycle(v + 1) {
					return true
				}
				// Backtracks here
			}
		}
		return false
	}

	g.path[0] = 0 // start from node 0 the search for a cycle
	if findCycle(1) {
		g.showHamiltonianPath()
		return
	} else {
		fmt.Println("No solution found!")
	}

}

func (g Graph) isValid(vertex, actual int) bool {
	// Are the two nodes connected?
	if g.adjMatrix[g.path[actual-1]][vertex] == 0 {
		return false
	}

	// Is this node already in Hamiltonian path?
	for i := 0; i < actual; i++ {
		if g.path[i] == vertex {
			return false
		}
	}
	return true
}
