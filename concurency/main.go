package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	waitGroup.Add(2)
	go foo()
	go bar()
	waitGroup.Wait()
}

func foo() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("foo : %d\n", i)
		time.Sleep(time.Duration(time.Microsecond * 3))
	}

	waitGroup.Done()
}

func bar() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("bar : %d\n", i)
		time.Sleep(time.Duration(time.Microsecond * 4))
	}

	waitGroup.Done()
}
