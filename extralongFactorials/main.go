package main

import (
	"fmt"
	"math/big"
)

func main() {
	extraLongFactorials(30)
}

func extraLongFactorials(n int32) {
	// prints the result and returns
	// var result *big.Int
	// result.SetInt64(1)

	result := big.NewInt(1)
	cache := make(map[int]*big.Int)

	cache[1] = new(big.Int).SetInt64(1)

	for i := 1; i <= int(n); i++ {
		if _, ok := cache[i]; ok { // if factorial(i) is in cache -> cache[i]
			result.Mul(result, cache[i]) // result = result * cache[i]
		}
		result.Mul(result, new(big.Int).SetInt64(int64(i))) // else result = result * i
		cache[i] = result                                   // put it in cache
	}

	fmt.Println(result.String())
}
