package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for j := 1; j < 10; j++ {
		fmt.Println("Foo:", j)
		time.Sleep(time.Duration(3 * time.Millisecond))
		// gives you interleaving foo prints then Bar adds gaps - concurrency running at the same time?
	}
	wg.Done()
}

func bar() {
	for i := 1; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
