package main

//Write a function countVowels that takes a string as input and returns the number of vowels (a, e, i, o, u) in the string.
import "fmt"

func main() {
	string := vowelsCounter("Hello gophers")
	fmt.Println(string)

	string = vowelsCounter("This is vowel counter code using pointers in Golang")
	fmt.Println(string)

	string = vowelsCounter("thanks")
	fmt.Println(string)
}
func is_vowel(char rune) bool {

	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') {

		return true

	} else {

		return false

	}

}

func vowelsCounter(str string) int {
	count := 0
	for _, char := range str {
		if is_vowel(char) {
			count = count + 1
		}
	}
	return count
}
