package main

import "fmt"

func main() {
	a := 90
	b := 50
	c := 10

	if a > b && a > c {
		fmt.Println("katta ", a)
	} else if b > a && b > c {
		fmt.Println("katta ", b)
	} else {
		fmt.Println(c)
	}

}
