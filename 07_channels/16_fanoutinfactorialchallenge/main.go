package main

import (
	"fmt"
	"sync"
)

func main() {
	c0 := factorial(gen())
	c1 := factorial(gen())
	c2 := factorial(gen())
	c3 := factorial(gen())
	c4 := factorial(gen())
	c5 := factorial(gen())
	c6 := factorial(gen())
	c7 := factorial(gen())
	c8 := factorial(gen())
	c9 := factorial(gen())
	// FAN OUT ABOVE

	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		fmt.Println(n)
	}
	// FAN IN ABOVE - use range since bloccking to ensure no deadlock
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= n
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
