// Mutex: Purpose - Lock our code to prevent overwritting - prevent race conditions when runnning concurrent processes
// mutex = mutually exclusive object - mutliple program threads can take turns sharing the same resource

// Have things running concurrently and different processes are trying to access the same variables - you can get some overwritting - the counter gets overwritten since foo & bar are in a race condition. Lock it down

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

// variable mutex is of type sync.Mutex
// needs to be a global variable since it is of higher scope - can lock down both incrementaor foo & bar

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Microsecond)
		{ // added extra curly to signify lockes - brings it over indent
			mutex.Lock()
			// mutex can access lock method
			x := counter
			x++
			counter = x
			fmt.Println(s, i, "Counter:", counter)
			mutex.Unlock()
		}
	}
	wg.Done()
}

// by using mutex counter now provideds the proper 40 sum for the amount of times the loop ran! not overwritten
