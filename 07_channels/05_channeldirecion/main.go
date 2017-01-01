// <-chan T   can only receive T
// chan<- T   can only send T
// chan T     bidirectonal - none (send or receive )
// note T is of type i.e. bool, int, etc.

package main

import "fmt"

func main() {
	c := incrementor()
	for n := range puller(c) {
		// refactor - could assign puller(c) to variablee and pass into range loop or could just put the call to puller c in the loop !

		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	// channels that you can only receive from more specific
	out2 := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out2 <- sum
		close(out2)
	}()
	return out2
}
