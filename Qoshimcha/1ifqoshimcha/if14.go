package main

import "fmt"

func main() {
	a := 4
	b := 7
	c := 3

	if (a <= b) && a <= c {
		fmt.Println("Eng kichik son: ", a)
	} else if b <= a && b <= c {
		fmt.Println("Eng kichik son: ", b)
	} else {
		fmt.Println("Eng kichik son: ", c)
	}

	if a >= b && a >= c {
		fmt.Println("Eng katta son: ", a)
	} else if b >= b && b >= c {
		fmt.Println("Eng katta son: ", b)
	} else {
		fmt.Println("Eng katta son: ", c)
	}
}
