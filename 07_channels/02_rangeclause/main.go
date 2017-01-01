package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		// when a chanells is closed can no longer put values into it
		// can still receive values from the channel but nothing new can go in
		// go func() is an anomyous function which self executs - goal of chennels cordinate all the different processes
		// send this wid receiver off on its own

	}()
	for n := range c {
		fmt.Println(n)
	}
	// for range loop - looping over the c and printing it instead of saying for places 0,1,2,3,4,5,6,7,8,9 range auto adjusts what to print out
	// c countious receives from the channel assigns to n and prints n

	// no longer needs time.sleep   - range clause blocks it
}

// func main is alos a go routine
