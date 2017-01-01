// IS THIS FAN OUT:
// Yes : we have launched several go routines which are publishing a message onto our channel We have mutiple functions reading from the same channel input.
// IS THIS FAN IN: NO
// FAN in- A funciton can read from mutiple inputs and proceed until all are closed by multiplexing the input channel onto single channel thats closed when all the inputs are closed. No channels are converging into one channel here.

package main

import (
	"fmt"
	"time"
)

var workerid int
var publisherid int

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Millisecond)
}

func publisher(out chan string) {
	publisherid++
	thisID := publisherid
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data \n", thisID)
		data := fmt.Sprintf("Data from publisher %d, Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerid++
	thisID := workerid
	for {
		fmt.Printf("%d: wating for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is:%s\n", thisID, input)
	}
}
