package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	n := len(scores)
	res := []int32{}           // keep alice position per game
	m := make(map[int32]int32) // map to hold unique scores & rankings
	k := int32(1)
	for _, score := range scores { // put scores in m
		if _, ok := m[score]; !ok {
			m[score] = k
			k++
		}
	}
	// start checking where Alice stands
	for _, sc := range alice {
		if _, ok := m[sc]; ok { // alice score is in map
			res = append(res, m[sc])
			continue
		} else {
			if sc < scores[n-1] { // alice is last for this game
				res = append(res, m[scores[n-1]]+1)
				continue
			}

		}
		// interpolate score
		i := sort.Search(len(scores), func(i int) bool { return sc > scores[i] })
		res = append(res, m[scores[i]])
	}

	// return result
	return res
}

func main() {

	//test1 = OK
	scores := []int32{100, 100, 50, 40, 40, 20, 10}
	alice := []int32{5, 25, 50, 120}

	result := climbingLeaderboard(scores, alice)
	fmt.Println(result)
	// test2 = OK
	scores = []int32{100, 90, 90, 80, 75, 60}
	alice = []int32{50, 65, 77, 90, 102}

	result = climbingLeaderboard(scores, alice)
	fmt.Println(result)
}
