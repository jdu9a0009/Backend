//POINTERS

package main

import "fmt"

func main() {
	number := 1
	example(&number)
	fmt.Println("Number: ", number)
}

func example(son *int) {
	fmt.Println("Son: ", *son)
	*son = 23
}
