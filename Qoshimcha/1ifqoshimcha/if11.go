package main

import "fmt"

func main() {
	a := 5
	b := 5

	if a != b {
		if a > b {
			b = a
			fmt.Println(a, b)
		} else {
			a = b
			fmt.Println(a, b)
		}

	} else {
		a = 0
		b = 0
		fmt.Println(a, b)
	}
}
