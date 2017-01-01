package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// init is special the first code that will be executed is func init() always

// run on more than 1 core - no longer needed in new go version uses more than 1 core
func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for j := 1; j < 10; j++ {
		time.Sleep(time.Duration(3 * time.Millisecond))
		fmt.Println("Foo:", j)
	}
	wg.Done()
}

func bar() {
	for i := 1; i < 10; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))

	}
	wg.Done()
}
