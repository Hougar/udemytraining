package main

import "fmt"

func main() {
	for n := range sq(gen(2, 3)) {
		// doing a range over the channel
		fmt.Println(n)
	}
}

func gen(numbers ...int) chan int {
	// returns a chanel
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
	// square takes a channel
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
