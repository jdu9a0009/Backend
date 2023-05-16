package main

import "fmt"

func main() {
	a := 8.9
	b := 5.6

	c := a
	a = b
	b = c

	fmt.Println(a, b)
}
