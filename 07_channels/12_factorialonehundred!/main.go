// execute 100 factorial computation concurrently and in parallel

package main

import "fmt"

func main() {
	in := gen()
	f := factorial(in)
	for n := range f {
		fmt.Println(n)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 2; j < 13; j++ {
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
