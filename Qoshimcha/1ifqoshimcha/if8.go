package main

import "fmt"

func main() {
	var a, b int
	a = 2
	b = 3
	if a > b {
		fmt.Println(a)
		fmt.Println(b)
	} else {
		fmt.Println(b)
		fmt.Println(a)
	}
}
