package main

import "os"

// func init() {
// 	nf, err := os.Create("log.txt")
// 	// create a log text file
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	log.SetOutput(nf)
// set the output to a new file
// OOH LOOK A TEXT FILE <---- log.txt
// instead of going to the terminal go to this txt file ! - time stamp
// }

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//	fmt.Println("error no tExt file:  ", err)
		// log.Println("error", err)
		// log.Fatalln(err)
		//  // exit status 1 - using OS exit - code 0 = sucess non-zero = error
		panic(err)
		//  // panic is good if you want to see some of the "stack stuff"
		//  // package builtin - function panic is built in - approximazation
		//  // un-comment any of the 1 - 4 above to deal with error
	}
}

// STDOUT VS STDERR - where it goes?
