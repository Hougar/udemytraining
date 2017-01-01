package main

import "fmt"

func main() {
	f := factorial(4)
	for n := range f {
		// taking values off
		fmt.Println("Total: ", n)
	}
}
func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		// putting values on
		close(out)
	}()
	return out
}

// Are we blocked from exiting main after processing happens? - yes when range f is created
// Are we putting values onto a channel and at the same time taking values off?

// Why would you want a factorial to be a channel?
// If running thousands of these programs would want to run in parallel & concurrently will speed it up !
