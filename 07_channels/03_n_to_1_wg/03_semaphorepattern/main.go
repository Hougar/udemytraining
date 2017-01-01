//semaphore - variable used to control access to resources - changed depending on conditions
// likea flag - tell a program to do something
// think of dancineg flag guuys on ships to signal messages
package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			// putting on
		}
		done <- true
		// true gets put into channel  done once for loop has writtent to int channel c
	}()
	go func() {
		<-done
		<-done
		// through values away dont neeed to assign blank identifier as in other cases
		// pulling off
		close(c)
		// both go routine channels need to be done before close executes
	}()
	// all 3 go rouutines launch together
	// the reason you need to make another go routine is that all go routines are executed simultaneoulsy - get blocked

	for n := range c {
		fmt.Println(n)
	}
	// range prints when all 3 above go routines are completed - c blocks because unbufffered channel
	// range sees when the channel is closed and drains the values off the channel, prints it and then exists
	// no wait groups required or use of sync package -
	// being purests - boolean logic to solve !
}
