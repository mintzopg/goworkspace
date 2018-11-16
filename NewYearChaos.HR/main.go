package main

import "fmt"

// Hacker Rank problem

func minimumBribes(q []int32) {
	cost := 0
	// val: value of item in normal queue i.e. 1, 2, 3, 4, ...., n
	for val := len(q); val > 0; val-- {
		i := val - 1 // slice index
		if int32(val) == q[i] {
			q = q[:i]
			continue // no bribe
		} else {
			if int32(val) == q[i-1] {
				cost += 1
				q = append(q[:i-1], q[i:]...)
				continue
			} else if int32(val) == q[i-2] {
				cost += 2
				q = append(q[:i-2], q[i-1:]...)
				continue
			} else {
				fmt.Println("Too chaotic")
				return
			}
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
