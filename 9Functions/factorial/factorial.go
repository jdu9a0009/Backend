package main

import "fmt"

func main() {
	result := factorial(5)
	fmt.Println(result) // prints 120

}
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
