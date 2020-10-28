package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// globals
var counter int
var wg sync.WaitGroup

// example of a race condition
func main() {
	wg.Add(2)
	go incrementer("foo")
	go incrementer("bar")

	// wait for all wait groups to finish
	wg.Wait()
}

func incrementer(identifier string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(identifier, i, "Counter :", counter)
	}

	// mark off that one process is done
	wg.Done()
}
