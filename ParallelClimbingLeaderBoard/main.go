package main

import (
	"fmt"
	"sort"
)

func climbingLeaderboard(scores []int, alice []int) []int {
	res := []int{} // keep alice position per game

	generate := func() <-chan int {
		c := make(chan int)
		go func(c chan int) {
			defer close(c)
			placeAlice(c, scores, alice)
		}(c)
		return c
	}

	for val := range generate() {
		res = append(res, val)
	}

	// return result
	return res
}

func placeAlice(c chan int, scores []int, alice []int) {
	n := len(scores)
	m := make(map[int]int) // map to hold unique scores & rankings
	k := 1
	for _, score := range scores { // put scores in m
		if _, ok := m[score]; !ok {
			m[score] = k
			k++
		}
	}
	// start checking where Alice stands
	for _, sc := range alice {
		if _, ok := m[sc]; ok { // alice score is in map
			c <- m[sc]
			continue
		} else {
			if sc < scores[n-1] { // alice is last for this game
				c <- m[scores[n-1]] + 1
				continue
			}

		}
		// interpolate score
		i := sort.Search(len(scores), func(i int) bool { return sc > scores[i] })
		c <- m[scores[i]]
	}

}

func main() {

	//test1 = OK
	scores := []int{100, 100, 50, 40, 40, 20, 10}
	alice := []int{5, 25, 50, 120}

	result := climbingLeaderboard(scores, alice)
	fmt.Println(result)
	// test2 = OK
	scores = []int{100, 90, 90, 80, 75, 60}
	alice = []int{50, 65, 77, 90, 102}

	result = climbingLeaderboard(scores, alice)
	fmt.Println(result)

}
