package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)
	// FAN OUT - distrubute work accross 2 go routines(c1, c2) that both read from in
	// mutiple functions c1, c2 pulling off 1 channel "out"

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
	// FAN IN - consume the merged input from multiple channels
}

func gen(numbs ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbs {
			// since we are ranging over a slice ( also applies to a map) we get thek KEY and then the VALUE therefore just want the value so blank ID used
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
			// range over a channel gives us the values of the channel
			// when you have a slice can use range
			out <- n * n
			// takes a number off the channel - i..e range of in and mutiplies it by itself and puts it back on to the channel
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	fmt.Printf("TYPEOFcs:%T\n", cs)
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
		// Passing in C back into the function  - C is a chan int
		// by adding c back as an argument limiting the scope of c to this go routine - c can be used elsewhere now
		// creates a new closure
		go func() {
			wg.Wait()
			close(out)
		}()
	}
	return out
}

// WHY FAN OUT AND FAN IN - we have a lot of values that need to be processed, 2 functions which can run on their own go routines , creates an assembly lign with values needing to be processed
