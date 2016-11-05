package main

// gersakbogdan committed on Dec 27, 2014 - taken from this source
import (
	"fmt"
	"time"
)

func fibevennumbers(max int) int {
	a, b, sum := 1, 2, 2
	for {
		a, b = a+2*b, 2*a+3*b
		if b >= max {
			return sum
		}
		sum += b
	}
}

func main() {
	start := time.Now()
	max := 3
	fmt.Println(fibevennumbers(max))
	fmt.Println("Elapsed: ", time.Since(start))
}
