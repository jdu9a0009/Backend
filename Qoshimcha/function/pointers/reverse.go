package main

//Write a Go function that takes a pointer to a string and reverses the string in-place (i.e., modifies the original string).
import "fmt"

func main() {
	s := "hello"
	fmt.Println("Before reverse:", s)
	reverseString(&s)
	fmt.Println("After reverse:", s)
}

func reverseString(s *string) {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
}
