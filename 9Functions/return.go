package main

import "fmt"

func main() {
	bir, ikki := split(17)
	fmt.Println("Bir: ", bir, "Ikki: ", ikki)
	l, r := sum(11, 22, 33, 44, 55)
	fmt.Println("Length: ", l, "Result: ", r)

}

func split(sum int) (y, x int) {
	x = sum + 10
	y = sum - x
	return
}

func sum(nums ...int) (length, result int) {
	result = 0
	for _, num := range nums {
		result += num
	}
	length = len(nums)
	return
}
