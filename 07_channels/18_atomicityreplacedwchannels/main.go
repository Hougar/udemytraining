package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := incrementor(2)
	var count int
	for n := range c {
		count++
		fmt.Println(n)
	}
	fmt.Println("Final count: ", count)

}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)
	// semaphore lettign us know when everything is finished

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing: ", k)
			}
			done <- true
		}(1)
	}
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()
	return c
	// c i sonly returned when the for loop above is done
}
