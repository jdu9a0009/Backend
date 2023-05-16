package main

import (
	"fmt"
	"time"
)

var input string

func main() {
	fmt.Println("This is Palindrom checker app\n please enter what words you want to check: ")
	time.Sleep(2 * time.Second)
	fmt.Scan(&input)

	result := []byte{}
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}

	if input == string(result) {
		fmt.Println("This is Palindrome")
	} else {
		fmt.Println("This is not Palindrome")
	}
}
