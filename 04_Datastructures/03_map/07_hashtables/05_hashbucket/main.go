package main

import "fmt"

func main() {
	n := hashbucket("Go", 12)
	// argument GO and 12
	fmt.Println(n)
}

func hashbucket(word string, buckets int) int {
	// paramater word of type string and buckets of type int
	letter := int(word[0])
	// variable letter definied as the int of the word at position 0 - changing string to int of first letter of word
	bucket := letter % buckets
	// bucket is the remainder of the leetter it and the buckets i..e 12 - how many buckets you want !
	return bucket
}
