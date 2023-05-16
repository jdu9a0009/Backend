package main

import "fmt"

func main() {
	a := 50
	b := 30
	c := 30

	if b == c {
		fmt.Println(1)
	} else if a == c {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}

}
