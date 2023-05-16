package main

import "fmt"

func main() {
	numbers := [...]int{22, 3, 4, 652, 8, 2}

	maxnumber := numbers[0]
	minnumbers := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if maxnumber < numbers[i] {
			maxnumber = numbers[i]
		}
	}
	fmt.Println(maxnumber)
	for i := 0; i < len(numbers); i++ {
		if minnumbers > numbers[i] {
			minnumbers = numbers[i]
		}
	}
	fmt.Println(minnumbers)

}
