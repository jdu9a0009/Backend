package main

import "fmt"

func main() {
	// Fibonachi sonlari

	var x, y, z = 1, 1, 0
	fmt.Println("Iltimos sonlarni kirting")
	for i := 0; i <= num.length; i++ {
		z = x + y
		x = y
		y = z
		fmt.Println("Fibonachi: ", z)
	}
}
