package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT: multiple functions pulling from the same channel
	// ditribute across 10 workers
	myChannels := []<-chan int{
		factorial(in),
		factorial(in),
		factorial(in),
		factorial(in),
		factorial(in),
		factorial(in),
		factorial(in),
		factorial(in),
		factorial(in),
		factorial(in),
	}

	//FAN IN
	// merge multiple channels into one
	var y int // for tracking purpose
	for n := range merge(myChannels) {
		y++ // track number of numebrs
		fmt.Printf("%d --- %d\n", y, n)
	}
}

func gen() <-chan int { // puts 100... numbers in a channel
	out := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {
			for j := 30; j < 40; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int { //sq
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(cs))

	out := make(chan int) // out channle merges all

	// start an output go routine for each input channel in cs
	// output copies values from  c to out until c is closed
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range cs {
		go output(c)
	}

	// start a go routine to close out once all the output go routines are done.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
