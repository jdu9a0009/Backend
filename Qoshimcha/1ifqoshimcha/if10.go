package main

import "fmt"

func main() {
	a := 9
	b := 5

	if a != b {
		a = a + b
		b = a
		fmt.Println(a, b)
	} else {
		a = 0
		b = 0
		fmt.Println(a, b)
	}
}
