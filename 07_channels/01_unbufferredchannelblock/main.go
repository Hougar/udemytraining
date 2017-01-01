package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	// Slicces, Maps & Channels all use make
	// these three types represent references to data structures tha must be initialized before use
	// unbuffered channel adding a buffer would look like: c := make(chan int, 10) - can hold ten
	// more advanced technique best to avoid

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// code stops until something is takeen off the channel
		// un buffered channels block until sent and received
	}()
	// taking 0 to 9 and putting <- it on on to the channel c
	go func() {
		for {
			fmt.Println(<-c)
			// taking the numbers off the channel ie. print  whats on channel c
		}
	}()
	// the two independently running go routines are passing data back and forth from eachother
	time.Sleep(time.Second)
	// wait  a second gives time for the go routines above to finish
}
