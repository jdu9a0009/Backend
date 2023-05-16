package main

import "fmt"

func main() {
	a := 30
	b := 30
	c := 30
	d := 20

	if b == c && c == d {
		fmt.Println(1)
	} else if a == c && a == d {
		fmt.Println(2)
	} else if a == b && b == d {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}

}
