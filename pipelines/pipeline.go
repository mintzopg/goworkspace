package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// pipeline example

var data = make(map[int]bool)
var closech bool

var wg sync.WaitGroup
var mutex sync.Mutex

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// keep sending values to channel "a" until is closed
func first(min, max int, out chan<- int) {
	for {
		mutex.Lock()
		if closech {
			close(out)
			wg.Done()
			return
		}
		mutex.Unlock()
		out <- random(min, max)
	}
}

// receive data from in channel "a" coming from "first" and keep sending them to out channel "b", until a random number already exists in data
func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := data[x]
		if ok {
			mutex.Lock()
			closech = true // if found in data tell first() to close the channel
			mutex.Unlock()
		} else {
			data[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
	wg.Done()
}

// keeps reading from in channel "b"  and computes
func third(in <-chan int) {
	sum := 0
	for k := range in {
		sum += k
	}
	// when channel is empty print the result
	fmt.Printf("The sum of the random numbers is %d\n", sum)
	wg.Done()
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		os.Exit(1)
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	// n1 := 2
	// n2 := 10

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())
	a := make(chan int)
	b := make(chan int)

	wg.Add(3)
	go first(n1, n2, a)
	go second(b, a)
	go third(b)
	wg.Wait()

	// go run pipeline.go 100 10000
	// https://play.golang.org/p/MQmjyVDCrOm

}
