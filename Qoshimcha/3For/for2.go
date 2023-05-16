package main

import "fmt"

func main() {
	a := 2
	b := 9
	for i := a; i <= b; i++ {
		fmt.Println(i)
	}
	fmt.Println("N: ", b-a+1)

}
