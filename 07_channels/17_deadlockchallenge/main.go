package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()
	xc := fanOut(in, 10)
	for n := range merge(xc...) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	// xc := make([]<-chan int, n) - having the n there creates deadlock why?
	// n = length and capacity of channel xc
	// n = 0 since using append
	// creeating more channels than are being used with append below
	var xc []<-chan int
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func factorial(in <-chan int) <-chan int {
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
