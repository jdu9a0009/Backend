package main

import "fmt"

func main() {
	a := 2
	b := 9
	for i := b; i >= a; i-- {
		fmt.Println(i)
	}
	fmt.Println("N: ", b-a+1)

}
