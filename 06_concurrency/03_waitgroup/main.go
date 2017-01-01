package main

import (
	"fmt"
	"sync"
)

// wg = Wait group see thr results of the go routines - make functions pause
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// add 2 groups to the wait group
	go foo()
	wg.Wait()
	wg.Add(1)
	go bar()
	wg.Wait()
	// wait until the wg turns to 0, 2-1, 1-2, 0
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
	// letting the wg know they are finished
	// will subtracct 1 from the wg
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
