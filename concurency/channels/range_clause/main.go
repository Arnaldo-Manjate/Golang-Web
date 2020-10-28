package main

// notice we do not add time.Sleep and add the range clause
// to pass the current value in the channel to back and forth between go rutines

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}

		close(channel)
	}()

	// use range to reciev the value that is currently in the channel
	for number := range channel {
		println(number)
	}

}
