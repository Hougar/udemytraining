package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}
	go func() {
		for i := 0; i < 10; i++ {
			<-done
		}
		close(c)
	}()
	// clsing 9 go routines from above !
	for n := range c {
		fmt.Println(n)
	}
	// this example many functions writting to the channel but only 1 function PULLIng form the channel
}
