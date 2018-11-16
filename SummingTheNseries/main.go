package main

import (
	"fmt"
	"math"
)

func summingSeries(n int64) int32 {
	m := int64(math.Pow10(9) + 7)
	sn := ((n % m) * (n % m)) % m

	return int32(sn % (int64(math.Pow10(9)) + 7))

}

func main() {
	fmt.Println(summingSeries(2))
}
