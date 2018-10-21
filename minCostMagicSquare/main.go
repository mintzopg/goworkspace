
package main

import (
	"fmt"
	"math"
)

const n = 3

func formingMagicSquare(s [][]int32) int32 {
	// s 3x3 matrix of integers in the inclusive range [1, 9]
	// generate vector in order
	v := [n*n]int{}
	index := 0
	for y:=0; y < n; y++{
		for x:= 0; x < n; x++{
			v[index] = int(s[y][x])
			index++
		}
	}
	//costs := [n*n]int{}
	cost := 1000
	validPerms := [][]int{} // valid permutations are the ones that form a magic square when put in a 3x3 matrix in row/column order

	// range over the channel and do something with the perms
	for perm := range generatePermutations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		if isMagic(perm) {
			validPerms = append(validPerms, perm)
		}
	}
	//fmt.Printf("%v \n", validPerms)
	index = 0
	for k, p := range validPerms{
		costperm:= 0
		for i, pval:= range p{
			costperm += int(math.Abs(float64(pval-v[i])))  // calculate the permutation cost
		}
		//costs[k]=costperm
		if costperm < cost {
			cost = costperm
			index = k
		}
	}
	fmt.Printf("min cost permutation: %v \n", validPerms[index])
	return int32(cost)
}


func generatePermutations(vec []int) <-chan []int {
	// Generator function returns a channel with the permutations
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, vec)
	}(c)
	return c
}

func permutate(c chan []int, inputs []int) {
	// compute the permutations
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

func isMagic(p []int) bool {
	square := [n][n]int{} // construct square
	// assign elements of slice with len = 9 to square
	// ROWS: y & COLUMNS: x
	k := 0
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			square[y][x] = p[k]
			k++
		}
	}

	// check if each row is a sum
	sum := 0
	for x := 0; x < n; x++ {
		sum += square[0][x]
	}
	for y := 1; y <= 2; y++ {
		tmp := 0
		for x := 0; x < 3; x++ {
			tmp += square[y][x]
		}
		if tmp != sum {
			return false
		}
	}

	// check if each column is same
	for x := 0; x <= 2; x++ {
		tmp := 0
		for y := 0; y < 3; y++ {
			tmp += square[y][x]
		}
		if tmp != sum {
			return false
		}
	}

	// check diagonal 1
	tmp := 0
	for i := 0; i < n; i++ {
		tmp += square[i][i]
	}
	if tmp != sum {
		return false
	}

	// check diagonal 2
	tmp = 0
	for i := 0; i < n; i++ {
		tmp += square[2-i][i]
	}
	if tmp != sum {
		return false
	}

	// default true
	return true
}

func main() {
	s1 := [][]int32{
		[]int32{4, 9, 2},
		[]int32{3, 5, 7},
		[]int32{8, 1, 5},
	}

	s2 := [][]int32{
		[]int32{4, 8, 2},
		[]int32{4, 5, 7},
		[]int32{6, 1, 6},
	}

	fmt.Println(formingMagicSquare(s1))  // 1 pass
	fmt.Println(formingMagicSquare(s2))	// 4 pass
}