package main

import "fmt"

func main() {
	k := 2000.50
	for i := 0; i <= 10; i++ {
		fmt.Println(i, " kg  konfet = ", float64(i)*k, " so'm")
	}

}
