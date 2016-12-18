package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701.txt")
	// declare error and the response from using get
	if err != nil {
		log.Fatal(err)
	}
	// log a fatal error if HTTP has no answer

	// getting the response body  = new scanner takes a reader
	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()
	// close res.body to save ram right before func main exits

	scanner.Split(bufio.ScanWords)
	// splitting moby dick up into words

	buckets := make([]int, 200)
	// creating a slice to hold counts

	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
		// add 1 the the int from hasbucket
	}

	fmt.Println(buckets[65:123])
	// printing from UtF-8 a to z
	// fmt.Println(buckets["********"])
	// for i := 28 ; i <=126; i++ {
	// fmt.Printf("%v - %c - %v \n",i,i buckets[i])
}

// there are how many words which start with A in moby dick? How many with B see result!
// which bucket would have how may entries -> i.e. A is bucket B is a bucket to z !
// are buckets arent even i..e 1 word starts with x but 20000 start with f etc

func HashBucket(word string) int {
	return int(word[0])
	// get the first letter of the word i..e position 0 and convert into an int - number
}
