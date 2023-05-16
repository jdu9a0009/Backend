package main

import "fmt"

func main() {
	a := 9
	b := 7
	c := 5

	if a > b && b > c {
		fmt.Println("O'rtacha son:", b)
	} else if a > c && a < b {
		fmt.Println("O'rtacha son:", a)
	} else {
		fmt.Println("O'rtacha son:", c)
	}

	if a >= b && a >= c {
		fmt.Println("Eng katta son: ", a)
	} else if b >= b && b >= c {
		fmt.Println("Eng katta son: ", b)
	} else {
		fmt.Println("Eng katta son: ", c)
	}
}
