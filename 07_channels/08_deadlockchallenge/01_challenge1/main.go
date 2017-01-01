package main

import "fmt"

func main() {
	c := make(chan int)
	// go func() {
	c <- 1
	// putting something onthe channel and nothing there to receive it - the code is blocked
	// }()
	fmt.Println(<-c)
}

// This code results in deadlock
// Why -? 1 is being put onto the channell & then print what was received from the channell
// all go routines are asleep - miscommunication between send and receive - how to fix it? go routine
// go func routine fixes it becaus func main launched then go routine then fmt Println
