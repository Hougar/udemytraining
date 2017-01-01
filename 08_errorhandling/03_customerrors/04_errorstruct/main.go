package main

import (
	"fmt"
	"log"
)

type mathError struct {
	lat  string
	long string
	err  error
}

func (n *mathError) Error() string {
	return fmt.Sprintf("A math error as occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		me := fmt.Errorf("math redux: imaginary root: %v", f)
		return 0, &mathError{"59.2289N", "99W", me}
	}
	return 42, nil
}
