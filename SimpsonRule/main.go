package main

import (
	. "codewarrior/SimpsonRule/kata"
	"fmt"
)

func main() {
	fmt.Println(Simpson(290))
	fmt.Println(Simpson(72))
	fmt.Println(Simpson(252))
	fmt.Println(Simpson(40))
}

// 290, 1.9999999986
// 72, 1.9999996367
// 252, 1.9999999975
// 40, 1.9999961668
