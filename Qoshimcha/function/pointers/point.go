package main

// Swapping 2 variable values using pointers

import (
	"fmt"
)

func main() {
	var num1, num2, temp int
	fmt.Println("Iltmos 1-va 2- sonlarni birin ketinlikda kiriting")
	fmt.Scanf("%d%d", &num1, &num2)
	fmt.Println("Almashtrishdan oldin")
	fmt.Println("Birinchi son: ", num1)
	fmt.Println("Ikkinchi son: ", num2)
	var x *int
	var y *int
	x = &num1 // x getting num1 addres
	y = &num2 //  y getting num2 adress
	temp = *x //swapping
	*x = *y
	*y = temp
	fmt.Println("Almashtrishdan kegin ")
	fmt.Println("Birinchi son: ", num1)
	fmt.Println("Ikkinchi son: ", num2)
}
