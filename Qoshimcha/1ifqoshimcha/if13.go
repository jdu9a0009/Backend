package main

import "fmt"

func main() {
	a := 30
	b := 50
	c := 10

	if a > b && b > c {
		fmt.Println("O'rtacha son:", b)
	} else if a > c && a < b {
		fmt.Println("O'rtacha son:", a)
	} else {
		fmt.Println("O'rtacha son:", c)
	}

}
