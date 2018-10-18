package main

import (
	"fmt"
)

//  https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
// Code from Joseph

func main() {
	perms := [][]int{}
	for perm := range GeneratePermutations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		// fmt.Println(perm)
		perms = append(perms, perm)
	}
	fmt.Println(len(perms)) // 362880 = n! for P(n , k) = n! / (n-k)!
}

// GeneratePermutations; func returns the permutations of {1, 2, ..., n}, n integeres
func GeneratePermutations(data []int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}
func permutate(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}
