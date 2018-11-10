package main

import (
	"fmt"
	"strings"
)

func matchingStrings(strings []string, queries []string) []int32 {
	freq := make([]int32, len(queries))

	for i, query := range queries {
		contains(strings, query, freq, i)
	}
	return freq
}

func contains(sls []string, qs string, freq []int32, index int) {
	for _, s := range sls {
		if strings.Contains(s, qs) && len(s) == len(qs) {
			freq[index]++
			fmt.Printf("freq:= %v\n", freq)
		}
	}
}

func main() {
	strings := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}
	fmt.Println(matchingStrings(strings, queries))
}
