package main

import "fmt"

func main() {
	length, summation := sum(10, 20, 30, 40, 50)
	fmt.Println("Lengths of slice: ", length)
	fmt.Println("Sum of slice: ", summation)
}

func sum(nums ...int) (int, int) {
	result := 0
	for _, num := range nums {
		result += num
	}
	return len(nums), result
}
