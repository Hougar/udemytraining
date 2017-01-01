// pipe line return channels from functions & accept channels in functions - as arguments
// a series of stages - connected by channels

// receive values from upstream via inbound channels
// perfrom some function on the data and produce a new value
// send values downstream via outbound channels

// example - square output

package main

import "fmt"

func main() {
	c := gen(2, 3)
	out := sq(c)
	// set up the pipeline

	fmt.Println(<-out)
	fmt.Println(<-out)
	//two and three squared print
}

func gen(numbers ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
