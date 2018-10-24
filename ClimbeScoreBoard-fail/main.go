package main

import (
	"fmt"
	"sort"
)

type player struct {
	score int32
	rank  int32
}

type rankings []player

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	res := []int32{} // keep alice position per game
	n := len(scores)

	// translate scores to type rankings data struct
	// initialize lraders slice of rankings
	players := rankings{
		player{
			score: scores[0],
			rank:  1,
		},
	}
	for i := 1; i < n; i++ {
		if scores[i] == players[i-1].score { // if they have the same score they have the same rank
			players = append(players, player{
				score: scores[i],
				rank:  players[i-1].rank,
			})
		} else {
			// scores[i] < scores[i-1] since scores is sorted descending
			players = append(players, player{
				score: scores[i],
				rank:  players[i-1].rank + 1,
			})
		}
	}
	// sort alice ascending
	sort.SliceStable(alice, func(i, j int) bool { return alice[i] < alice[j] })

	for _, sc := range alice {
		for j, pl := range players {
			if sc >= pl.score {
				res = append(res, pl.rank) // take this position
				rankIncrease(players, j)   // update the rest with rank + 1
				break
			} else {
				if j == n-1 { // if you reach the last element in players and sc < rank, then alice is last for this round
					res = append(res, pl.rank+1)
				}
				continue
			}
		}
	}
	return res
}

func rankIncrease(r rankings, i int) {
	for j := i; j < len(r); j++ {
		r[j].rank++
	}
}

func main() {
	var scores, alice []int32

	//test1 = OK
	// scores = []int32{100, 100, 50, 40, 40, 20, 10}
	// alice = []int32{5, 25, 50, 120}

	// result := climbingLeaderboard(scores, alice)
	// fmt.Println(result)
	// test2 = OK
	// scores = []int32{100, 90, 90, 80, 75, 60}
	// alice = []int32{50, 65, 77, 90, 102}

	// result = climbingLeaderboard(scores, alice)
	// fmt.Println(result)

	////////////////////////////////////////

	///////////////////////////////////////

	fmt.Println(len(scores), len(alice))
	result := climbingLeaderboard(scores, alice)
	fmt.Println(result)
}
