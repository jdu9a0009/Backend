package main

import "fmt"

func main() {
	a := "Hello"
	fmt.Println(a)

	fmt.Printf("Salom inglizchada %v-> sozi\n", a)

	fmt.Printf("%v %v %v", 11, 22.5, true)
	x, y, z, b := 1, 2.3, "Hello", true

	fmt.Printf("1: %T 2: %T 3: %T 4: %T", x, y, z, b) //TYPE

}
