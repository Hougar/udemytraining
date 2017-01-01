// Have things running concurrently and different processes are trying to access the same variables - you can get some overwritting - the counter gets overwritten since foo & bar are in a race condition

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

// global variable counter is being accessed and overwritten - is the race condition

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter", counter)
	// the counter ran 20 + 20 times = 40 but this program will give you < 40 since overwritting happening
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Microsecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// in cmd prompt can rund go run -race main.go - to see if you have a racecondition
/// the - race is a "flag" cmd
