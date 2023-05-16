package main

import "fmt"

func main() {
	a := 2
	b := 5
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(sum)

}
