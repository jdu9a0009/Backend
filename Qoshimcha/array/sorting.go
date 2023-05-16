package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := [...]int{22, 3, 4, 652, 8, 2}
	fmt.Println("Befor Sorting", numbers)
	sort.Ints(numbers[:]) //// this  thing can sort array or slice
	fmt.Println("After sorting: ", numbers)

}

/*
Given an array arr[] of size N, the task is to sort this array in ascending order in Go lang.
Examples:

Input: arr[] = {0, 23, 14, 12, 9}
Output: {0, 9, 12, 14, 23}

Input: arr[] = {7, 0, 2}
Output: {0, 2, 7}
*/
