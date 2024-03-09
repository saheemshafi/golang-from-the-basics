package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels for communication between go routines")

	channel := make(chan int)
	quit := make(chan int)

	waitGroup := sync.WaitGroup{}
	defer waitGroup.Wait()

	// go fibonacci(25, channel)

	// for value := range channel {
	// 	fmt.Println("Multiplication result:", value)
	// }

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()
	fibonacci(channel, quit)

}

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c) // useful to notify if no more values are coming like in a range block
// }

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // select statement can be used to listen on multiple channels
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quitting")
			return
		}
	}
}

/*
	Deadlock: A state when nobody is listening to the channel and
	we are trying to push something to it
	OR: We are reading the channel when nothing has been pushed to it
*/

/*
	Receiving values from channel is a blocking operation
*/
