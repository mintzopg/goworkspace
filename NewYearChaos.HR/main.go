package main

import "fmt"

// Hacker Rank problem

func minimumBribes(q []int32) {
	q0 := make([]int32, len(q))
	var cost int32
	for i := range q {
		q0[i] = int32(i + 1)
	}
	for i, v := range q {
		// check unmoved
		if q0[i] == v {
			continue
		}
		diff := v - int32(i+1)
		if diff < 0 {
			continue
		}
		if diff >= 1 && diff <= 2 {
			cost += diff
		} else {
			fmt.Println("Too chaotic")
			return
		}
	}

	fmt.Println(cost)
}
func main() {
	// queue := []int32{2, 1, 5, 3, 4}    // OK (3)
	// queue := []int32{2, 5, 1, 3, 4}  // OK (Too chaotic)
	// queue := []int32{5, 1, 2, 3, 7, 8, 6, 4} // OK (Too chaotic)
	queue := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	minimumBribes(queue)
}
