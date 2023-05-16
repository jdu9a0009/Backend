package main

//Write a Go function that takes a pointer to an integer slice and returns the sum of all the elements in the slice.
import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := sumOfSlice(&slice)
	fmt.Println(sum)
}
func sumOfSlice(ptr *[]int) int {
	sum := 0
	for _, num := range *ptr {
		sum += num
	}
	return sum
}
