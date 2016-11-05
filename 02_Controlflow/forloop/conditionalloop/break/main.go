package main

import "fmt"

func main() {
	i := 11
	for {
		fmt.Println(i)
		if i >= 10 { // here is the condition if i >=10 then break i.e. stop
			break
		}
		i++
	}
}
