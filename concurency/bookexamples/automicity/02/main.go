package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)
	fmt.Println("Starting execution")
	go doWork("routine A")
	go doWork("routine B")
	// Give the goroutines time to run.
	time.Sleep(1 * time.Second)

	// Safely flag it is time to shutdown.
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		// Do we need to shutdown. (read it safeley)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

/***** output  ******
Starting execution
Doing routine B Work
Doing routine A Work
Doing routine A Work
Doing routine B Work
Doing routine A Work
Doing routine B Work
Doing routine A Work
Doing routine B Work
Shutdown Now
Shutting routine B Down
Shutting routine A Down
*/
