package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func retrievTests() ([]int32, []int32, []int32) {
	scores := []int32{}
	ali := []int32{}
	res := []int32{}

	var num int64
	var tt string

	// Read player scores
	file, err := os.Open("/home/gimin/dev/GoWorkspace/src/gimin/trackgolang/ClimbeLeaderBoard/tests/scores200k")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tt = scanner.Text()
		num, _ = strconv.ParseInt(tt, 10, 32)
		scores = append(scores, int32(num))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Read Alice scores
	file, err = os.Open("/home/gimin/dev/GoWorkspace/src/gimin/trackgolang/ClimbeLeaderBoard/tests/alice100k")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tt = scanner.Text()
		num, _ = strconv.ParseInt(tt, 10, 32)
		ali = append(ali, int32(num))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// read result
	file, err = os.Open("/home/gimin/dev/GoWorkspace/src/gimin/trackgolang/ClimbeLeaderBoard/tests/output200k100k")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		tt = scanner.Text()
		num, _ = strconv.ParseInt(tt, 10, 32)
		res = append(res, int32(num))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scores, ali, res

}

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

	// scores, alice, output := retrievTests()
	// fmt.Println(len(scores), len(alice), len(output)) // read success
	// result := climbingLeaderboard(scores, alice)
	// fmt.Println(len(result))
	// fmt.Println(len(output))

	// diff := []int32{}
	// for i, v := range result {
	// 	if v-output[i] != 0 {
	// 		// fmt.Println(i, v-output[i])
	// 		diff = append(diff, v-output[i])
	// 	}
	// }
	// fmt.Println(len(diff))
	// fmt.Println(result[139], output[139])
	// fmt.Println(alice[138], alice[139])

}
