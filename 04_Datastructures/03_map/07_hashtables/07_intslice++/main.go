package main

import "fmt"

func main() {
	buckets := make([]int, 1)
	fmt.Println(buckets[0])
	// slice buckets at position 0 has default value of nil (0)
	buckets[0] = 42
	// set position 0 to int value 42
	fmt.Println(buckets[0])
	// print 42 which is now at position 0
	buckets[0]++
	// add 1 to position 0 i.e same as buckets[0] = buckets[0] +1 the value at buckets not the position! - pass by value
	fmt.Println(buckets[0])
	// print 42 which is now at position 0

}
