// can return channles & take channels as arguments
package main

import "fmt"

func main() {
	c := incrementor()
	cSUm := puller(c)
	// passing the channel c into puller
	for n := range cSUm {
		fmt.Println(n)
		// print all the elements of CSum in this case only 1 element
	}
}

func incrementor() chan int {
	// returns type of channel int
	out := make(chan int)
	// creates a channel
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		// launch go routine
		// putting values onto channel out through for loop
		close(out)
		// after all the values are put onto the channel close the channel
	}()
	return out
	// return out assigned to c
}

func puller(c chan int) chan int {
	// takes chanel int assigned to variable c from incrementor
	out2 := make(chan int)
	// creating another out channel within the scop of func puller
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		// go routiine where var sum int assigns sum of all elements c
		out2 <- sum
		close(out2)
	}()
	// assign the sum of all in c to channel out and close it
	return out2
	// return out to be assigned as variable cSum
}

// when assigning a channel always make sure thereis somehting to take information off the channel or else you have a "DEADLOCK" - could also be vice versa
