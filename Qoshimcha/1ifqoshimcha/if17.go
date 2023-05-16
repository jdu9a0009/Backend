package main

import "fmt"

func main() {
	var a, b, c float32
	a = 5.0
	b = 7.0
	c = 9.0
	if (a > b && b > c) || (a < b && b < c) {
		fmt.Printf("a: %f\n", 2.0*a)
		fmt.Printf("b:%f\n ", 2*b)
		fmt.Printf("c: %f\n", 2*c)
	} else {
		fmt.Printf("a: %f\n", -1*a)
		fmt.Printf("b: %f\n", -1*b)
		fmt.Printf("c: %f\n", -1*c)
	}

}
