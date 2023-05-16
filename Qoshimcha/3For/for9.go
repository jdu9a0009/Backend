package main

import "fmt"

func main() {
	a := 2
	b := 5
	res := 0
	for i := a; i <= b; i++ {
		res += i * i
	}
	fmt.Println(res)

}
