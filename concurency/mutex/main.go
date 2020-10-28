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

// a mutex is an object that is create to allow multiple  threads
//	can take turns sharing the same resource such as a file
var mutex sync.Mutex

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
		mutex.Lock()
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(identifier, i, "Counter :", counter)
		mutex.Unlock()
	}

	// mark off that one process is done
	wg.Done()
}
