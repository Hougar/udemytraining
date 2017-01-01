package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		// close the channel before go routine exists
		// so that range can actually exit and not go on forever
	}()
	for n := range c {
		fmt.Println(n)
	}
}

// why is only zero being printed? - need to range over <-c
// if your gonna rangeneed to close your channel
