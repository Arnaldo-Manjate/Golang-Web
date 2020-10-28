package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// notice the change to an int64
var counter int64
var wg sync.WaitGroup

// atomicity allows you to share one specific resource with threads

// example of a race condition
func main() {
	wg.Add(2)
	go incrementer("foo")
	go incrementer("bar")

	fmt.Println("Final Counter :", counter)
	// wait for all wait groups to finish
	wg.Wait()
}

func incrementer(identifier string) {
	for i := 0; i < 20; i++ {
		// lock the process

		// x := counter
		// x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// update the counter using atomicity
		atomic.AddInt64(&counter, 1)
		fmt.Println(identifier, i, "Counter :", counter)
	}

	// mark off that one process is done
	wg.Done()
}
