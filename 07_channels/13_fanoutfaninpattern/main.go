// -------------->=============================================>=======================================>-------------->
//Some channel     Mutiple functions reading from that channel  mutiple function writing to same channel  some channel
//            FANOUT                                                                                 FAN IN
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are both boring I am leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		// not in deadlock because no i condition on this loop
	}()
	return c
}

// FAN IN

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
			// take the value off input1 and put onto channel c
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
