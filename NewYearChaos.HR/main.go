package main

import (
	"fmt"
	"math"
)

type person struct {
	t0 int32 // initial position
	ti int32 // position not
	bn int32 // number of bribes
}

func minimumBribes(q []int32) {
	var persons = []person{}

	for i, val := range q {
		j := int32(i) + 1
		persons = append(persons, person{
			t0: val,
			ti: j,
			bn: int32(math.Abs(float64(val - j))),
		})
	}
	fmt.Printf("persons: %v\n", persons)
}
func main() {
	queue := []int32{1, 2, 3, 5, 4, 7, 6, 8}
	minimumBribes(queue)
}
