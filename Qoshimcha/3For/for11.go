package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	res := 0
	for i := 0; i <= a; i++ {
		res += int(math.Pow(a+i), 2)
	}
	fmt.Println(res)

}
