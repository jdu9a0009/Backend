package main

import "fmt"

func main() {
	/*
		Foydalanuvchi son kirtadi
		shu son 3ga bo'linib 6ga bo'linmasa Good chiqsin aks xolda Bad chiqsin
	*/

	var integer int
	fmt.Printf("Iltmos Butun son kirting\n")
	fmt.Scanf("%v", &integer)

	if integer%3 == 0 && integer%6 != 0 {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}
