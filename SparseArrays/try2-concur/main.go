package main

import (
	"fmt"
	"strings"
	"sync"
)

func matchingStrings(strings []string, queries []string) []int32 {
	var wg sync.WaitGroup

	freq := make([]int32, len(queries))

	for i, query := range queries { // launch len(queries) number of concurrent go routines
		wg.Add(1)
		go contains(strings, query, freq, i, &wg)
	}
	wg.Wait()
	return freq
}

func contains(sls []string, qs string, freq []int32, index int, wg *sync.WaitGroup) {
	// fmt.Println("go routine #", runtime.NumGoroutine())
	for _, s := range sls {
		if strings.Contains(s, qs) && len(s) == len(qs) {
			freq[index]++ // freq[index] is not shared memory location => no race here
		}
	}
	wg.Done()
}

func main() {
	strings := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}
	fmt.Println(matchingStrings(strings, queries))
}
