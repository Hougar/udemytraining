// another way to sync things which are occuring in different processses

// atomicity - locking a smalller part relative to a mutex - in package sync
// only  one person can aceess it at a time

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)
0
var wg sync.WaitGroup
var counter int64

// when variable counter = int64 likrly doing atomic

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
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// the final counter is correct but the individual counters do not come out in sequence - i.e. when order matters use mutex if not can use atomicity
