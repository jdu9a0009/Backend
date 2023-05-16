package main

import "fmt"

func main() {
	k := 2000.50
	for i := 0.1; i <= 1; i += 0.1 {
		fmt.Println(i, " kg  konfet = ", float64(i)*k, " so'm")
	}

}
