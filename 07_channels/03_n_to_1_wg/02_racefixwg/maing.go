// you can have n many functions writting to the same channel

package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()
	// this func is the waiting for the wg group to finish i.e. wg from the 2 above to be done and then will excecute close
	// wg is the same variable everywhere

	for n := range c {
		fmt.Println(n)
	}
	// range is wi; execute once close - i..e the channel is drained
}

// output is still interleaving but NO race condition !
