package main

import (
	"fmt"
	"time"
)

// sheww ziyabuya !!
// so chanels allows us to sync processes without using a wait groups
// or mutext locks by using channels we follow the go  proverb :
// "Dont not communicate by sharing memory , share memory by communicating"
// we achive this syncing by pointing a value to a channel and listening
// for that value in s seperate part of the code ,
// NB: until the code listening for the value passed to the channel is done executing
//     the execution of the code that passed the value into the channel will be paused

func main() {
	// making a new channel that will transport an intiger
	channel := make(chan int)

	// execute first go routine
	go func() {
		for i := 0; i < 10; i++ {
			// pass the counter into the channel
			channel <- i
		}
	}()

	// execute second go routine
	go func() {
		for {
			// get the next value of the unbuffered channel then
			// print the value/int that was passed into the unbuffered channel
			fmt.Println(<-channel)
		}
	}()

	// one second is more than enough for us to see the values being printed out
	time.Sleep(time.Second)
}
